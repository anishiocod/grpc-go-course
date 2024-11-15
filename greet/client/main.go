package main

import (
	"log"

	pb "github.com/anishiocod/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect on:%v\n", err)

	}
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)
	doGreet(c)

}
