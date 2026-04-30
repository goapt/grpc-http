package main

import (
	"log"
	"net/http"

	"github.com/goapt/grpc-http/example/codec"
	apiv1 "github.com/goapt/grpc-http/example/proto/gen/user/v1"
	"github.com/goapt/grpc-http/example/service"
)

func main() {
	mux := http.NewServeMux()
	apiv1.RegisterUserServiceHTTPServer(mux, codec.JSONCodec{}, &service.UserService{})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("HTTP server listening on :8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server startup failed: %v", err)
	}
}
