# gRPC examples

## Build Binaries

```
make all
```

## Run Binaries

Start the server:

```
./bin/server
```

Then run the client a few times:

```
./bin/client
```

Now visit `http://localhost:8080/swagger-ui` to make REST calls interactively.

Alternatively, use curl:

```
curl -X GET "http://localhost:8080/v1/make-box?height=5&width=4&depth=3" -H "accept: application/json"
```

### Regenerating code after proto file changes

The gRPC and REST bindings plus the swagger file are generated automatically from the proto file. The generated files are committed to the repo so you don't need to run these commands to try the code.

However, if you make changes to the proto file you'll need to regenerate like this:

```
make generate
```
