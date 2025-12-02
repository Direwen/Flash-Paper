# Stage 1: Build the Go application
FROM golang:1.25-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first (for better caching)
# Docker caches layers, so if these don't change, dependencies won't re-download
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the binary
# -o main: output name
# ./cmd/api: entry point
RUN go build -o main ./cmd/api

FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your app runs on
EXPOSE 8080

# Run the binary
CMD ["./main"]
