# Build the Go application
FROM golang:1.21 AS go-build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main ./cmd/main

# Build the Svelte frontend
FROM node:16 AS svelte-build
WORKDIR /app/svelte-frontend
COPY svelte-frontend/package*.json ./
RUN npm install
COPY svelte-frontend/ ./
RUN npm run build

# Final stage
FROM alpine:latest

# Install CA certificates to allow HTTPS requests
RUN apk --no-cache add ca-certificates

# Copy the Go binary from go-build stage
COPY --from=go-build /app/main /root/main

# Copy the Svelte built assets from svelte-build stage
COPY --from=svelte-build /app/svelte-frontend/public/build /root/static

# Set the Current Working Directory inside the container
WORKDIR /root/

# Expose port 8080 to the outside
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
