package curl

import (
	"io"
	"net"
	"net/http"
	"time"
)

// Curl
type Curl interface {
	// send get request
	Get(url string) (Response, error)
	// send post request
	Post(url string, body io.Reader) (Response, error)
	// send put request
	Put(url string, body io.Reader) (Response, error)
	// send patch request
	Patch(url string, body io.Reader) (Response, error)
	// send delete request
	Delete(url string) (Response, error)
	// post file
	PostFile(url string, files map[string]string, params map[string]string) (Response, error)
	// send http request
	Send(request *http.Request) (Response, error)
	// send post json request
	PostJson(url string, body io.Reader) (Response, error)
}

// Response
type Response interface {
	// get string
	String() string
	// get byte
	Byte() []byte
	// json marshal
	BindJson(object interface{}) error
	// get reader
	Reader() io.Reader
}

// New return curl instance
func New(opts ...Option) Curl {
	options := options{
		timeout: time.Second * 30,
		transport: &http.Transport{ // 配置连接池
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			IdleConnTimeout: 3 * time.Second,
		},
	}

	// apply option
	for _, option := range opts {
		option.apply(&options)
	}
	return &curl{
		client: http.Client{
			Transport: options.transport,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       options.timeout,
		},
	}
}
