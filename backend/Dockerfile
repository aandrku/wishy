# 1. Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /server

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the binary
RUN go build -o server cmd/server/main.go

# 2. Final stage (slim runtime)
FROM alpine:latest

WORKDIR /server

# Copy the compiled binary from builder
COPY --from=builder /server/server .

# Run the binary
CMD ["./server", "-tls", "-port=443"]
