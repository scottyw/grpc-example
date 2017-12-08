## gRPC examples

Compile the protobuf like this:

```
mkdir boxes
protoc \
  -I/usr/local/include \
  -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:boxes \
  --swagger_out=logtostderr=true:boxes \
  --grpc-gateway_out=logtostderr=true:boxes \
  --proto_path proto boxes.proto
 ```

### Simple RPC

To try the simple rpc example, start the server:

```
go run factoryserver/server.go
```

Then run the client a few times:

```
go run factoryclient/client.go
```

### Bi-directional streaming

To try the bi-directional streaming example, start the server:

```
go run storeserver/server.go
```

Then run the client:

```
go run storeclient/client.go
```


###

Start the store sevrer and its REST counterpart - both need to be running:
```
go run storeserver/server.go
go run storerestserver/rest.go
```

Now curl the REST API:

```
curl -G  http://localhost:8080/v1/CheckOnline
```