# ⚡ kiri-aws

[![Go Version](https://img.shields.io/badge/go-1.25+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Docker](https://img.shields.io/badge/docker-ready-blue?logo=docker)](docker-compose.yml)
[![CI](https://img.shields.io/github/actions/workflow/status/Brilhante29/kiri-aws/ci.yml?branch=main)](https://github.com/Brilhante29/kiri-aws/actions)

**kiri-aws** is a next-generation, high-performance local AWS emulator written in Go. Designed for developer productivity and machine-speed automation, **kiri-aws** features a **deterministic Billing Engine** and a **Time Machine API** (`POST /_kiri/time/advance`) that allows developers to simulate hours, days, or months of AWS infrastructure uptime and cost generation in milliseconds — 100% offline.

---

## ✨ Features

- 🚀 **Sub-millisecond Latency:** Instant boot, zero-cloud dependency, written in Go.
- ⏱️ **Time Machine API (`/_kiri/time/advance`):** Advance virtual time by days or months instantly to trigger recurring billing calculations and schedule events.
- 💰 **AWS Cost Explorer (CE) & Budgets Emulation:** Full support for `GetCostAndUsage` requests matching official AWS pricing rates (KMS, S3, DynamoDB, SQS, EC2, Lambda).
- 🔌 **Universal AWS SDK Support:** Works out of the box with AWS SDK Go v2, Boto3 (Python), AWS SDK JS/TS, AWS CLI, and Terraform.
- 🐳 **Docker & Docker Compose Ready:** Light memory footprint for CI/CD pipelines and local microservice dev environments.

---

## ⚡ Quickstart

### 1. Start kiri-aws Server

```bash
# Run locally with Go
go run ./cmd/kiri --port 4566
```

Or run via Docker Compose:

```bash
docker-compose up -d
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
curl -X POST "http://localhost:4566/_kiri/time/advance?days=30"
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

## ## Supported Services (116 services)

<!-- BEGIN SERVICES -->
### Storage
| Service | Description |
|---------|-------------|
| DataSync | Automated data transfer service |
| DynamoDB | NoSQL database |
| DynamoDB Streams | DynamoDB change data capture |
| EBS | Block storage |
| EFS | Elastic file system |
| ElastiCache | In-memory caching |
| FSx | High-performance file systems |
| Glacier | Archive storage |
| MemoryDB | Redis-compatible database |
| S3 | Object storage |
| S3 Control | S3 account-level operations |
| S3 Tables | S3 table buckets |
| Storage Gateway | Hybrid cloud storage connection |

### Compute
| Service | Description |
|---------|-------------|
| App Runner | Containerized web application service |
| Auto Scaling | EC2 auto scaling management |
| Batch | Batch computing |
| EC2 | Virtual machines |
| Elastic Beanstalk | Application deployment |
| Lambda | Serverless functions |
| Lightsail | Virtual private server service |
| Serverless Application Repository | Serverless application catalog |

### Container
| Service | Description |
|---------|-------------|
| ECR | Container registry |
| ECS | Container orchestration |
| EKS | Kubernetes service |

### Database
| Service | Description |
|---------|-------------|
| DocumentDB | MongoDB-compatible database |
| Keyspaces | Apache Cassandra-compatible database |
| Neptune | Graph database |
| QLDB | Quantum Ledger Database |
| RDS | Relational database service |
| Redshift | Data warehousing |
| Timestream | Time-series database service |

### Messaging & Integration
| Service | Description |
|---------|-------------|
| EventBridge | Event bus |
| Firehose | Data delivery |
| IoT Core | IoT device connection and messaging |
| IoT Data Plane | Real-time IoT message broker |
| Kinesis | Real-time streaming |
| MQ | Message broker (ActiveMQ/RabbitMQ) |
| MSK (Kafka) | Managed streaming for Kafka |
| Pipes | Event-driven integration |
| SNS | Pub/Sub messaging |
| SQS | Message queuing |

### Security & Identity
| Service | Description |
|---------|-------------|
| ACM | Certificate management |
| Cognito | User authentication |
| GuardDuty | Threat detection service |
| IAM | Identity and access management |
| Inspector v2 | Automated vulnerability management |
| KMS | Key management |
| Macie | Data security and privacy |
| Resource Access Manager | Cross-account resource sharing |
| STS | Security token service |
| Secrets Manager | Secret storage |
| Security Lake | Security data lake |
| Signer | Code signing service |
| WAFv2 | Web application firewall v2 |

### Monitoring & Logging
| Service | Description |
|---------|-------------|
| CloudTrail | API audit logging |
| CloudWatch | Metrics and alarms |
| CloudWatch Logs | Log management |
| CloudWatch Synthetics | Synthetic canary monitoring |
| X-Ray | Distributed tracing |

### Networking & Content Delivery
| Service | Description |
|---------|-------------|
| API Gateway | API management (REST API) |
| API Gateway v2 | API management (HTTP/WebSocket API) |
| App Mesh | Service mesh |
| CloudFront | CDN |
| ELBv2 | Load balancing |
| Global Accelerator | Network acceleration |
| Location | Location-based services |
| Route 53 | DNS service |
| Route 53 Resolver | DNS resolver |

### Application Integration
| Service | Description |
|---------|-------------|
| Amplify | Full-stack application hosting |
| AppSync | GraphQL API |
| Pinpoint SMS Voice v2 | SMS messaging |
| SES | Email service |
| SES v2 | Email service (v2 API) |
| Scheduler | Task scheduling |
| Step Functions | Workflow orchestration |

### Management & Configuration
| Service | Description |
|---------|-------------|
| AWS Health | Service status and health events |
| Backup | Centralized backup service |
| Budgets | AWS budget tracking and alerts |
| Cloud Control API | Unified CRUD API for cloud resources |
| CloudFormation | Infrastructure as code |
| CodeConnections | Source code connections |
| Config | Resource configuration |
| Organizations | Multi-account management |
| SSM | Systems Manager |
| Service Quotas | Service limit management |

### Analytics & ML
| Service | Description |
|---------|-------------|
| Athena | SQL query service |
| Bedrock | Generative AI foundation models and runtime |
| Comprehend | NLP service |
| Data Exchange | Data marketplace |
| EMR | Big data processing platform |
| Entity Resolution | Entity matching |
| Forecast | Time-series forecasting |
| Glue | ETL service |
| Kendra | Intelligent enterprise search |
| Kinesis Video Streams | Video streaming and analytics |
| OpenSearch | Search and analytics suite |
| Polly | Text-to-speech synthesis |
| QuickSight | Business intelligence service |
| Rekognition | Image/video analysis |
| SageMaker | Machine learning |
| Textract | Document text and data extraction |
| Transcribe | Speech-to-text recognition |
| Translate | Neural machine translation |

### Developer Tools
| Service | Description |
|---------|-------------|
| CodeArtifact | Artifact and package repository |
| CodeBuild | Continuous integration build service |
| CodeCommit | Source control repository service |
| CodeGuru Profiler | Application profiling |
| CodeGuru Reviewer | Automated code review |
| CodePipeline | Continuous delivery workflow service |

### Other Services
| Service | Description |
|---------|-------------|
| Cost Explorer | Cost analysis |
| DLM | Data lifecycle manager |
| Directory Service | Microsoft AD |
| EMR Serverless | Big data processing |
| FinSpace | Financial data management |
| GameLift | Game server hosting |
| Resilience Hub | Application resilience |
<!-- END SERVICES -->

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

## 📜 Upstream Attribution & License

`kiri-aws` is licensed under the [MIT License](LICENSE).

### Upstream Open-Source Attribution
`kiri-aws` was built upon and evolved from the open-source AWS emulation engine foundation created by **sivchari/kumo** (and awsim). We express our sincere appreciation to the original authors and open-source contributors for their foundational work.

Copyright (c) 2025 sivchari  
Copyright (c) 2026 Kiro-AWS / Kiri-AWS Contributors
