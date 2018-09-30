package main

import (
	"context"
	"log"

	"github.com/scottyw/grpc-example/factory"
)

type factoryServer struct {
}

func (*factoryServer) MakeBox(context context.Context, spec *factory.BoxSpecification) (*factory.Box, error) {
	log.Println("Making a box ...")
	return &factory.Box{Volume: spec.Depth * spec.Height * spec.Width}, nil
}
