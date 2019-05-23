FROM puppet/gogrpc:latest AS builder

ENV GO111MODULE=on

RUN mkdir -p /go/src/github.com/scottyw/grpc-example
WORKDIR /go/src/github.com/scottyw/grpc-example

ADD ./ /go/src/github.com/scottyw/grpc-example
RUN make all

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/scottyw/grpc-example/bin/server /app/
COPY --from=builder /go/src/github.com/scottyw/grpc-example/factory/proto/factory.swagger.json /app/www/swagger.json

WORKDIR /app
CMD ["./server"]