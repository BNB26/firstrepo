# syntax=docker/dockerfile:1

# Stage 1: Build the Go application
FROM golang:1.21 AS builder
WORKDIR /app

# Copy go.mod and go.sum first to leverage caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire application source
COPY . .

# Build the application
RUN go build -o main .

# Stage 2: Create a minimal runtime image
FROM gcr.io/distroless/base-debian12
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose the application port (change if needed)
EXPOSE 8088

# Run the binary
CMD ["/app/main"]