package main

import (
	"context"

	pb "github.com/dracuxan/GoRPC/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello Weeb!",
	}, nil
}
