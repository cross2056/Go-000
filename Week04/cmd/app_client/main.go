package main

import (
	"context"
	"log"
	"os"
	"time"
	"week04/api/hello/v1"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:9000"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	client := hello.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.SayHello(ctx, &hello.HelloRequest{Name: name})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
