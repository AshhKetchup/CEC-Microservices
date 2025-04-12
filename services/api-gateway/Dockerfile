FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy module files first
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o api-gateway ./main.go

FROM alpine:latest
WORKDIR /app

# Copy binary and generated files
COPY --from=builder /app/api-gateway .
COPY --from=builder /app/gen ./gen

EXPOSE 8080
CMD ["./api-gateway"]