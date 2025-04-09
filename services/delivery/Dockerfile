FROM golang:1.21-alpine AS builder
WORKDIR /app

RUN apk add protobuf protobuf-dev && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY services/delivery/proto/ ./proto/
COPY services/order/proto/order.proto ./proto/external/

RUN mkdir -p gen && \
    protoc -I proto --go_out=gen --go-grpc_out=gen \
    proto/delivery.proto proto/external/order.proto

COPY services/delivery/ .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o delivery-service ./cmd/main.go

FROM alpine:latest
WORKDIR /app
RUN mkdir -p /app/data
COPY --from=builder /app/delivery-service .
COPY --from=builder /app/gen ./gen
COPY services/delivery/config/config.yaml ./config/
VOLUME /app/data
EXPOSE 50054
CMD ["./delivery-service"]