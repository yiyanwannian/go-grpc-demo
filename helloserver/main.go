package main

import (
	"context"
	"fmt"
	"github.com/albinism/grpc-test/protos/hello"
	"google.golang.org/grpc"
	"log"
	"net"
)

type HelloServiceServer struct {
	hello.UnimplementedHelloServiceServer
}

func (s HelloServiceServer) Hello(ctx context.Context, request *hello.HelloRequest) (*hello.HelloResponse, error) {
	return &hello.HelloResponse{Greeting: fmt.Sprintf("hello %s", request.Name)}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	hello.RegisterHelloServiceServer(grpcServer, &HelloServiceServer{})

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
