# Stage 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your project
COPY . .

# Install migrate CLI with PostgreSQL support
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest && \
    cp /go/bin/migrate /usr/local/bin/


# Build your Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o activity-bot .

# Stage 2: Run
FROM alpine:latest

WORKDIR /root/

# Install certs (useful for HTTPS if needed)
RUN apk --no-cache add ca-certificates

# Install required tools
RUN apk add --no-cache postgresql-client

# Copy binary from builder
COPY --from=builder /app/activity-bot .
COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate
COPY --from=builder /app/migrations /app/migrations
COPY --from=builder /app/migrations/dbseed /app/migrations/dbseed

# Copy entrypoint
COPY entrypoint.sh /app/entrypoint.sh

# Copy your `.env` file if needed
COPY .env .env

# Run it (this prints QR code to Docker stdout)
# CMD ["sh", "-c", "./activity-bot && migrate -path migrations -database postgres://postgres:postgres1234@postgres:5432/activity_tracker?sslmode=disable up"]
ENTRYPOINT ["/app/entrypoint.sh"]

# CMD ["sh", "-c", "./activity-bot"]
