GRPC_GATEWAY_DIR := $(shell go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway 2> /dev/null)
GO_INSTALLED := $(shell which go)
PROTOC_INSTALLED := $(shell which protoc)
BINDATA_INSTALLED := $(shell which go-bindata 2> /dev/null)
PGGG_INSTALLED := $(shell which protoc-gen-grpc-gateway 2> /dev/null)
PGS_INSTALLED := $(shell which protoc-gen-swagger 2> /dev/null)
PGG_INSTALLED := $(shell which protoc-gen-go 2> /dev/null)

all: build

install-tools:
ifndef PROTOC_INSTALLED
	$(error "protoc is not installed, please run 'brew install protobuf'")
endif

PROTOBUF_INCLUDES += -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
PROTOBUF_INCLUDES += -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/
PROTOBUF_INCLUDES += -I${GOPATH}/src/

generate: install-tools
	@mkdir -p factory
	@protoc \
		-I/usr/local/include \
		-I. \
		${PROTOBUF_INCLUDES} \
		--go_out=plugins=grpc:factory \
		--swagger_out=logtostderr=true:factory \
		--grpc-gateway_out=logtostderr=true:factory \
		proto/factory.proto

build: generate
	@rm -rf bin
	@mkdir -p bin
	@CGO_ENABLED=0 GOOS=linux go build -o bin/server factoryserver/*.go
	@echo "Success! Binaries can be found in 'bin' dir"
