# Stage 1: Build stage
FROM golang:1.26-alpine AS builder

# Install build dependencies
RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o api ./cmd/api/

# Stage 2: Runtime stage
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache ca-certificates sqlite-libs

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/api .

# Create directory for SQLite database
RUN mkdir -p /app/data

# Expose default port
EXPOSE 4000

# Set port env only (no secrets)
ENV PORT=4000

# Run the application
CMD ["./api"]