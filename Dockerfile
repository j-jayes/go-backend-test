# Build stage
FROM golang:1.21 AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main ./cmd/main

# Final stage
FROM alpine:latest

# Install CA certificates to allow HTTPS requests
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary from the previous stage
COPY --from=build /app/main .

# Expose port 8080 to the outside
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
