# Use official Golang Alpine image
FROM golang:1.24-alpine

# Set working directory
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy all files from current directory
COPY . .


EXPOSE 50053
# Run the application
CMD ["go", "run", "."]