package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/scottyw/grpc-example/boxes"
	"google.golang.org/grpc"
)

type storeFrontServer struct {
}

func (*storeFrontServer) CheckOnline(context context.Context, _ *boxes.Empty) (*boxes.BoolValue, error) {
	log.Println("Checking if the store is online ...")
	return &boxes.BoolValue{Value: true}, nil
}

func (*storeFrontServer) PlaceOrder(stream boxes.StoreFront_PlaceOrderServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Order %d: Need a box of size %dx%dx%d", in.OrderNumber, in.Spec.Height, in.Spec.Width, in.Spec.Depth)
		time.Sleep(time.Duration(rand.Int31n(5000)) * time.Millisecond)
		volume := in.Spec.Height * in.Spec.Width * in.Spec.Depth
		log.Printf("Order %d: Box with volume %d is ready for delivery!", in.OrderNumber, volume)
		stream.Send(&boxes.Delivery{OrderNumber: in.OrderNumber, Box: &boxes.Box{Volume: volume}})
	}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 5577))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	boxes.RegisterStoreFrontServer(grpcServer, &storeFrontServer{})
	log.Println("Ready to sell some boxes ...")
	grpcServer.Serve(lis)
}
