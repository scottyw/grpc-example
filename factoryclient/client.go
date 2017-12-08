package main

import (
	"context"
	"log"

	"github.com/scottyw/grpc-example/boxes"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5566", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := generated.NewBoxFactoryClient(conn)
	box, err := client.MakeBox(context.Background(), &boxes.Spec{Height: 2, Width: 3, Depth: 4})
	if err != nil {
		log.Fatalf("failed to make a box: %v", err)
	}
	log.Printf("Got a lovely box with volume %d", box.Volume)
}
