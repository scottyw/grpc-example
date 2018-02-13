package main

import (
	"context"
	"log"
	"time"

	"github.com/scottyw/grpc-example/boxes"
	"google.golang.org/grpc"
)

func main() {
	// Wait up to 3000ms for a server connection e.g. try starting the client before the server
	ctx, cancelTimeoutFunc := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	conn, err := grpc.DialContext(ctx, "localhost:5566",
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	cancelTimeoutFunc()
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	log.Printf("Dialled OK ...")
	defer conn.Close()
	client := boxes.NewBoxFactoryClient(conn)
	log.Printf("Created BoxFactoryClient ...")
	box, err := client.MakeBox(context.Background(), &boxes.Spec{Height: 2, Width: 3, Depth: 4})
	if err != nil {
		log.Fatalf("Failed to make a box: %v", err)
	}
	log.Printf("Got a lovely box with volume %d", box.Volume)
}
