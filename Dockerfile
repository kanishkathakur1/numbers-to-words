# Pull lightweight golang container, to build the app
FROM golang:1.24.4-alpine AS builder

WORKDIR /app

# Copy all the required files
COPY go.mod .

COPY src/ .

# Run tests before building the binary, and fail if any issues
RUN go test -v ./...

RUN CGO_ENABLED=0 go build -o /app/numbers-to-words .

# Use an alpine container to run the CLI binary
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder
COPY --from=builder /app/numbers-to-words .

# Run the binary, on entrypoint
ENTRYPOINT [ "./numbers-to-words" ]

# Default to no args
CMD []

