package curl

import (
	"io"
	"net/http"
	"time"
)

var _default = New(WithTimeout(time.Second * 20))

// Default return default instance of Curl
func Default() Curl {
	return _default
}
// Get send get request
func Get(url string) (Response, error) {
	return Default().Get(url)
}
// Post send post request
func Post(url string, body io.Reader) (Response, error) {
	return Default().Post(url, body)
}

// PostJson send post request with content-type: application/json
func PostJson(url string, body io.Reader) (Response, error) {
	return Default().PostJson(url, body)
}
// Put send put request
func Put(url string, body io.Reader) (Response, error) {
	return Default().Put(url, body)
}
// Patch send patch request
func Patch(url string, body io.Reader) (Response, error) {
	return Default().Patch(url, body)
}
// Delete send delete request
func Delete(url string) (Response, error) {
	return Default().Delete(url)
}

// PostFile upload file by post request
func PostFile(url string, files map[string]string, params map[string]string) (Response, error) {
	return Default().PostFile(url, files, params)
}
// Send send custom http request
func Send(request *http.Request) (Response, error) {
	return Default().Send(request)
}
