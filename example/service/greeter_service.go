package service

import (
	"context"
	"fmt"

	apiv1 "github.com/goapt/grpc-http/example/proto/gen/demo/v1"
)

type GreeterService struct{}

func (s *GreeterService) SayHello(_ context.Context, req *apiv1.HelloRequest) (*apiv1.HelloReply, error) {
	name := req.GetName()
	if name == "" {
		name = "world"
	}
	return &apiv1.HelloReply{
		Message: fmt.Sprintf("hello, %s", name),
	}, nil
}
