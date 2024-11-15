package main

import (
	"context"
	"log"

	pb "github.com/anishiocod/grpc-go-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet Was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Clement",
	})
	if err != nil {
		log.Fatalf("Could not greet %v\n", err)
	}
	log.Printf("Greeting: %s\n", res.Result)

}
