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
