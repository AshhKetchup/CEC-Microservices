FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o product-service ./cmd/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/product-service .
COPY config/config.yaml ./config/

EXPOSE 50052
CMD ["./product-service"]