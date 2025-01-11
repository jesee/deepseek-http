# Build stage
FROM golang:1.24rc1-alpine3.20 AS builder
RUN apk add --no-cache git gcc musl-dev

WORKDIR /go/src/app

# Copy go mod files first
COPY go.mod ./
RUN go mod download

COPY . .

# Build with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/app .

# Final stage
FROM alpine

RUN apk --no-cache add ca-certificates && \
    adduser -D -H -u 10001 appuser

WORKDIR /app

COPY --from=builder /go/bin/app .

# Use non-root user
USER appuser

# Use JSON array format for ENTRYPOINT
ENTRYPOINT ["/app/app"]

# Metadata
LABEL Name=deepseekhttp \
      Version=0.0.1 \
      Maintainer="maintainer@example.com"

# Health check
HEALTHCHECK --interval=30s --timeout=3s \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8089/health || exit 1

EXPOSE 8089
