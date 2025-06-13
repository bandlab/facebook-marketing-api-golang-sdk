package fb

import "net/http"

type Options struct {
	transport func(http.RoundTripper) http.RoundTripper
}

type Option func(*Options)

// WithTransport allows modifying the transport of the http.Client.
func WithTransport(fn func(http.RoundTripper) http.RoundTripper) Option {
	return func(o *Options) {
		o.transport = fn
	}
}
