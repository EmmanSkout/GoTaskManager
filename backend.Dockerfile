# Use an official Go image as the build environment
FROM golang:1.17-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Build the Go binary
RUN go build -o backend-binary .

# Start a new stage for the final image
FROM alpine:3.19.1

# Set the working directory inside the container
WORKDIR /app

# Copy only the built binary from the previous stage
COPY --from=builder /app/backend-binary /app/backend-binary

# Expose the port that the Go backend will run on
EXPOSE 3000

# Command to run the Go backend
CMD ["./backend-binary"]
