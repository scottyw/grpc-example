package main

import (
	"context"
	"log"

	"github.com/scottyw/grpc-example/factory"
)

type factoryServer struct {
}

func (*factoryServer) MakeBox(context context.Context, spec *factory.BoxSpec) (*factory.Box, error) {
	log.Println("Making a box ...")
	return &factory.Box{Volume: spec.Depth * spec.Height * spec.Width}, nil
}

// func (*factoryServer) Status(context context.Context, service *factory.Service) (*factory.StatusMessage, error) {
// 	log.Printf("Checking status for %s ...", service.Name)
// 	return &factory.StatusMessage{
// 		ServiceName: service.Name,
// 		Ok:          true,
// 	}, nil
// }
