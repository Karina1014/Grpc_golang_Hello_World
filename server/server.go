package main

import (
	hello "architectural-styles/grpc/proto"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	hello.UnimplementedGreeterServer
}

// send the message to client
func (s *server) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	return &hello.HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

// port
func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		(panic(err))
	}
	s := grpc.NewServer()
	hello.RegisterGreeterServer(s, &server{})
	fmt.Println("Port Listening on: 9000")
	s.Serve(lis)
}
