# gRPC examples

## Build Container

```
make image
```

## Run Container

```
docker run --expose 8080 grpc-example:latest
```

Now visit `http://localhost:8080/swagger-ui` to make REST calls interactively.

Alternatively, use curl:

```
curl -X GET "http://localhost:8080/v1/make-box?height=5&width=4&depth=3" -H "accept: application/json"
curl -X GET "http://localhost:8080/v1/status" -H "accept: application/json"
```
