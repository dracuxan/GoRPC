package main

import (
	"context"

	"github.com/dracuxan/GoRPC/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *proto.NoParam) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Message: "Hello Weeb!",
	}, nil
}
