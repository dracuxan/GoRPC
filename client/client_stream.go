package main

import (
	"context"
	"log"
	"time"

	pb "github.com/dracuxan/GoRPC/proto"
)

func callSayHelloClientSideStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Streaming started!")

	stream, err := client.SayHelloClientStream(context.Background())
	if err != nil {
		log.Fatalln("error while connecting:", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending: %v", err)
		}

		log.Printf("Sent request with name: %v", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving: %v", err)
	}
	log.Printf("Streaming completed")
	log.Printf("%v", res.Messages)
}
