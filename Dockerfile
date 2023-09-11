# Use the official Go image from the DockerHub
FROM golang:1.21 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main ./cmd/main/

# Use a smaller base image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

# Copy the pre-built binary from the previous stage
COPY --from=builder /app/main /main

# This container exposes port 8080 to the outside world
EXPOSE 8090

# Command to run the executable
CMD ["/main"]
