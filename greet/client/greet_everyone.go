package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/suhriar/go-grpc-project/greet/proto"
)

func doGreetEveryone(c pb.GreetServcieClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Suhri"},
		{FirstName: "Ainur"},
		{FirstName: "Rifky"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("rror while receiving: %v\n", err)
			}

			log.Printf("Receiving: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
