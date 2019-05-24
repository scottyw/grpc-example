FROM puppet/gogrpc:latest AS builder

ENV GO111MODULE=on

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