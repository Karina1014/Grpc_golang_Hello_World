package main

import (
	hello "architectural-styles/grpc/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	//gRPC "channel" for destination URI
	conn, err := grpc.NewClient("localhost:9000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//send request to server
	c := hello.NewGreeterClient(conn)
	resp, err := c.SayHello(context.Background(), &hello.HelloRequest{Name: "world"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Message)
}
