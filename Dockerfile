# Use the official Golang image as the builder
FROM golang:1.24-alpine AS builder

# Set environment variables for Go
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the gRPC server
RUN go build -o server cmd/gateway/main.go

# ---- Final Image ----
FROM alpine:3.18

# Add metadata to the image
LABEL maintainer="Your Name bikramgyawali57@gmail.com" \
      version="1.0" \
      description="gRPC server built with Go"

WORKDIR /app

COPY --from=builder /app/server .

# If you have any additional dependencies, install them here
# RUN apk add --no-cache <packages>

# Health check (optional, if you have a health endpoint)
# HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
# CMD curl -f http://localhost:8080/health || exit 1

EXPOSE 8080

CMD ["./server"]
