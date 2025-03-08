package main

import (
	"log"
	"time"

	pb "github.com/dracuxan/GoRPC/proto"
)

func (s *helloServer) SayHelloServerStream(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamServer) error {
	log.Printf("Got Request with names: %v", req.Names)

	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}

		time.Sleep(2 * time.Second)
	}

	return nil
}
