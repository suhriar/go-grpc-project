package main

import (
	"context"
	"log"
	"time"

	pb "github.com/suhriar/go-grpc-project/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServcieClient, timeout time.Duration) {
	log.Println("doGreet was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := pb.GreetRequest{
		FirstName: "Suhri",
	}

	res, err := c.GreetWithDeadline(ctx, &req)
	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Printf("Deadline Exceeded!")
				return
			} else {
				log.Fatalf("unexpected gRPC error: %v\n", err)
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
	}

	log.Printf("GreetWithDeadline: %s\n", res.Result)

}
