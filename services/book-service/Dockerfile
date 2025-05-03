FROM golang:1.21 AS builder

WORKDIR /app

# Copy go.mod and go.sum for dependency management
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o book-service

# Use a minimal image for production
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/book-service .

# Expose the service port
EXPOSE 5000

# Start the application
CMD ["./book-service"]