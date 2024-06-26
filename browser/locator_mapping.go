package browser

import (
	"fmt"

	"github.com/dop251/goja"

	"github.com/grafana/xk6-browser/common"
	"github.com/grafana/xk6-browser/k6ext"
)

// mapLocator API to the JS module.
func mapLocator(vu moduleVU, lo *common.Locator) mapping {
	return mapping{
		"clear": func(opts goja.Value) error {
			ctx := vu.Context()

			copts := common.NewFrameFillOptions(lo.Timeout())
			if err := copts.Parse(ctx, opts); err != nil {
				return fmt.Errorf("parsing clear options: %w", err)
			}

			return lo.Clear(copts) //nolint:wrapcheck
		},
		"click": func(opts goja.Value) (*goja.Promise, error) {
			popts, err := parseFrameClickOptions(vu.Context(), opts, lo.Timeout())
			if err != nil {
				return nil, err
			}

			return k6ext.Promise(vu.Context(), func() (any, error) {
				return nil, lo.Click(popts) //nolint:wrapcheck
			}), nil
		},
		"dblclick":     lo.Dblclick,
		"check":        lo.Check,
		"uncheck":      lo.Uncheck,
		"isChecked":    lo.IsChecked,
		"isEditable":   lo.IsEditable,
		"isEnabled":    lo.IsEnabled,
		"isDisabled":   lo.IsDisabled,
		"isVisible":    lo.IsVisible,
		"isHidden":     lo.IsHidden,
		"fill":         lo.Fill,
		"focus":        lo.Focus,
		"getAttribute": lo.GetAttribute,
		"innerHTML":    lo.InnerHTML,
		"innerText":    lo.InnerText,
		"textContent":  lo.TextContent,
		"inputValue":   lo.InputValue,
		"selectOption": lo.SelectOption,
		"press":        lo.Press,
		"type":         lo.Type,
		"hover":        lo.Hover,
		"tap": func(opts goja.Value) (*goja.Promise, error) {
			copts := common.NewFrameTapOptions(lo.DefaultTimeout())
			if err := copts.Parse(vu.Context(), opts); err != nil {
				return nil, fmt.Errorf("parsing locator tap options: %w", err)
			}
			return k6ext.Promise(vu.Context(), func() (any, error) {
				return nil, lo.Tap(copts) //nolint:wrapcheck
			}), nil
		},
		"dispatchEvent": func(typ string, eventInit, opts goja.Value) error {
			popts := common.NewFrameDispatchEventOptions(lo.DefaultTimeout())
			if err := popts.Parse(vu.Context(), opts); err != nil {
				return fmt.Errorf("parsing locator dispatch event options: %w", err)
			}
			return lo.DispatchEvent(typ, exportArg(eventInit), popts) //nolint:wrapcheck
		},
		"waitFor": lo.WaitFor,
	}
}
