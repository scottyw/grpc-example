package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sync"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/scottyw/grpc-example/F.adding-firestore/factory"
	"google.golang.org/grpc"
)

func startGRPC(firebaseClient *firestore.Client) {
	lis, err := net.Listen("tcp", "localhost:5566")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	factory.RegisterBoxFactoryServer(grpcServer, &FactoryServer{firebaseClient: firebaseClient})
	log.Println("gRPC server 'ready' ...")
	grpcServer.Serve(lis)
}

func startHTTP() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Connect to the GRPC server
	conn, err := grpc.Dial("localhost:5566", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	// Register grpc-gateway
	rmux := runtime.NewServeMux()
	client := factory.NewBoxFactoryClient(conn)
	err = factory.RegisterBoxFactoryHandlerClient(ctx, rmux, client)
	if err != nil {
		log.Fatal(err)
	}

	// Serve the swagger, swagger-ui and grpc-gateway REST bindings on 8080
	mux := http.NewServeMux()
	mux.Handle("/v1/", rmux)
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("www"))))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("REST server ready on http://0.0.0.0:%s ...", port)
	err = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), mux)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	conf := &firebase.Config{ProjectID: "puppet-sentry-dev"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	go startGRPC(client)

	go startHTTP()

	// Block forever
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()

}
