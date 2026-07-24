# ⚡ Kiro-AWS

[![Go Version](https://img.shields.io/badge/go-1.25+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Docker](https://img.shields.io/badge/docker-ready-blue?logo=docker)](docker-compose.yml)
[![Docs](https://img.shields.io/badge/docs-Lightpanda_Style-8b5cf6)](docs/index.html)

**Kiro-AWS** is an ultra-fast, local AWS emulator built for developer joy and machine speed. It features a **deterministic Billing Engine** and a **Time Machine API** (`POST /_kiro/time/advance`) that allows developers to simulate hours, days, or months of infrastructure uptime and cost generation in milliseconds — 100% offline.

---

## ✨ Features

- 🚀 **Sub-millisecond Latency:** Instant boot, zero-cloud dependency, written in Go.
- ⏱️ **Time Machine API (`/_kiro/time/advance`):** Advance virtual time by days or months instantly to trigger recurring billing calculations.
- 💰 **AWS Cost Explorer (CE) Emulation:** Full support for `GetCostAndUsage` requests matching official AWS pricing rates (KMS, S3, DynamoDB, SQS, EC2, Lambda).
- 🔌 **Universal AWS SDK Support:** Works out of the box with AWS SDK Go v2, Boto3 (Python), AWS SDK JS/TS, AWS CLI, and Terraform.
- 🐳 **Docker & Docker Compose Ready:** Light memory footprint for CI/CD pipelines and local microservice dev.

---

## ⚡ Quickstart (Run in 10 Seconds)

### 1. Start Kiro-AWS Server

```bash
# Clone & run locally
go run ./cmd/kiro --port 4566
```

Or run via Docker:

```bash
docker run -p 4566:4566 kiro-aws:latest
```

### 2. Configure Environment & AWS CLI

```bash
export AWS_ACCESS_KEY_ID="test"
export AWS_SECRET_ACCESS_KEY="test"
export AWS_DEFAULT_REGION="us-east-1"
export AWS_ENDPOINT_URL="http://localhost:4566"

# Create S3 Bucket
aws s3 mb s3://my-billing-bucket

# Create DynamoDB Table
aws dynamodb create-table \
  --table-name Orders \
  --attribute-definitions AttributeName=ID,AttributeType=S \
  --key-schema AttributeName=ID,KeyType=HASH \
  --billing-mode PAY_PER_REQUEST
```

---

## ⏱️ Time Machine API (Time Travel Simulation)

Simulate 30 days of resource uptime instantly:

```bash
curl -X POST "http://localhost:4566/_kiro/time/advance?days=30"
```

**Response:**

```json
{
  "status": "success",
  "advanced_by": "720h0m0s",
  "current_virtual_time": "2026-08-21T05:13:21Z"
}
```

---

## 💰 Query Cost Explorer (Simulated AWS Invoice)

```bash
curl -X POST "http://localhost:4566/" \
  -H "Content-Type: application/x-amz-json-1.1" \
  -H "X-Amz-Target: AWSInsightsIndexService.GetCostAndUsage" \
  -d '{
    "TimePeriod": {"Start": "2026-01-01", "End": "2026-12-31"},
    "Granularity": "MONTHLY",
    "Metrics": ["UnblendedCost"]
  }'
```

---

## 💻 SDK Examples

### Go (AWS SDK v2)

```go
package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	ctx := context.Background()
	cfg, _ := config.LoadDefaultConfig(ctx,
		config.WithRegion("us-east-1"),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:4566"}, nil
			},
		)),
	)

	client := s3.NewFromConfig(cfg)
	out, _ := client.ListBuckets(ctx, &s3.ListBucketsInput{})
	fmt.Printf("Buckets count: %d\n", len(out.Buckets))
}
```

---

## 📚 Documentation Site

For full interactive docs, API reference, and copy-paste code snippets, open `docs/index.html` in your browser.

---

## 📄 License

MIT License &copy; 2026 Kiro-AWS Contributors.
