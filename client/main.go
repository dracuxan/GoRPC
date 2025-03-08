package main

import (
	"fmt"
	"log"

	pb "github.com/dracuxan/GoRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameList{
		Names: []string{"Nisarg", "Niraj", "Jayesh"},
	}

	for {
		opp := 4
		fmt.Println()
		fmt.Println("Enter action to perform:\n1. Unary Stream\n2. Server Stream\n3. Client Stream\n4. Bi Directional Stream\n5. Exit")
		fmt.Print(">> ")
		fmt.Scanf("%d", &opp)
		switch opp {
		case 1:
			callSayHello(client)
			fmt.Println()
		case 2:
			callSayHelloServerStream(client, names)
			fmt.Println()
		case 3:
			callSayHelloClientSideStreaming(client, names)
			fmt.Println()
		case 4:
			callSayHelloBiDiStrem(client, names)
		case 5:
			return
		}
	}
}
