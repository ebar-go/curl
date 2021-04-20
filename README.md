# curl
基于http.Client实现的curl客户端，提供更为快捷的http接口调用方式。

# 安装
```
go get github.com/ebar-go/curl
```

# 使用说明
## GET请求
```go
package main
import (
	"fmt"
	"github.com/ebar-go/curl"
)
func main() {
    var address = "http://localhost:8080/api"

    response, err := curl.Get(address)
    if err != nil {
        panic(err)
    }
    fmt.Println(response.String()) // 可以通过response获取string类型的响应
    fmt.Println(response.Byte()) // 也可以是byte
    var obj struct{}
    err = response.BindJson(&obj) // 也可以解析到一个结构体
}
```

## POST请求
```
params := make(url.Values)
params.Set("id", "1")
response, err := curl.Post("localhost:8080/test", strings.NewReader(params.Encode()))
```

## POST JSON请求
```
body := strings.NewReader(`{"username":"curl"}`)
response, err := curl.PostJson("localhost:8080/test", body)
```
## PUT请求
```
response, err := curl.Put(address, nil)
```

## PUT请求
```
response, err := curl.Put(address, nil)
```


## Patch请求
```
response, err := curl.Patch(address, nil)
```

## Delete请求
```
response, err := curl.Delete(address, nil)
```

## 自定义请求
```
request, _ := http.NewRequest(http.MethodPost, url, body)
response, err := curl.Send(request)
if err != nil {
    return err
}
fmt.Println(response.String())
```
