FROM golang:1.27-alpine@sha256:8a5910f31396cd4d89662f56c68b3ae31d374308270a1c3bd96672ee5ed43414 AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git ca-certificates

# Copy go mod files
COPY go.mod go.sum* ./
RUN go mod download

# Copy source
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /kiri ./cmd/kiri

# Final image
FROM alpine:3.24@sha256:294b683cb724975bec92580e1e685676bd4b50bda910ddb8c51d4cabeaec77e6

RUN apk add --no-cache ca-certificates \
    && addgroup -g 10000 -S kiri && adduser -u 10000 -S kiri -G kiri \
    && mkdir -p /data && chown kiri:kiri /data

COPY --from=builder /kiri /usr/local/bin/kiri

USER kiri

EXPOSE 4566

ENTRYPOINT ["kiri"]
CMD ["--host", "0.0.0.0", "--port", "4566"]
