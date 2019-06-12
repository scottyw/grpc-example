GRPC_GATEWAY_DIR := $(shell go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway 2> /dev/null)
GO_INSTALLED := $(shell which go)
PROTOC_INSTALLED := $(shell which protoc)
PGGG_INSTALLED := $(shell which protoc-gen-grpc-gateway 2> /dev/null)
PGG_INSTALLED := $(shell which protoc-gen-go 2> /dev/null)

all: build

install-tools:
ifndef PROTOC_INSTALLED
	$(error "go is not installed, please run 'brew install go'")
endif
ifndef PROTOC_INSTALLED
	$(error "protoc is not installed, please run 'brew install protobuf'")
endif
ifndef PGGG_INSTALLED
	@go get -u github.com/grpc-ecosystem/grpc-gateway/...
endif
ifndef PGG_INSTALLED
	@go get -u github.com/golang/protobuf/protoc-gen-go
endif

generate: install-tools
	@rm -rf factory
	@rm -f www/swagger.json
	@mkdir -p factory
	@protoc \
		-I/usr/local/include \
		-I. \
		-I$(GRPC_GATEWAY_DIR)/third_party/googleapis \
		--go_out=plugins=grpc:factory \
		--swagger_out=logtostderr=true:factory \
		--grpc-gateway_out=logtostderr=true:factory \
		--proto_path proto factory.proto
	@cp factory/factory.swagger.json www/swagger.json

build: generate
	@rm -rf bin
	@mkdir -p bin
	@go build -o bin/server factoryserver/*.go
	@go build -o bin/client factoryclient/*.go
	@echo "Success! Binaries can be found in 'bin' dir"
