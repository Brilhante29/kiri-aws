FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git ca-certificates

# Copy go mod files
COPY go.mod go.sum* ./
RUN go mod download

# Copy source
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /kiro ./cmd/kiro

# Final image
FROM alpine:3.20

RUN apk add --no-cache ca-certificates \
    && addgroup -g 10000 -S kiro && adduser -u 10000 -S kiro -G kiro \
    && mkdir -p /data && chown kiro:kiro /data

COPY --from=builder /kiro /usr/local/bin/kiro

USER kiro

EXPOSE 4566

ENTRYPOINT ["kiro"]
CMD ["--host", "0.0.0.0", "--port", "4566"]
