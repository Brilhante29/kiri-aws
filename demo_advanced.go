package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamodbtypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

// In AWS SDK Go v2, Cost Explorer isn't in the default set of mock services, but we can call it using HTTP directly 
// or by importing github.com/aws/aws-sdk-go-v2/service/costexplorer. We will just use HTTP for the Kiro CE endpoint 
// since we mapped it via X-Amz-Target anyway!

func main() {
	ctx := context.TODO()

	// Load AWS SDK config pointing to Kiro
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("test", "test", "")),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL:           "http://127.0.0.1:4566",
					SigningRegion: "us-east-1",
				}, nil
			},
		)),
	)
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	fmt.Println("🚀 Iniciando Simulação de E-Commerce (Arquitetura Orientada a Eventos)...")

	// 1. Criar recursos (S3, DynamoDB, SQS, KMS)
	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})
	dyClient := dynamodb.NewFromConfig(cfg)
	sqsClient := sqs.NewFromConfig(cfg)
	kmsClient := kms.NewFromConfig(cfg)

	fmt.Println("📦 Provisionando infraestrutura...")
	
	// Create S3
	bucketName := "ecommerce-invoices"
	_, _ = s3Client.CreateBucket(ctx, &s3.CreateBucketInput{Bucket: &bucketName})

	// Create DynamoDB
	tableName := "Orders"
	_, _ = dyClient.CreateTable(ctx, &dynamodb.CreateTableInput{
		TableName: &tableName,
		KeySchema: []dynamodbtypes.KeySchemaElement{
			{AttributeName: aws.String("OrderID"), KeyType: dynamodbtypes.KeyTypeHash},
		},
		AttributeDefinitions: []dynamodbtypes.AttributeDefinition{
			{AttributeName: aws.String("OrderID"), AttributeType: dynamodbtypes.ScalarAttributeTypeS},
		},
		BillingMode: dynamodbtypes.BillingModePayPerRequest,
	})

	// Create SQS
	queueName := "order-processing-queue"
	qResp, _ := sqsClient.CreateQueue(ctx, &sqs.CreateQueueInput{QueueName: &queueName})
	
	// Create KMS Key
	_, _ = kmsClient.CreateKey(ctx, &kms.CreateKeyInput{})

	fmt.Println("✅ Infraestrutura pronta (S3, DynamoDB, SQS, KMS)")

	fmt.Println("🛒 Simulando pedidos (Uso de recursos)...")
	// Simulate Usage
	if qResp != nil && qResp.QueueUrl != nil {
		_, _ = sqsClient.SendMessage(ctx, &sqs.SendMessageInput{
			QueueUrl:    qResp.QueueUrl,
			MessageBody: aws.String(`{"order_id": "123", "amount": 99.99}`),
		})
	}
	
	_, _ = s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    aws.String("invoice-123.pdf"),
		Body:   strings.NewReader("mock pdf content"),
	})

	fmt.Println("⏳ Acionando a 'Máquina do Tempo' do Kiro (Avançando 30 dias)...")
	
	// Advance time using our new Time Travel API
	req, _ := http.NewRequest("POST", "http://127.0.0.1:4566/_kiri/time/advance?days=30", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("failed to advance time: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("⏱️ Tempo avançado com sucesso: %s\n", string(body))

	fmt.Println("💰 Consultando o AWS Cost Explorer...")
	
	// Query CE via JSON API mapping
	ceReq, _ := http.NewRequest("POST", "http://127.0.0.1:4566/", strings.NewReader(`{"TimePeriod": {"Start": "2026-01-01", "End": "2026-12-31"}, "Granularity": "MONTHLY", "Metrics": ["UnblendedCost"]}`))
	ceReq.Header.Set("Content-Type", "application/x-amz-json-1.1")
	ceReq.Header.Set("X-Amz-Target", "AWSInsightsIndexService.GetCostAndUsage")
	ceResp, err := http.DefaultClient.Do(ceReq)
	if err != nil {
		log.Fatalf("failed to call CE: %v", err)
	}
	defer ceResp.Body.Close()
	ceBody, _ := io.ReadAll(ceResp.Body)
	
	fmt.Println("\n🧾 Fatura Simulada (30 dias depois):")
	fmt.Println(string(ceBody))
}
