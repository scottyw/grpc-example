## gRPC examples

Start the server:

```
go run factoryserver/*.go
```

Then run the client a few times:

```
go run factoryclient/*.go
```

Now visit `http://localhost:8080/swagger-ui` to make REST calls interactively.

Alternatively, use curl:

```
curl -X GET "http://localhost:8080/v1/make-box?height=5&width=4&depth=3" -H "accept: application/json"
curl -X GET "http://localhost:8080/v1/status" -H "accept: application/json"
```

### Regenerating code after proto file changes

The gRPC and REST bindings plus the swagger file are generated automatically from the proto file. The generated files are committed to the repo so you don't need to run these commands to try the code. 

However, if you make changes to the proto file you'll need to regenerate like this:

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
