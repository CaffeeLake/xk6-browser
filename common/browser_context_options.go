package common

import (
	"context"
	"fmt"

	"github.com/grafana/sobek"

	"github.com/grafana/xk6-browser/k6ext"
)

// Geolocation represents a geolocation.
type Geolocation struct {
	Latitude  float64 `js:"latitude"`
	Longitude float64 `js:"longitude"`
	Accuracy  float64 `js:"accuracy"`
}

// NewGeolocation creates a new instance of Geolocation.
func NewGeolocation() *Geolocation {
	return &Geolocation{}
}

// Validate validates the [Geolocation].
func (g *Geolocation) Validate() error {
	if g.Longitude < -180 || g.Longitude > 180 {
		return fmt.Errorf(`invalid longitude "%.2f": precondition -180 <= LONGITUDE <= 180 failed`, g.Longitude)
	}
	if g.Latitude < -90 || g.Latitude > 90 {
		return fmt.Errorf(`invalid latitude "%.2f": precondition -90 <= LATITUDE <= 90 failed`, g.Latitude)
	}
	if g.Accuracy < 0 {
		return fmt.Errorf(`invalid accuracy "%.2f": precondition 0 <= ACCURACY failed`, g.Accuracy)
	}
	return nil
}

// Parse parses the geolocation options.
func (g *Geolocation) Parse(ctx context.Context, sopts sobek.Value) error { //nolint:cyclop
	var newgl Geolocation

	if !sobekValueExists(sopts) {
		return fmt.Errorf("geolocation options are required")
	}

	opts := sopts.ToObject(k6ext.Runtime(ctx))
	for _, k := range opts.Keys() {
		switch k {
		case "accuracy":
			newgl.Accuracy = opts.Get(k).ToFloat()
		case "latitude":
			newgl.Latitude = opts.Get(k).ToFloat()
		case "longitude":
			newgl.Longitude = opts.Get(k).ToFloat()
		}
	}

	*g = newgl

	return nil
}

// BrowserContextOptions stores browser context options.
type BrowserContextOptions struct {
	AcceptDownloads   bool              `js:"acceptDownloads"`
	DownloadsPath     string            `js:"downloadsPath"`
	BypassCSP         bool              `js:"bypassCSP"`
	ColorScheme       ColorScheme       `js:"colorScheme"`
	DeviceScaleFactor float64           `js:"deviceScaleFactor"`
	ExtraHTTPHeaders  map[string]string `js:"extraHTTPHeaders"`
	Geolocation       *Geolocation      `js:"geolocation"`
	HasTouch          bool              `js:"hasTouch"`
	HTTPCredentials   Credentials       `js:"httpCredentials"`
	IgnoreHTTPSErrors bool              `js:"ignoreHTTPSErrors"`
	IsMobile          bool              `js:"isMobile"`
	JavaScriptEnabled bool              `js:"javaScriptEnabled"`
	Locale            string            `js:"locale"`
	Offline           bool              `js:"offline"`
	Permissions       []string          `js:"permissions"`
	ReducedMotion     ReducedMotion     `js:"reducedMotion"`
	Screen            Screen            `js:"screen"`
	TimezoneID        string            `js:"timezoneID"`
	UserAgent         string            `js:"userAgent"`
	VideosPath        string            `js:"videosPath"`
	Viewport          Viewport          `js:"viewport"`
}

// NewBrowserContextOptions creates a default set of browser context options.
func NewBrowserContextOptions() *BrowserContextOptions {
	return &BrowserContextOptions{
		ColorScheme:       ColorSchemeLight,
		DeviceScaleFactor: 1.0,
		ExtraHTTPHeaders:  make(map[string]string),
		JavaScriptEnabled: true,
		Locale:            DefaultLocale,
		Permissions:       []string{},
		ReducedMotion:     ReducedMotionNoPreference,
		Screen:            Screen{Width: DefaultScreenWidth, Height: DefaultScreenHeight},
		Viewport:          Viewport{Width: DefaultScreenWidth, Height: DefaultScreenHeight},
	}
}

// GrantPermissionsOptions is used by BrowserContext.GrantPermissions.
type GrantPermissionsOptions struct {
	Origin string
}
