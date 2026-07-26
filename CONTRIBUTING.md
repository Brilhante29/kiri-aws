# Contributing to `kiri-aws`

Thank you for your interest in contributing to `kiri-aws`!

## Development Setup

### Prerequisites
- **Go:** 1.25+
- **Docker:** Required for running containerized builds & integration tests

### Building Locally

```bash
# Build kiri binary
make build

# Run unit tests
make test
```

### Running in Docker

```bash
docker build -f docker/Dockerfile -t kiri-aws:local .
docker run -p 4566:4566 kiri-aws:local
```

## Pull Request Guidelines

1. Ensure code is formatted with `gofmt` and passes `go test ./...`.
2. Keep commits concise and descriptive using Conventional Commits (`feat:`, `fix:`, `docs:`, `refactor:`).
3. Update relevant documentation or tests when modifying service handlers.
