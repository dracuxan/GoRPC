package main

import (
	"io"
	"log"

	"github.com/dracuxan/GoRPC/proto"
)

func (s *helloServer) SayHelloClientStream(stream proto.GreetService_SayHelloClientStreamServer) error {
	var message []string
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		log.Printf("Got Request with name: %v", req.Name)

		message = append(message, "Hello "+req.Name)
	}
}
