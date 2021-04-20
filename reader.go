package curl

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)


// http响应内容reader
type responseReader struct {
	pool sync.Pool
}


var bodyReader = newResponseReader()

// getResponseReader
func newResponseReader() *responseReader {
	return &responseReader{pool: sync.Pool{New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 4096))
	}}}
}

// Read 读取流数据
func (adapter *responseReader) read(reader io.Reader) ([]byte, error) {
	buffer := adapter.pool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer func() {
		if buffer != nil {
			adapter.pool.Put(buffer)
			buffer = nil
		}
	}()
	_, err := io.Copy(buffer, reader)

	if err != nil {
		return nil, fmt.Errorf("failed to read respone:%s", err.Error())
	}

	return buffer.Bytes(), nil
}
