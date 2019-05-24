package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/scottyw/grpc-example/F.adding-firestore/factory"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FactoryServer struct {
	firebaseClient *firestore.Client
}

func (f *FactoryServer) MakeBox(ctx context.Context, spec *factory.BoxSpecification) (*factory.Empty, error) {
	log.Println("Making a box ...")
	_, _, err := f.firebaseClient.Collection("boxes").Add(ctx, map[string]interface{}{
		"name":   spec.Name,
		"width":  spec.Width,
		"depth":  spec.Depth,
		"height": spec.Height,
	})
	if err != nil {
		return &factory.Empty{}, status.Error(codes.Internal, err.Error())
	}
	log.Printf("box created ...")
	return &factory.Empty{}, nil
}

func (f *FactoryServer) GetBoxes(ctx context.Context, spec *factory.Empty) (*factory.Boxes, error) {
	output := factory.Boxes{}
	iter := f.firebaseClient.Collection("boxes").Documents(ctx)
	log.Printf("getting boxes ...")
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		var box factory.BoxSpecification
		doc.DataTo(&box)

		output.Boxes = append(output.Boxes, &box)
	}
	log.Printf("done getting boxes ...")
	return &output, nil
}

func (f *FactoryServer) Status(ctx context.Context, service *factory.Empty) (*factory.StatusMessage, error) {
	log.Printf("Checking status ...")
	return &factory.StatusMessage{
		ServiceName: "grpc-example-server",
		Ok:          true,
	}, nil
}
