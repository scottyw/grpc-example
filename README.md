## gRPC examples

Compile the protobuf like this:

```
mkdir boxes
protoc --go_out=plugins=grpc:boxes --proto_path proto boxes.proto
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
