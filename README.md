# protoc-gen-go-http

从 protobuf 文件中生成使用 golang http handler


## 安装

```shell
go install github.com/goapt/cmd/protoc-gen-go-http@latest
```

## 生成

```shell
protoc --proto_path=. \
--proto_path=./third_party \
--go-http_out=paths=source_relative:. \
./example.proto
```

## 合约
你需要实现合约`[contract](contract)`中的接口
Codec: 用户解析请求和返回数据
ServeMux: 可以使用http.ServeMux

```go
mux := http.NewServeMux()

apiv1.RegisterUserServiceHTTPServer(mux,codec,userService)

server := &http.Server{
    Addr:    ":8080", 
    Handler: mux, 
}

if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
    fmt.Printf("Server startup failed: %v\n", err)
}
```
