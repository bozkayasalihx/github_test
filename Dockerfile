# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum* ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main

# Final stage
FROM alpine:3.19

WORKDIR /app

# Add non root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Copy binary from builder
COPY --from=builder /app/main .

# Use non root user
USER appuser

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./main"]
