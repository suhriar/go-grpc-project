package main

import (
	"context"
	"io"
	"log"

	pb "github.com/suhriar/go-grpc-project/greet/proto"
)

func doGreetManyTimes(c pb.GreetServcieClient) {
	log.Println("doGreetmanyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Suhri",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}
