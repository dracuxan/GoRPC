package main

import (
	"context"
	"log"
	"time"

	pb "github.com/dracuxan/GoRPC/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	log.Printf("%v", res.Message)
}
