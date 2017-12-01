package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/scottyw/grpc-example/boxes"
	"google.golang.org/grpc"
)

type boxFactoryServer struct {
}

func (*boxFactoryServer) CheckOpen(context context.Context, _ *boxes.Empty) (*boxes.BoolValue, error) {
	log.Println("Checking if the factory is open ...")
	return &boxes.BoolValue{Value: true}, nil
}

func (*boxFactoryServer) MakeBox(context context.Context, spec *boxes.Spec) (*boxes.Box, error) {
	log.Println("Making a box ...")
	return &boxes.Box{Volume: spec.Depth * spec.Height * spec.Width}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 5566))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	boxes.RegisterBoxFactoryServer(grpcServer, &boxFactoryServer{})
	log.Println("Ready to make boxes ...")
	grpcServer.Serve(lis)
}
