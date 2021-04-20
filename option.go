package curl

import (
	"net/http"
	"time"
)

type Option interface {
	apply(options *options)
}

// options 配置项
type options struct {
	// 超时
	timeout time.Duration
	// transport
	transport *http.Transport
}

type timeoutOption time.Duration

func (o timeoutOption) apply(options *options) {
	options.timeout = time.Duration(o)
}
// WithTimeout set http client timeout option
func WithTimeout(timeout time.Duration) Option {
	return timeoutOption(timeout)
}

type httpTransportOption struct {
	transport *http.Transport
}

func (o *httpTransportOption) apply(options *options) {
	if o.transport == nil {
		return
	}
	options.transport =o.transport
}

// WithHttpTransport set http transport option
func WithHttpTransport(transport *http.Transport) Option {
	return &httpTransportOption{transport: transport}
}