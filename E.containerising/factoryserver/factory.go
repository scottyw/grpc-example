package main

import (
	"context"
	"log"

	"github.com/scottyw/grpc-example/E.containerising/factory"
)

type factoryServer struct {
}

func (*factoryServer) MakeBox(context context.Context, spec *factory.BoxSpecification) (*factory.Box, error) {
	log.Println("Making a box ...")
	return &factory.Box{Volume: spec.Depth * spec.Height * spec.Width}, nil
}

func (*factoryServer) Status(context context.Context, service *factory.Empty) (*factory.StatusMessage, error) {
	log.Printf("Checking status ...")
	return &factory.StatusMessage{
		ServiceName: "grpc-example-server",
		Ok:          true,
	}, nil
}
