FROM golang:1.19.0

# Set the working directory to /app
WORKDIR /app

# Copy the source code into the container
COPY . .

# Download all dependencies
RUN go mod download

# Build the Go application
RUN go build -o abax-api

EXPOSE ${HTTP_PORT}

# Set the container command
CMD ["./abax-api"]
