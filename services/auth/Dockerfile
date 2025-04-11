# Use a minimal Go base image
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the Go binary
RUN go build -o auth-service ./main.go

# Create a minimal final image
FROM alpine:latest

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/auth-service .

# Expose gRPC port
EXPOSE 50051

COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

CMD ["/wait-for-it.sh", "mysql:3306", "--", "./auth-service"]
