FROM golang:1.21-alpine AS builder
WORKDIR /app

# Install protoc and dependencies
RUN apk add protobuf protobuf-dev && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Copy and compile protos
COPY services/product/proto/ ./proto/
RUN mkdir -p gen && \
    protoc -I proto --go_out=gen --go-grpc_out=gen \
    proto/product.proto proto/common/types.proto

# Build service
COPY services/product/ .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o product-service ./cmd/main.go

FROM alpine:latest
WORKDIR /app
RUN mkdir -p /app/data
COPY --from=builder /app/product-service .
COPY --from=builder /app/gen ./gen
COPY services/product/config/config.yaml ./config/
VOLUME /app/data
EXPOSE 50052
CMD ["./product-service"]