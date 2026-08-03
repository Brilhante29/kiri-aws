<div align="center">

# kiri-aws 霧

**The cloud, brought down to your machine.**

[![CI](https://img.shields.io/github/actions/workflow/status/Brilhante29/kiri-aws/ci.yml?branch=main&label=CI&logo=github)](https://github.com/Brilhante29/kiri-aws/actions)
[![Release](https://img.shields.io/github/v/release/Brilhante29/kiri-aws?logo=github&label=release&sort=semver)](https://github.com/Brilhante29/kiri-aws/releases/latest)
[![Go](https://img.shields.io/badge/go-1.25%2B-00ADD8?logo=go&logoColor=white)](https://go.dev)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/Brilhante29/kiri-aws/badge)](https://securityscorecards.dev/viewer/?uri=github.com/Brilhante29/kiri-aws)
[![Go Report Card](https://goreportcard.com/badge/github.com/Brilhante29/kiri-aws)](https://goreportcard.com/report/github.com/Brilhante29/kiri-aws)

A single binary that emulates **116 AWS services** on one local endpoint.
Real client compatible, offline, free. Point your Go, Python, Node, or Java SDK at it.
Point the AWS CLI, Terraform, or plain HTTP at it. Build, test, and price a whole AWS
architecture without an account, a credential, or a bill.

`kiri` (霧) is fog: that same cloud at ground level, running locally on your laptop.

**Part of the kiri family** — [**kiri-gcp**](https://github.com/Brilhante29/kiri-gcp)
does the same for Google Cloud, with the same CLI shape, the same `KIRI_*`
configuration, and the same release guarantees.

</div>

---

## Why kiri

Cloud test environments are slow, shared, and billed. Test suites that touch S3,
DynamoDB, or SQS end up mocked into meaninglessness, or they run against a real
account and become flaky and expensive. `kiri-aws` gives you the real client
protocol on `localhost`: your SDK calls are unchanged, the wire format is the
real one, and the whole surface starts in a container in under a second.

It also answers a question mocks cannot: **what will this architecture cost?**
A deterministic billing engine tracks resource usage, and the Time Machine API
fast-forwards the clock so a month of accrual happens in milliseconds.

- **116 AWS services in one binary** — S3, DynamoDB, SQS, SNS, Lambda, EC2, IAM,
  KMS, Step Functions, EventBridge, and [more](#supported-services).
- **No credentials, no account.** Any access key works; nothing leaves the machine.
- **Works with the tools you already use** — AWS SDKs (Go, Python, JS, Java), the
  AWS CLI, and Terraform, by pointing them at one endpoint.
- **Cost surface** — Cost Explorer `GetCostAndUsage` and Budgets, backed by a
  pricing catalog.
- **Time Machine** — `POST /_kiri/time/advance` moves the virtual clock so
  recurring billing and scheduled events fire on demand.
- **Optional persistence** — set `KIRI_DATA_DIR` and state survives restarts.

---

## Install

Every release ships signed binaries for linux, macOS, and Windows (amd64 and
arm64), a multi-arch container image, an SBOM, and SLSA build provenance.

```bash
# Container (recommended)
docker run -d -p 4566:4566 --name kiri-aws ghcr.io/brilhante29/kiri-aws:latest

# Go toolchain
go install github.com/Brilhante29/kiri-aws/cmd/kiri@latest

# Helm
helm install kiri oci://ghcr.io/brilhante29/charts/kiri
```

Or grab a binary from the [latest release](https://github.com/Brilhante29/kiri-aws/releases/latest).
See [RELEASING.md](RELEASING.md) to verify signatures and provenance.

---

## Quickstart

### Option A: Run the published image (recommended)

```bash
docker run -d -p 4566:4566 --name kiri-aws ghcr.io/brilhante29/kiri-aws:latest
```

Verify that the emulator is running:

```bash
curl http://localhost:4566/health
# {"status":"healthy"}
```

### Option B: Docker Compose

A [`docker-compose.yml`](docker-compose.yml) ships with the repository:

```bash
docker compose up -d
```

### Option C: Go module (in-process testing)

Import the module and run the emulator inside your test binary — no container, no
port juggling, and it shuts down with the test:

```go
import (
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/service/s3"
    kiri "github.com/Brilhante29/kiri-aws"
)

func TestUploadsReport(t *testing.T) {
    srv := kiri.NewServer() // random port on localhost
    defer srv.Close()

    client := s3.NewFromConfig(cfg, func(o *s3.Options) {
        o.BaseEndpoint = aws.String(srv.URL)
        o.UsePathStyle = true
    })
    // ... exercise the code under test against client
}
```

Or run it straight from a checkout:

```bash
go run ./cmd/kiri --port 4566
```

### Point your tools at it

```bash
export AWS_ACCESS_KEY_ID="test"
export AWS_SECRET_ACCESS_KEY="test"
export AWS_DEFAULT_REGION="us-east-1"
export AWS_ENDPOINT_URL="http://localhost:4566"

aws s3 mb s3://my-billing-bucket

aws dynamodb create-table \
  --table-name Orders \
  --attribute-definitions AttributeName=ID,AttributeType=S \
  --key-schema AttributeName=ID,KeyType=HASH \
  --billing-mode PAY_PER_REQUEST
```

---

## Language & tooling setup

### Go SDK

```go
import (
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))

client := s3.NewFromConfig(cfg, func(o *s3.Options) {
    o.BaseEndpoint = aws.String("http://localhost:4566")
    o.UsePathStyle = true
})
```

### Python (boto3)

```python
import boto3

s3 = boto3.client(
    "s3",
    endpoint_url="http://localhost:4566",
    region_name="us-east-1",
    aws_access_key_id="test",
    aws_secret_access_key="test",
)
```

### Node.js / TypeScript

```javascript
const { S3Client } = require('@aws-sdk/client-s3');

const s3 = new S3Client({
  endpoint: 'http://localhost:4566',
  region: 'us-east-1',
  forcePathStyle: true,
  credentials: { accessKeyId: 'test', secretAccessKey: 'test' },
});
```

### Terraform

```hcl
provider "aws" {
  region                      = "us-east-1"
  access_key                  = "test"
  secret_key                  = "test"
  skip_credentials_validation = true
  skip_requesting_account_id  = true
  s3_use_path_style           = true

  endpoints {
    s3       = "http://localhost:4566"
    dynamodb = "http://localhost:4566"
    sqs      = "http://localhost:4566"
  }
}
```

### AWS CLI

```bash
aws --endpoint-url http://localhost:4566 s3 ls
```

---

## Cost surface (Cost Explorer analogue)

`kiri-aws` tracks resource usage against a pricing catalog, so you can project
what an architecture would cost before it exists. Query it with the real
Cost Explorer API:

```bash
aws --endpoint-url http://localhost:4566 ce get-cost-and-usage \
  --time-period Start=2026-01-01,End=2026-12-31 \
  --granularity MONTHLY \
  --metrics UnblendedCost
```

### Time Machine

Billing only becomes interesting once time passes. Advance the virtual clock and
every accrual, schedule, and expiry moves with it:

```bash
curl -X POST "http://localhost:4566/_kiri/time/advance?days=30"
```

```json
{
  "status": "success",
  "advanced_by": "720h0m0s",
  "current_virtual_time": "2026-08-21T05:13:21Z"
}
```

Thirty days of uptime accrue in milliseconds, then query Cost Explorer again to
see the bill that architecture would have produced.

---

## Supported services

116 services are registered and reachable on the same endpoint. The list below is
generated from each service's own metadata by `make readme`, and a test fails the
build if it drifts.

<details>
<summary><strong>Show all 116 services</strong></summary>

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

</details>

---

## Configuration

| Variable | Default | Description |
|---|---|---|
| `KIRI_HOST` | `0.0.0.0` | Bind address |
| `KIRI_PORT` | `4566` | HTTP port for every AWS protocol |
| `KIRI_LOG_LEVEL` | `info` | Logging verbosity: `debug`, `info`, `warn`, `error` |
| `KIRI_DATA_DIR` | *(unset)* | Directory for state snapshots (enables persistence) |
| `KIRI_INIT_DIR` | *(unset)* | Directory of init scripts executed on startup |
| `KIRI_PPROF` | *(unset)* | Set to `1`/`true`/`on` to expose the pprof endpoints |

---

## Architecture

```
                 ┌──────────────────────────────────────────────┐
                 │                 kiri Server                  │
                 ├──────────────────────────────────────────────┤
                 │        HTTP Router / Dispatchers (:4566)     │
                 │   AWS JSON 1.x  ·  Query  ·  REST  ·  CBOR   │
                 └──────────────────────┬───────────────────────┘
                                        │
                                        ▼
                 ┌──────────────────────────────────────────────┐
                 │           Unified Service Registry           │
                 │                (116 Services)                │
                 └──────────────────────┬───────────────────────┘
                                        │
                        ┌───────────────┴───────────────┐
                        ▼                               ▼
            ┌───────────────────────┐      ┌────────────────────────┐
            │  Billing engine +     │      │  Storage + persistence │
            │  virtual clock        │      │  ($KIRI_DATA_DIR)      │
            └───────────────────────┘      └────────────────────────┘
```

One Go process fronts every service. Requests arrive on a single port and are
routed by protocol: AWS JSON (`X-Amz-Target`), Query (form-encoded), REST, and
Smithy RPC v2 CBOR. Services register themselves through `init()` hooks, share a
virtual clock, and report usage into the billing engine.

---

## Examples

Runnable samples live in [`test/integration`](test/integration): each file drives
a service with the real AWS SDK for Go v2 against a running emulator, which
doubles as executable documentation for the supported request shapes.

---

## Releases and supply chain

Releases are automated end to end — see [RELEASING.md](RELEASING.md). Every tag
publishes signed binaries (linux, macOS, Windows on amd64/arm64), checksums
signed keylessly with [cosign](https://sigstore.dev), an SBOM per archive, SLSA
build provenance, a multi-arch image on `ghcr.io`, and the Helm chart.

```bash
cosign verify ghcr.io/brilhante29/kiri-aws:latest \
  --certificate-identity-regexp '^https://github.com/Brilhante29/kiri-aws/' \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com
```

---

## Contributing

Issues and pull requests are welcome. Read [CONTRIBUTING.md](CONTRIBUTING.md)
first: it covers the layout, how to add a service, and the local checks. Pull
request titles follow [Conventional Commits](https://www.conventionalcommits.org)
because the release automation derives the next version from them.

---

## License

`kiri-aws` is licensed under the [MIT License](LICENSE).

### Upstream attribution

`kiri-aws` was built upon and evolved from the open-source AWS emulation engine
foundation created by **sivchari/kumo** (and awsim). We express our sincere
appreciation to the original authors and open-source contributors for their
foundational work.

Copyright (c) 2025 sivchari  
Copyright (c) 2026 Guilherme Brilhante and the kiri-aws contributors
