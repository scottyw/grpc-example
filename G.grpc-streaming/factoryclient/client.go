package main

import (
	"context"
	"log"
	"time"

	"github.com/scottyw/grpc-example/G.grpc-streaming/factory"

	"google.golang.org/grpc"
)

func main() {
	// Dial the server, waiting up to 3s for a connection (so you can start the client before the server)
	ctx, cancelTimeoutFunc := context.WithTimeout(context.Background(), 3*time.Second)
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

	// Create the gRPC client
	client := factory.NewBoxFactoryClient(conn)
	log.Printf("Created BoxFactoryClient ...")

	// Make a remote call
	streamClient, err := client.StartProductionLine(context.Background(), &factory.BoxSpecification{Height: 2, Width: 3, Depth: 4})
	if err != nil {
		log.Fatalf("Failed to get the stream client: %v", err)
	}
	for {
		box, err := streamClient.Recv()
		if err != nil {
			log.Fatalf("Stream closed")
		}
		log.Printf("Got a lovely box with volume %d", box.Volume)
	}
}
