.PHONY: install
install:
	cd ./cmd/protoc-gen-go-http && go install ./...
.PHONY: proto
proto:
	cd example/proto && buf generate

