package browser

import (
	"fmt"

	"github.com/grafana/sobek"

	"github.com/grafana/xk6-browser/common"
	"github.com/grafana/xk6-browser/k6ext"
)

func mapKeyboard(vu moduleVU, kb *common.Keyboard) mapping {
	return mapping{
		"down": func(key string) *sobek.Promise {
			return k6ext.Promise(vu.Context(), func() (any, error) {
				return nil, kb.Down(key) //nolint:wrapcheck
			})
		},
		"up": func(key string) *sobek.Promise {
			return k6ext.Promise(vu.Context(), func() (any, error) {
				return nil, kb.Up(key) //nolint:wrapcheck
			})
		},
		"press": func(key string, opts sobek.Value) *sobek.Promise {
			return k6ext.Promise(vu.Context(), func() (any, error) {
				kbdOpts := common.NewKeyboardOptions()
				if err := kbdOpts.Parse(vu.Context(), opts); err != nil {
					return nil, fmt.Errorf("parsing keyboard options: %w", err)
				}

				return nil, kb.Press(key, kbdOpts) //nolint:wrapcheck
			})
		},
		"type": func(text string, opts sobek.Value) *sobek.Promise {
			return k6ext.Promise(vu.Context(), func() (any, error) {
				return nil, kb.Type(text, opts) //nolint:wrapcheck
			})
		},
		"insertText": func(text string) *sobek.Promise {
			return k6ext.Promise(vu.Context(), func() (any, error) {
				return nil, kb.InsertText(text) //nolint:wrapcheck
			})
		},
	}
}
