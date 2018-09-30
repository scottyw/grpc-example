package main

import (
	"log"
	"net"
	"sync"

	"github.com/scottyw/grpc-example/factory"
	"google.golang.org/grpc"
)

func startGRPC() {
	lis, err := net.Listen("tcp", "localhost:5566")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	factory.RegisterBoxFactoryServer(grpcServer, &factoryServer{})
	log.Println("gRPC server ready...")
	grpcServer.Serve(lis)
}

func main() {

	go startGRPC()

	// Block forever
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()

}
