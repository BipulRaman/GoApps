# Stage 1: Build the Go application
FROM golang:1.20-alpine AS builder

# Install necessary build tools
RUN apk add --no-cache git build-base

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application with static linking for better performance
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Stage 2: Create a minimal runtime image
FROM alpine:3.18

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/main .

# Expose the port the app runs on
EXPOSE 8081

# Command to run the executable
CMD ["./main"]