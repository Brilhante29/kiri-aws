package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	fmt.Println("🚀 Starting Kiro S3 Demo...")

	// Create a custom endpoint resolver that points to the local kiri server
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:           "http://host.docker.internal:4566",
			SigningRegion: "us-east-1",
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID: "test", SecretAccessKey: "test",
			}, nil
		})),
	)
	if err != nil {
		log.Fatalf("❌ unable to load SDK config: %v", err)
	}

	// Use path style for local testing (http://localhost:4566/bucket instead of http://bucket.localhost:4566)
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	fmt.Println("⏳ Attempting to list buckets in Kiro (should be empty)...")
	out, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Fatalf("❌ failed to list buckets: %v", err)
	}
	fmt.Printf("🪣 Initial Buckets: %v\n", out.Buckets)

	bucketName := "kiri-demo-bucket"
	fmt.Printf("🔨 Creating a new bucket '%s'...\n", bucketName)
	_, err = client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Fatalf("❌ failed to create bucket: %v", err)
	}
	fmt.Println("✅ Bucket created successfully!")

	fmt.Println("🔍 Listing buckets again to confirm creation...")
	out, err = client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Fatalf("❌ failed to list buckets: %v", err)
	}
	for _, b := range out.Buckets {
		fmt.Printf("👉 Found bucket: %s\n", *b.Name)
	}
	
	fmt.Println("🎉 Demo completed successfully!")
}
