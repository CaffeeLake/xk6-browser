package browser

import (
	"context"
	"reflect"
	"strings"
	"testing"

	"github.com/dop251/goja"
	"github.com/stretchr/testify/require"

	"github.com/grafana/xk6-browser/common"

	k6common "go.k6.io/k6/js/common"
	k6modulestest "go.k6.io/k6/js/modulestest"
	k6lib "go.k6.io/k6/lib"
	k6metrics "go.k6.io/k6/metrics"
)

// customMappings is a list of custom mappings for our module API.
// Some of them are wildcards, such as query to $ mapping; and
// others are for publicly accessible fields, such as mapping
// of page.keyboard to Page.getKeyboard.
func customMappings() map[string]string {
	return map[string]string{
		// wildcards
		"pageAPI.query":             "$",
		"pageAPI.queryAll":          "$$",
		"frameAPI.query":            "$",
		"frameAPI.queryAll":         "$$",
		"elementHandleAPI.query":    "$",
		"elementHandleAPI.queryAll": "$$",
		// getters
		"pageAPI.getKeyboard":    "keyboard",
		"pageAPI.getMouse":       "mouse",
		"pageAPI.getTouchscreen": "touchscreen",
		// internal methods
		"elementHandleAPI.objectID":    "",
		"frameAPI.id":                  "",
		"frameAPI.loaderID":            "",
		"JSHandleAPI.objectID":         "",
		"browserAPI.close":             "",
		"frameAPI.evaluateWithContext": "",
		// TODO: browser.on method is unexposed until more event
		// types other than 'disconnect' are supported.
		// See: https://github.com/grafana/xk6-browser/issues/913
		"browserAPI.on": "",
	}
}

// TestMappings tests that all the methods of the API (api/) are
// to the module. This is to ensure that we don't forget to map
// a new method to the module.
func TestMappings(t *testing.T) {
	t.Parallel()

	type test struct {
		apiInterface any
		mapp         func() mapping
	}

	var (
		vu = &k6modulestest.VU{
			RuntimeField: goja.New(),
			InitEnvField: &k6common.InitEnvironment{
				TestPreInitState: &k6lib.TestPreInitState{
					Registry: k6metrics.NewRegistry(),
				},
			},
		}
		customMappings = customMappings()
	)

	// testMapping tests that all the methods of an API are mapped
	// to the module. And wildcards are mapped correctly and their
	// methods are not mapped.
	testMapping := func(t *testing.T, tt test) {
		t.Helper()

		var (
			typ    = reflect.TypeOf(tt.apiInterface).Elem()
			mapped = tt.mapp()
			tested = make(map[string]bool)
		)
		for i := 0; i < typ.NumMethod(); i++ {
			method := typ.Method(i)
			require.NotNil(t, method)

			// goja uses methods that starts with lowercase.
			// so we need to convert the first letter to lowercase.
			m := toFirstLetterLower(method.Name)

			cm, cmok := isCustomMapping(customMappings, typ.Name(), m)
			// if the method is a custom mapping, it should not be
			// mapped to the module. so we should not find it in
			// the mapped methods.
			if _, ok := mapped[m]; cmok && ok {
				t.Errorf("method %q should not be mapped", m)
			}
			// a custom mapping with an empty string means that
			// the method should not exist on the API.
			if cmok && cm == "" {
				continue
			}
			// change the method name if it is mapped to a custom
			// method. these custom methods are not exist on our
			// API. so we need to use the mapped method instead.
			if cmok {
				m = cm
			}
			if _, ok := mapped[m]; !ok {
				t.Errorf("method %q not found", m)
			}
			// to detect if a method is redundantly mapped.
			tested[m] = true
		}
		// detect redundant mappings.
		for m := range mapped {
			if !tested[m] {
				t.Errorf("method %q is redundant", m)
			}
		}
	}

	for name, tt := range map[string]test{
		"browser": {
			apiInterface: (*browserAPI)(nil),
			mapp: func() mapping {
				return mapBrowser(moduleVU{VU: vu})
			},
		},
		"browserContext": {
			apiInterface: (*browserContextAPI)(nil),
			mapp: func() mapping {
				return mapBrowserContext(moduleVU{VU: vu}, &common.BrowserContext{})
			},
		},
		"page": {
			apiInterface: (*pageAPI)(nil),
			mapp: func() mapping {
				return mapPage(moduleVU{VU: vu}, &common.Page{
					Keyboard:    &common.Keyboard{},
					Mouse:       &common.Mouse{},
					Touchscreen: &common.Touchscreen{},
				})
			},
		},
		"elementHandle": {
			apiInterface: (*elementHandleAPI)(nil),
			mapp: func() mapping {
				return mapElementHandle(moduleVU{VU: vu}, &common.ElementHandle{})
			},
		},
		"jsHandle": {
			apiInterface: (*common.JSHandleAPI)(nil),
			mapp: func() mapping {
				return mapJSHandle(moduleVU{VU: vu}, &common.BaseJSHandle{})
			},
		},
		"frame": {
			apiInterface: (*frameAPI)(nil),
			mapp: func() mapping {
				return mapFrame(moduleVU{VU: vu}, &common.Frame{})
			},
		},
		"mapRequest": {
			apiInterface: (*requestAPI)(nil),
			mapp: func() mapping {
				return mapRequest(moduleVU{VU: vu}, &common.Request{})
			},
		},
		"mapResponse": {
			apiInterface: (*responseAPI)(nil),
			mapp: func() mapping {
				return mapResponse(moduleVU{VU: vu}, &common.Response{})
			},
		},
		"mapWorker": {
			apiInterface: (*workerAPI)(nil),
			mapp: func() mapping {
				return mapWorker(moduleVU{VU: vu}, &common.Worker{})
			},
		},
		"mapLocator": {
			apiInterface: (*locatorAPI)(nil),
			mapp: func() mapping {
				return mapLocator(moduleVU{VU: vu}, &common.Locator{})
			},
		},
		"mapConsoleMessage": {
			apiInterface: (*consoleMessageAPI)(nil),
			mapp: func() mapping {
				return mapConsoleMessage(moduleVU{VU: vu}, &common.ConsoleMessage{})
			},
		},
		"mapTouchscreen": {
			apiInterface: (*touchscreenAPI)(nil),
			mapp: func() mapping {
				return mapTouchscreen(moduleVU{VU: vu}, &common.Touchscreen{})
			},
		},
	} {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			testMapping(t, tt)
		})
	}
}

// toFirstLetterLower converts the first letter of the string to lower case.
func toFirstLetterLower(s string) string {
	// Special cases.
	// Instead of loading up an acronyms list, just do this.
	// Good enough for our purposes.
	special := map[string]string{
		"ID":        "id",
		"JSON":      "json",
		"JSONValue": "jsonValue",
		"URL":       "url",
	}
	if v, ok := special[s]; ok {
		return v
	}
	if s == "" {
		return ""
	}

	return strings.ToLower(s[:1]) + s[1:]
}

// isCustomMapping returns true if the method is a custom mapping
// and returns the name of the method to be called instead of the
// original one.
func isCustomMapping(customMappings map[string]string, typ, method string) (string, bool) {
	name := typ + "." + method

	if s, ok := customMappings[name]; ok {
		return s, ok
	}

	return "", false
}

// ----------------------------------------------------------------------------
// JavaScript API definitions.
// ----------------------------------------------------------------------------

// browserAPI is the public interface of a CDP browser.
type browserAPI interface {
	Close()
	Context() *common.BrowserContext
	CloseContext()
	IsConnected() bool
	NewContext(opts goja.Value) (*common.BrowserContext, error)
	NewPage(opts goja.Value) (*common.Page, error)
	On(string) (bool, error)
	UserAgent() string
	Version() string
}

// browserContextAPI is the public interface of a CDP browser context.
type browserContextAPI interface {
	AddCookies(cookies []*common.Cookie) error
	AddInitScript(script goja.Value, arg goja.Value) error
	Browser() *common.Browser
	ClearCookies() error
	ClearPermissions()
	Close()
	Cookies(urls ...string) ([]*common.Cookie, error)
	GrantPermissions(permissions []string, opts goja.Value)
	NewPage() (*common.Page, error)
	Pages() []*common.Page
	SetDefaultNavigationTimeout(timeout int64)
	SetDefaultTimeout(timeout int64)
	SetGeolocation(geolocation goja.Value)
	SetHTTPCredentials(httpCredentials goja.Value)
	SetOffline(offline bool)
	WaitForEvent(event string, optsOrPredicate goja.Value) (any, error)
}

// pageAPI is the interface of a single browser tab.
type pageAPI interface {
	BringToFront()
	Check(selector string, opts goja.Value)
	Click(selector string, opts goja.Value) error
	Close(opts goja.Value) error
	Content() string
	Context() *common.BrowserContext
	Dblclick(selector string, opts goja.Value)
	DispatchEvent(selector string, typ string, eventInit goja.Value, opts goja.Value)
	EmulateMedia(opts goja.Value)
	EmulateVisionDeficiency(typ string)
	Evaluate(pageFunc goja.Value, arg ...goja.Value) any
	EvaluateHandle(pageFunc goja.Value, arg ...goja.Value) (common.JSHandleAPI, error)
	Fill(selector string, value string, opts goja.Value)
	Focus(selector string, opts goja.Value)
	Frames() []*common.Frame
	GetAttribute(selector string, name string, opts goja.Value) goja.Value
	GetKeyboard() *common.Keyboard
	GetMouse() *common.Mouse
	GetTouchscreen() *common.Touchscreen
	Goto(url string, opts goja.Value) (*common.Response, error)
	Hover(selector string, opts goja.Value)
	InnerHTML(selector string, opts goja.Value) string
	InnerText(selector string, opts goja.Value) string
	InputValue(selector string, opts goja.Value) string
	IsChecked(selector string, opts goja.Value) bool
	IsClosed() bool
	IsDisabled(selector string, opts goja.Value) bool
	IsEditable(selector string, opts goja.Value) bool
	IsEnabled(selector string, opts goja.Value) bool
	IsHidden(selector string, opts goja.Value) bool
	IsVisible(selector string, opts goja.Value) bool
	Locator(selector string, opts goja.Value) *common.Locator
	MainFrame() *common.Frame
	On(event string, handler func(*common.ConsoleMessage) error) error
	Opener() pageAPI
	Press(selector string, key string, opts goja.Value)
	Query(selector string) (*common.ElementHandle, error)
	QueryAll(selector string) ([]*common.ElementHandle, error)
	Reload(opts goja.Value) *common.Response
	Screenshot(opts goja.Value) goja.ArrayBuffer
	SelectOption(selector string, values goja.Value, opts goja.Value) []string
	SetContent(html string, opts goja.Value)
	SetDefaultNavigationTimeout(timeout int64)
	SetDefaultTimeout(timeout int64)
	SetExtraHTTPHeaders(headers map[string]string)
	SetInputFiles(selector string, files goja.Value, opts goja.Value)
	SetViewportSize(viewportSize goja.Value)
	Tap(selector string, opts goja.Value) (*goja.Promise, error)
	TextContent(selector string, opts goja.Value) string
	ThrottleCPU(common.CPUProfile) error
	ThrottleNetwork(common.NetworkProfile) error
	Title() string
	Type(selector string, text string, opts goja.Value)
	Uncheck(selector string, opts goja.Value)
	URL() string
	ViewportSize() map[string]float64
	WaitForFunction(fn, opts goja.Value, args ...goja.Value) (any, error)
	WaitForLoadState(state string, opts goja.Value)
	WaitForNavigation(opts goja.Value) (*common.Response, error)
	WaitForSelector(selector string, opts goja.Value) (*common.ElementHandle, error)
	WaitForTimeout(timeout int64)
	Workers() []*common.Worker
}

// consoleMessageAPI is the interface of a console message.
type consoleMessageAPI interface {
	Args() []common.JSHandleAPI
	Page() *common.Page
	Text() string
	Type() string
}

// frameAPI is the interface of a CDP target frame.
type frameAPI interface {
	Check(selector string, opts goja.Value)
	ChildFrames() []*common.Frame
	Click(selector string, opts goja.Value) error
	Content() string
	Dblclick(selector string, opts goja.Value)
	DispatchEvent(selector string, typ string, eventInit goja.Value, opts goja.Value)
	// EvaluateWithContext for internal use only
	EvaluateWithContext(ctx context.Context, pageFunc goja.Value, args ...goja.Value) (any, error)
	Evaluate(pageFunc goja.Value, args ...goja.Value) any
	EvaluateHandle(pageFunc goja.Value, args ...goja.Value) (common.JSHandleAPI, error)
	Fill(selector string, value string, opts goja.Value)
	Focus(selector string, opts goja.Value)
	FrameElement() (*common.ElementHandle, error)
	GetAttribute(selector string, name string, opts goja.Value) goja.Value
	Goto(url string, opts goja.Value) (*common.Response, error)
	Hover(selector string, opts goja.Value)
	InnerHTML(selector string, opts goja.Value) string
	InnerText(selector string, opts goja.Value) string
	InputValue(selector string, opts goja.Value) string
	IsChecked(selector string, opts goja.Value) bool
	IsDetached() bool
	IsDisabled(selector string, opts goja.Value) bool
	IsEditable(selector string, opts goja.Value) bool
	IsEnabled(selector string, opts goja.Value) bool
	IsHidden(selector string, opts goja.Value) bool
	IsVisible(selector string, opts goja.Value) bool
	ID() string
	LoaderID() string
	Locator(selector string, opts goja.Value) *common.Locator
	Name() string
	Query(selector string) (*common.ElementHandle, error)
	QueryAll(selector string) ([]*common.ElementHandle, error)
	Page() *common.Page
	ParentFrame() *common.Frame
	Press(selector string, key string, opts goja.Value)
	SelectOption(selector string, values goja.Value, opts goja.Value) []string
	SetContent(html string, opts goja.Value)
	SetInputFiles(selector string, files goja.Value, opts goja.Value)
	Tap(selector string, opts goja.Value) (*goja.Promise, error)
	TextContent(selector string, opts goja.Value) string
	Title() string
	Type(selector string, text string, opts goja.Value)
	Uncheck(selector string, opts goja.Value)
	URL() string
	WaitForFunction(pageFunc, opts goja.Value, args ...goja.Value) (any, error)
	WaitForLoadState(state string, opts goja.Value)
	WaitForNavigation(opts goja.Value) (*common.Response, error)
	WaitForSelector(selector string, opts goja.Value) (*common.ElementHandle, error)
	WaitForTimeout(timeout int64)
}

// elementHandleAPI is the interface of an in-page DOM element.
type elementHandleAPI interface {
	common.JSHandleAPI

	BoundingBox() *common.Rect
	Check(opts goja.Value)
	Click(opts goja.Value) error
	ContentFrame() (*common.Frame, error)
	Dblclick(opts goja.Value)
	DispatchEvent(typ string, props goja.Value)
	Fill(value string, opts goja.Value)
	Focus()
	GetAttribute(name string) goja.Value
	Hover(opts goja.Value)
	InnerHTML() string
	InnerText() string
	InputValue(opts goja.Value) string
	IsChecked() bool
	IsDisabled() bool
	IsEditable() bool
	IsEnabled() bool
	IsHidden() bool
	IsVisible() bool
	OwnerFrame() (*common.Frame, error)
	Press(key string, opts goja.Value)
	Query(selector string) (*common.ElementHandle, error)
	QueryAll(selector string) ([]*common.ElementHandle, error)
	Screenshot(opts goja.Value) goja.ArrayBuffer
	ScrollIntoViewIfNeeded(opts goja.Value)
	SelectOption(values goja.Value, opts goja.Value) []string
	SelectText(opts goja.Value)
	SetInputFiles(files goja.Value, opts goja.Value)
	Tap(opts goja.Value) (*goja.Promise, error)
	TextContent() string
	Type(text string, opts goja.Value)
	Uncheck(opts goja.Value)
	WaitForElementState(state string, opts goja.Value)
	WaitForSelector(selector string, opts goja.Value) (*common.ElementHandle, error)
}

// requestAPI is the interface of an HTTP request.
type requestAPI interface {
	AllHeaders() map[string]string
	Frame() *common.Frame
	HeaderValue(string) goja.Value
	Headers() map[string]string
	HeadersArray() []common.HTTPHeader
	IsNavigationRequest() bool
	Method() string
	PostData() string
	PostDataBuffer() goja.ArrayBuffer
	ResourceType() string
	Response() *common.Response
	Size() common.HTTPMessageSize
	Timing() goja.Value
	URL() string
}

// responseAPI is the interface of an HTTP response.
type responseAPI interface {
	AllHeaders() map[string]string
	Body() goja.ArrayBuffer
	Frame() *common.Frame
	HeaderValue(string) goja.Value
	HeaderValues(string) []string
	Headers() map[string]string
	HeadersArray() []common.HTTPHeader
	JSON() goja.Value
	Ok() bool
	Request() *common.Request
	SecurityDetails() goja.Value
	ServerAddr() goja.Value
	Size() common.HTTPMessageSize
	Status() int64
	StatusText() string
	URL() string
}

// locatorAPI represents a way to find element(s) on a page at any moment.
type locatorAPI interface {
	Clear(opts *common.FrameFillOptions) error
	Click(opts goja.Value) error
	Dblclick(opts goja.Value)
	Check(opts goja.Value)
	Uncheck(opts goja.Value)
	IsChecked(opts goja.Value) bool
	IsEditable(opts goja.Value) bool
	IsEnabled(opts goja.Value) bool
	IsDisabled(opts goja.Value) bool
	IsVisible(opts goja.Value) bool
	IsHidden(opts goja.Value) bool
	Fill(value string, opts goja.Value)
	Focus(opts goja.Value)
	GetAttribute(name string, opts goja.Value) goja.Value
	InnerHTML(opts goja.Value) string
	InnerText(opts goja.Value) string
	TextContent(opts goja.Value) string
	InputValue(opts goja.Value) string
	SelectOption(values goja.Value, opts goja.Value) []string
	Press(key string, opts goja.Value)
	Type(text string, opts goja.Value)
	Hover(opts goja.Value)
	Tap(opts goja.Value) (*goja.Promise, error)
	DispatchEvent(typ string, eventInit, opts goja.Value)
	WaitFor(opts goja.Value)
}

// keyboardAPI is the interface of a keyboard input device.
// TODO: map this to page.GetKeyboard(). Currently, the common.Keyboard type
// mapping is not tested using this interface. We use the concrete type
// without testing its exported methods.
type keyboardAPI interface { //nolint: unused
	Down(key string)
	InsertText(char string)
	Press(key string, opts goja.Value)
	Type(text string, opts goja.Value)
	Up(key string)
}

// touchscreenAPI is the interface of a touchscreen.
type touchscreenAPI interface {
	Tap(x float64, y float64) *goja.Promise
}

// mouseAPI is the interface of a mouse input device.
// TODO: map this to page.GetMouse(). Currently, the common.Mouse type
// mapping is not tested using this interface. We use the concrete type
// without testing its exported methods.
type mouseAPI interface { //nolint: unused
	Click(x float64, y float64, opts goja.Value)
	DblClick(x float64, y float64, opts goja.Value)
	Down(x float64, y float64, opts goja.Value)
	Move(x float64, y float64, opts goja.Value)
	Up(x float64, y float64, opts goja.Value)
	// Wheel(opts goja.Value)
}

// workerAPI is the interface of a web worker.
type workerAPI interface {
	URL() string
}
