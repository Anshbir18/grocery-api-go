# Start from Go image
FROM golang:1.22-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and download deps
COPY go.mod ./
RUN go mod download

# Copy everything
COPY . .

# Build the Go binary
RUN go build -o main ./cmd/main.go

# Run it
CMD ["./main"]