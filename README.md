# protoc-gen-go-http

从 `protobuf` 文件生成 Go HTTP Handler Stub，减少样板代码编写，专注业务 `service` 实现。

## 目标

- 基于 `google.api.http` 注解生成 HTTP 路由与处理函数。
- 业务只需要实现 `XXXHTTPServer` 接口，不再手写重复的 `Decode/Validate/Encode` 模板代码。
- 不依赖外部 runtime 包：运行时接口按包生成到 `http_runtime.pb.go`。

## 安装

```shell
go install github.com/goapt/grpc-http/cmd/protoc-gen-go-http@latest
```

## 生成

```shell
protoc --proto_path=. \
--proto_path=./third_party \
--go-http_out=paths=source_relative:. \
./example.proto
```

如果使用 `buf`，可参考仓库示例 `example/proto/buf.gen.yaml`：

```yaml
version: v2
plugins:
  - local: protoc-gen-go-http
    out: gen
    opt: paths=source_relative
```

## 生成结果

每个 `go_package` 会生成：

- `*_http.pb.go`：路由注册 + handler stub + `Register<Service>HTTPServer`
- `http_runtime.pb.go`：运行时接口（每个包只生成一份，避免多 service 冲突）
- 其他 pb 文件（由 `protoc-gen-go` / `protoc-gen-go-grpc` 生成）

`http_runtime.pb.go` 中包含两个接口：

- `ServeMux`：用于注册路由（`http.ServeMux` 可直接使用）
- `Codec`：用于请求解码、响应编码和参数校验

## 使用方式

### 1. 定义 proto（含 http 注解）

```proto
service GreeterService {
  rpc SayHello(HelloRequest) returns (HelloReply) {
    option (google.api.http) = {get: "/greeter/hello/{name}"};
  }
}
```

### 2. 实现 service 接口

生成后会有接口 `GreeterServiceHTTPServer`，你只需要实现业务逻辑方法。

### 3. 实现 Codec（示例）

```go
type JSONCodec struct{}

func (JSONCodec) Encode(w http.ResponseWriter, _ *http.Request, v any) {}
func (JSONCodec) Decode(r *http.Request, v any) error { return nil }
func (JSONCodec) Validate(v any) error { return nil }
```

### 4. 注册服务并启动 HTTP Server

```go
mux := http.NewServeMux()

apiv1.RegisterUserServiceHTTPServer(mux, codec, userService)
apiv1.RegisterGreeterServiceHTTPServer(mux, codec, greeterService)

server := &http.Server{
    Addr:    ":8080",
    Handler: mux,
}

if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
    fmt.Printf("Server startup failed: %v\n", err)
}
```

## 路由说明

- 路由 pattern 直接来自 `google.api.http`。
- 当前推荐使用标准库可直接支持的 pattern（例如 `GET /path/{id}`）。
- Query 参数与请求体字段解析行为由你的 `Codec.Decode()` 决定。

## 示例项目

完整可运行示例见：

- `example/proto/demo/v1/user.proto`
- `example/proto/demo/v1/greeter.proto`
- `example/main.go`
- `example/codec/json_codec.go`
- `example/service/*.go`

本地体验：

```shell
make install
cd example/proto && buf generate
cd ../ && go run .
```
