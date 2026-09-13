# Build Stage
FROM golang:alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
# CGO_ENABLED=0 is required for Alpine to build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /api-app ./cmd/main.go

# Final Stage (using a very small image)
FROM alpine:latest

WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /api-app .

# Expose port (must match your APP_PORT in .env, usually 8080)
EXPOSE 8080

# Command to run the executable
CMD ["./api-app"]
