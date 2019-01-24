## gRPC examples

This repo requires Go 1.11 or later and uses Go modules for dependency management. The repo does not need to be in your Go path.

To enable modules if required:

```
export GO111MODULE=on
```

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

If you make changes to the proto file then run `make` to regenerate all boilerplate code (and install the required tools if necessary):

```
make
```

For the details, have a look in the Makefile.
