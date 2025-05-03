# Use official Golang image
FROM golang:1.21 AS builder

# Set working directory
WORKDIR /app

# Copy Go modules files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o borrow-service

# Use minimal image for production
FROM alpine:latest

WORKDIR /root/

# Copy the compiled binary from builder stage
COPY --from=builder /app/borrow-service .

# Expose the service port
EXPOSE 8080

# Set environment variables for DB connection
ENV DB_HOST=borrow-db
ENV DB_USER=user
ENV DB_PASSWORD=password
ENV DB_NAME=borrow_service

# Start the service
CMD ["./borrow-service"]
