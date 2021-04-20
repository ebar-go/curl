package curl

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)


// curl instance of Curl interface
type curl struct {
	// http client
	client http.Client
}


// Get
func (c *curl) Get(url string) (Response, error) {
	return c.dispatch(http.MethodGet, url, nil)
}

// Post
func (c *curl) Post(url string, body io.Reader) (Response, error) {
	return c.dispatch(http.MethodPost, url, body)
}

// Put
func (c *curl) Put(url string, body io.Reader) (Response, error) {
	return c.dispatch(http.MethodPut, url, body)
}

func (c *curl) PostJson(url string, body io.Reader) (Response, error){
	return c.dispatchWithContentType(http.MethodPost, url, body, "application/json;charset=utf8")
}

// Patch
func (c *curl) Patch(url string, body io.Reader) (Response, error) {
	return c.dispatch(http.MethodDelete, url, nil)
}

// Delete
func (c *curl) Delete(url string) (Response, error) {
	return c.dispatch(http.MethodDelete, url, nil)
}

// PostFile 上传文件
func (c *curl) PostFile(url string, files map[string]string, params map[string]string) (Response, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	// 添加form参数
	for name, value := range params {
		_ = writer.WriteField(name, value)
	}

	// 写入文件流
	for field, path := range files {
		// 读取文件
		file, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("Open File: %v", err)
		}
		_ = file.Close()

		// 写入writer
		part, err := writer.CreateFormFile(field, filepath.Base(path))
		if err != nil {
			return nil, fmt.Errorf("Create Form File: %v", err)
		}
		_, err = io.Copy(part, file)
		if err != nil {
			return nil, err
		}
	}

	// 必须close，这样writer.FormDataContentType()才正确
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("Close Writer: %v", err)
	}

	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())

	return c.Send(request)
}

// Send send http request
func (c *curl) Send(request *http.Request) (Response, error) {
	resp, err := c.client.Do(request)
	// close response
	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()

	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, fmt.Errorf("no response")
	}

	body, err := bodyReader.read(resp.Body)
	if err != nil {
		return nil, err
	}
	return &response{body: body}, nil
}

// dispatch
func (c *curl) dispatch(method, url string, body io.Reader) (Response, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	return c.Send(request)
}
func (c *curl) dispatchWithContentType(method, url string, body io.Reader, contentType string) (Response, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", contentType)
	return c.Send(request)
}
