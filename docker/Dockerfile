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
FROM alpine:3.24@sha256:294b683cb724975bec92580e1e685676bd4b50bda910ddb8c51d4cabeaec77e6

RUN apk add --no-cache ca-certificates \
    && addgroup -g 10000 -S kiri && adduser -u 10000 -S kiri -G kiri \
    && mkdir -p /data && chown kiri:kiri /data

COPY --from=builder /kiri /usr/local/bin/kiri

USER kiri

EXPOSE 4566

ENTRYPOINT ["kiri"]
CMD ["--host", "0.0.0.0", "--port", "4566"]
