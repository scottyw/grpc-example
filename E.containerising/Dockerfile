FROM golang:latest as builder

RUN apt update && apt install unzip

WORKDIR /tmp
RUN curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
RUN unzip protoc-3.5.1-linux-x86_64.zip -d protoc3
RUN mv protoc3/bin/* /usr/local/bin/
RUN mv protoc3/include/* /usr/local/include/
RUN rm -rf protoc3

WORKDIR /go
RUN go get github.com/grpc-ecosystem/grpc-gateway/...

ENV GO111MODULE=on

RUN go get github.com/githubnemo/CompileDaemon && cp /go/bin/CompileDaemon /CompileDaemon

RUN mkdir -p /go/src/github.com/scottyw/E.containerising/grpc-example
WORKDIR /go/src/github.com/scottyw/E.containerising/grpc-example

ADD ./ /go/src/github.com/scottyw/E.containerising/grpc-example
RUN make all

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/scottyw/E.containerising/grpc-example/bin/server /app/
COPY --from=builder /go/src/github.com/scottyw/E.containerising/grpc-example/factory/factory.swagger.json /app/www/swagger.json

WORKDIR /app
CMD ["./server"]