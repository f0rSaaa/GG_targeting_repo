FROM golang:1.22-alpine AS builder

# Install git and build dependencies
RUN apk add --no-cache git build-base

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Set GOTOOLCHAIN to local to use the installed Go version
ENV GOTOOLCHAIN=local

# Download and tidy dependencies
RUN go mod tidy

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o targeting-engine

# Final stage
FROM alpine:latest

# Install CA certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/targeting-engine .

# Expose port
EXPOSE 8080

# Set the entry point
CMD ["./targeting-engine"] 