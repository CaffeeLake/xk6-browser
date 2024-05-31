package browser

import (
	"github.com/dop251/goja"

	"github.com/grafana/xk6-browser/common"
	"github.com/grafana/xk6-browser/k6ext"
)

// mapResponse to the JS module.
func mapResponse(vu moduleVU, r *common.Response) mapping { //nolint:funlen
	if r == nil {
		return nil
	}
	maps := mapping{
		"allHeaders": func() *goja.Promise {
			return k6ext.Promise(vu.Context(), func() (any, error) {
				return r.AllHeaders(), nil
			})
		},
		"body": func() *goja.Promise {
			return k6ext.Promise(vu.Context(), func() (any, error) {
				body, err := r.Body()
				if err != nil {
					return nil, err //nolint: wrapcheck
				}
				buf := vu.Runtime().NewArrayBuffer(body)
				return &buf, nil
			})
		},
		"frame": func() mapping {
			return mapFrame(vu, r.Frame())
		},
		"headerValue": func(name string) *goja.Promise {
			return k6ext.Promise(vu.Context(), func() (any, error) {
				v, ok := r.HeaderValue(name)
				if !ok {
					return nil, nil
				}
				return v, nil
			})
		},
		"headerValues": func(name string) *goja.Promise {
			return k6ext.Promise(vu.Context(), func() (any, error) {
				return r.HeaderValues(name), nil
			})
		},
		"headers": r.Headers,
		"headersArray": func() *goja.Promise {
			return k6ext.Promise(vu.Context(), func() (any, error) {
				return r.HeadersArray(), nil
			})
		},
		"json": func() *goja.Promise {
			return k6ext.Promise(vu.Context(), func() (any, error) {
				return r.JSON() //nolint: wrapcheck
			})
		},
		"ok": r.Ok,
		"request": func() mapping {
			return mapRequest(vu, r.Request())
		},
		"securityDetails": func() *goja.Promise {
			return k6ext.Promise(vu.Context(), func() (any, error) {
				return r.SecurityDetails(), nil
			})
		},
		"serverAddr": func() *goja.Promise {
			return k6ext.Promise(vu.Context(), func() (any, error) {
				return r.ServerAddr(), nil
			})
		},
		"size": func() *goja.Promise {
			return k6ext.Promise(vu.Context(), func() (any, error) {
				return r.Size(), nil
			})
		},
		"status":     r.Status,
		"statusText": r.StatusText,
		"url":        r.URL,
		"text": func() *goja.Promise {
			return k6ext.Promise(vu.Context(), func() (any, error) {
				return r.Text() //nolint:wrapcheck
			})
		},
	}

	return maps
}
