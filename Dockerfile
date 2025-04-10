# Use the official golang image as a parent image
FROM golang:1.24-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
# COPY go.mod go.sum ./

# Download all dependencies
# RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the app listens on
EXPOSE 8080

# Run the executable
CMD ["./main"]