package chromium

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/grafana/xk6-browser/common"
)

// Allocator provides facilities for finding, running, and interacting with a Chromium browser.
type Allocator struct {
	execPath string    // path to the Chromium executable
	storage  DataStore // stores temporary data for the extension and user
}

// Allocate starts a new Chromium browser process and returns it.
// flags are CLI flags to pass to the Chromium executable,
// and env is environment variables to pass to the Chromium executable.
//nolint:funlen
func (a *Allocator) Allocate(
	ctx context.Context,
	opts *common.LaunchOptions,
	flags map[string]interface{},
	env []string,
) (_ *common.BrowserProcess, rerr error) {
	// use the provided directory or create a temporary one.
	if err := a.storage.Make("", flags["user-data-dir"]); err != nil {
		return nil, fmt.Errorf("cannot make temp data directory: %w", err)
	}
	// add dir to flags so that parseArgs can parse it.
	flags["user-data-dir"] = a.storage.Dir

	args, err := parseArgs(flags)
	if err != nil {
		return nil, fmt.Errorf("cannot parse args: %w", err)
	}

	a.execPath = opts.ExecutablePath
	if a.execPath == "" {
		a.execPath = findExecPath()
	}
	ctx, cancel := context.WithCancel(ctx)
	cmd := exec.CommandContext(ctx, a.execPath, args...) //nolint:gosec
	killAfterParent(cmd)

	defer func() {
		if rerr != nil {
			cancel()
			a.storage.Cleanup()
		}
	}()

	// Pipe stderr to stdout
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("cannot pipe stdout: %w", err)
	}
	defer func() {
		_ = stdout.Close()
	}()
	cmd.Stderr = cmd.Stdout

	// Set up environment variable for process
	if len(env) > 0 {
		cmd.Env = append(os.Environ(), env...)
	}

	// We must start the cmd before calling cmd.Wait, as otherwise the two
	// can run into a data race.
	err = cmd.Start()
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("cannot find browser executable in %s", a.execPath)
	}
	if err != nil {
		return nil, fmt.Errorf("cannot start browser executable: %w", err)
	}
	if ctx.Err() != nil {
		return nil, fmt.Errorf("context err after command start: %w", ctx.Err())
	}

	go func() {
		_ = cmd.Wait()
		a.storage.Cleanup()
	}()

	ctxTimeout, cancel := context.WithTimeout(ctx, opts.Timeout)
	defer cancel()

	wsURL, err := parseWebsocketURL(ctxTimeout, stdout)
	if err != nil {
		return nil, fmt.Errorf("cannot parse websocket url: %w", err)
	}

	return common.NewBrowserProcess(ctx, cancel, cmd.Process, wsURL, a.storage.Dir), nil
}

// parseArgs parses command-line arguments and returns them.
func parseArgs(flags map[string]interface{}) ([]string, error) {
	// Build command line args list
	var args []string
	for name, value := range flags {
		switch value := value.(type) {
		case string:
			args = append(args, fmt.Sprintf("--%s=%s", name, value))
		case bool:
			if value {
				args = append(args, fmt.Sprintf("--%s", name))
			}
		default:
			return nil, errors.New("invalid browser command line flag")
		}
	}
	if _, ok := flags["no-sandbox"]; !ok && os.Getuid() == 0 {
		// Running as root, for example in a Linux container. Chromium
		// needs --no-sandbox when running as root, so make that the
		// default, unless the user set "no-sandbox": false.
		args = append(args, "--no-sandbox")
	}
	if _, ok := flags["remote-debugging-port"]; !ok {
		args = append(args, "--remote-debugging-port=0")
	}

	// Force the first page to be blank, instead of the welcome page;
	// --no-first-run doesn't enforce that.
	// args = append(args, "about:blank")
	// args = append(args, "--no-startup-window")
	return args, nil
}

// parseWebsocketURL grabs the websocket address from chrome's output and returns it.
func parseWebsocketURL(ctx context.Context, rc io.Reader) (wsURL string, _ error) {
	type result struct {
		wsURL string
		err   error
	}
	c := make(chan result, 1)
	go func() {
		const prefix = "DevTools listening on "

		scanner := bufio.NewScanner(rc)
		for scanner.Scan() {
			if s := scanner.Text(); strings.HasPrefix(s, prefix) {
				c <- result{
					strings.TrimPrefix(strings.TrimSpace(s), prefix),
					nil,
				}
				return
			}
		}
		if err := scanner.Err(); err != nil {
			c <- result{"", fmt.Errorf("scanner err: %w", err)}
		}
	}()
	select {
	case r := <-c:
		return r.wsURL, r.err
	case <-ctx.Done():
		return "", fmt.Errorf("ctx err: %w", ctx.Err())
	}
}

// fincExecPath finds the path to the Chromium executable and returns it.
func findExecPath() string {
	for _, path := range [...]string{
		// Unix-like
		"headless_shell",
		"headless-shell",
		"chromium",
		"chromium-browser",
		"google-chrome",
		"google-chrome-stable",
		"google-chrome-beta",
		"google-chrome-unstable",
		"/usr/bin/google-chrome",

		// Windows
		"chrome",
		"chrome.exe", // in case PATHEXT is misconfigured
		`C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`,
		`C:\Program Files\Google\Chrome\Application\chrome.exe`,
		filepath.Join(os.Getenv("USERPROFILE"), `AppData\Local\Google\Chrome\Application\chrome.exe`),

		// Mac (from https://commondatastorage.googleapis.com/chromium-browser-snapshots/index.html?prefix=Mac/857950/)
		"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
		"/Applications/Chromium.app/Contents/MacOS/Chromium",
	} {
		if _, err := exec.LookPath(path); err == nil {
			return path
		}
	}

	return ""
}
