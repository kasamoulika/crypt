#################################
# Build Stage
#################################
FROM golang:1.17 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY . .

RUN go mod download
RUN go mod verify

# Build the Go application
RUN go build -o crypt

#################################
# Package Stage
#################################
FROM golang:1.17

ENV SRC_DIR=/app
RUN mkdir -p "$SRC_DIR"
WORKDIR $SRC_DIR

COPY --from=builder "$SRC_DIR"/crypt $SRC_DIR

# Expose a port (if your application listens on a specific port)
EXPOSE 8080

# Run the Go application when the container starts
CMD ["./crypt"]