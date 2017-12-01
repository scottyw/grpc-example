package main

import (
	"context"
	"io"
	"log"
	"math/rand"

	"github.com/scottyw/grpc-example/boxes"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5577", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := boxes.NewStoreFrontClient(conn)
	stream, err := client.PlaceOrder(context.Background())
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Received order %d which is a box with a volume of %d", in.OrderNumber, in.Box.Volume)
		}
	}()
	for i := int32(0); i < 10; i++ {
		spec := &boxes.Spec{Height: rand.Int31n(100), Width: rand.Int31n(100), Depth: rand.Int31n(100)}
		log.Printf("Sending Order %d for a box of size %dx%dx%d", i, spec.Height, spec.Width, spec.Depth)
		if err := stream.Send(&boxes.Order{OrderNumber: i, Spec: spec}); err != nil {
			log.Fatalf("Failed to send: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc
}
