package service

import (
	"context"
	"log"
	"time"

	apiv1 "github.com/goapt/grpc-http/example/proto/gen/user/v1"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type UserService struct{}

func (s *UserService) Create(_ context.Context, req *apiv1.CreateRequest) (*emptypb.Empty, error) {
	log.Printf("Create called: name=%q email=%q", req.GetName(), req.GetEmail())
	return &emptypb.Empty{}, nil
}

func (s *UserService) HTML(_ context.Context, _ *emptypb.Empty) (*httpbody.HttpBody, error) {
	return &httpbody.HttpBody{
		ContentType: "text/html; charset=utf-8",
		Data:        []byte("<html><body><h1>Hello UserServiceHTTPServer</h1></body></html>"),
	}, nil
}

func (s *UserService) List(_ context.Context, req *apiv1.ListRequest) (*apiv1.ListResponse, error) {
	now := timestamppb.New(time.Now())
	log.Printf("List called: page=%d", req.GetPage())
	return &apiv1.ListResponse{
		Results: []*apiv1.User{
			{
				Id:        1,
				Name:      "Alice",
				Email:     "alice@example.com",
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
	}, nil
}

func (s *UserService) Put(_ context.Context, req *apiv1.ListRequest) (*emptypb.Empty, error) {
	log.Printf("Put called: page=%d", req.GetPage())
	return &emptypb.Empty{}, nil
}
