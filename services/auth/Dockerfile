FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o auth-service ./cmd/main.go

FROM alpine:latest
WORKDIR /app
RUN mkdir -p /app/data  # Create data directory
COPY --from=builder /app/auth-service .
COPY config.yaml .
VOLUME /app/data  # Persist user data
EXPOSE 50051
CMD ["./auth-service"]