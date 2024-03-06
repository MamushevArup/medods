# Stage 1: Build the Go application
FROM golang:1.21 as builder

WORKDIR /auth

COPY go.mod go.sum ./

RUN go mod download

COPY .env .env

COPY config.yml config.yml

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /auth/cmd cmd/main.go

# Stage 2: Build the final image
FROM scratch

WORKDIR /app

# Copy the binary from the first stage
COPY --from=builder /auth/cmd /app/cmd

COPY --from=builder /auth/.env /app/

COPY --from=builder /auth/config.yml /app/
# Command to run the executable
CMD ["/app/cmd/main"]