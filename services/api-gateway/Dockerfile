FROM golang:1.24-alpine

WORKDIR /app

# Copy module files and download dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code.
COPY . .

# Expose the API Gateway port.
EXPOSE 8080

# Run the application.
CMD ["go", "run", "."]
