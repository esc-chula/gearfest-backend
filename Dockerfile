FROM golang:1.21.5-alpine3.18

# Working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o server ./src/.

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./server"]