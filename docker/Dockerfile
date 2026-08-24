FROM golang:1.27-alpine@sha256:4c9fe60190a2a3350ddc51de80d0224b8a6698d12bdfc999fee45ea9d6c46dbc AS builder

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
FROM alpine:3.24@sha256:28bd5fe8b56d1bd048e5babf5b10710ebe0bae67db86916198a6eec434943f8b

RUN apk add --no-cache ca-certificates \
    && addgroup -g 10000 -S kiri && adduser -u 10000 -S kiri -G kiri \
    && mkdir -p /data && chown kiri:kiri /data

COPY --from=builder /kiri /usr/local/bin/kiri

USER kiri

EXPOSE 4566

ENTRYPOINT ["kiri"]
CMD ["--host", "0.0.0.0", "--port", "4566"]
