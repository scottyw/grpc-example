## gRPC examples

Compile the protobuf and generate swagger like this:

```
rm -rf factory
mkdir -p factory
protoc \
  -I/usr/local/include \
  -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:factory \
  --swagger_out=logtostderr=true:factory \
  --grpc-gateway_out=logtostderr=true:factory \
  --proto_path proto factory.proto
go generate ./...
 ```

Start the server:

```
go run factoryserver/server.go
```

Then run the client a few times:

```
go run factoryclient/client.go
```

Now visit `http://localhost:8080/swagger-ui`
