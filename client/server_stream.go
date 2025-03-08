package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/dracuxan/GoRPC/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Streaming has started!")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.SayHelloServerStream(ctx, names)
	if err != nil {
		log.Fatalf("Could not send request: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while Streaming: %v", err)
		}

		log.Printf("%v\n", message)
	}
	log.Println("Streaming finished!")
}
