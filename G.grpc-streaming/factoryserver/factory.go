package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/scottyw/grpc-example/G.grpc-streaming/factory"
)

type factoryServer struct {
}

func (*factoryServer) MakeBox(context context.Context, spec *factory.BoxSpecification) (*factory.Box, error) {
	log.Println("Making a box ...")
	return &factory.Box{Volume: spec.Depth * spec.Height * spec.Width}, nil
}

func (*factoryServer) StartProductionLine(spec *factory.BoxSpecification, server factory.BoxFactory_StartProductionLineServer) error {
	log.Println("Starting production line ...")

	for {
		volume := int32(rand.Intn(1000))
		log.Printf("Making a box with volumne %d ...\n", volume)
		box := &factory.Box{Volume: volume}
		err := server.Send(box)
		if err != nil {
			panic(err)
		}
		time.Sleep(2 * time.Second)
	}

}
