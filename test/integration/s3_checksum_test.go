//go:build integration

package integration

import (
	"io"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// objectChecksums gathers the per-algorithm checksum fields that the SDK
// output types share by name.
type objectChecksums struct {
	crc32, crc32c, crc64nvme, sha1, sha256, sha512 *string
}

// sdkChecksumAlgorithms are the algorithms the SDK can compute for an upload.
// The S3 enum also lists MD5 and XXHASH variants, which the SDK cannot send.
var sdkChecksumAlgorithms = []types.ChecksumAlgorithm{
	types.ChecksumAlgorithmCrc32,
	types.ChecksumAlgorithmCrc32c,
	types.ChecksumAlgorithmCrc64nvme,
	types.ChecksumAlgorithmSha1,
	types.ChecksumAlgorithmSha256,
	types.ChecksumAlgorithmSha512,
}

func (c objectChecksums) get(algorithm types.ChecksumAlgorithm) string {
	return aws.ToString(map[types.ChecksumAlgorithm]*string{
		types.ChecksumAlgorithmCrc32:     c.crc32,
		types.ChecksumAlgorithmCrc32c:    c.crc32c,
		types.ChecksumAlgorithmCrc64nvme: c.crc64nvme,
		types.ChecksumAlgorithmSha1:      c.sha1,
		types.ChecksumAlgorithmSha256:    c.sha256,
		types.ChecksumAlgorithmSha512:    c.sha512,
	}[algorithm])
}

// TestS3_ChecksumRoundTrip drives the SDK's integrity path for every
// additional-checksum algorithm: PutObject sends the checksum the SDK
// computes, kiri verifies and stores it, and GetObject and HeadObject with
// checksum mode enabled return it, so the SDK validates the downloaded body.
func TestS3_ChecksumRoundTrip(t *testing.T) {
	client := newS3Client(t)
	ctx := t.Context()
	bucket := "test-checksum-roundtrip"
	body := "region,total\nsa-east-1,42\n"

	if _, err := client.CreateBucket(ctx, &s3.CreateBucketInput{Bucket: aws.String(bucket)}); err != nil {
		t.Fatalf("failed to create bucket: %v", err)
	}

	for _, algorithm := range sdkChecksumAlgorithms {
		key := "report-" + strings.ToLower(string(algorithm)) + ".csv"

		put, err := client.PutObject(ctx, &s3.PutObjectInput{
			Bucket:            aws.String(bucket),
			Key:               aws.String(key),
			Body:              strings.NewReader(body),
			ChecksumAlgorithm: algorithm,
		})
		if err != nil {
			t.Fatalf("%s: failed to put object: %v", algorithm, err)
		}

		want := objectChecksums{put.ChecksumCRC32, put.ChecksumCRC32C, put.ChecksumCRC64NVME, put.ChecksumSHA1, put.ChecksumSHA256, put.ChecksumSHA512}.get(algorithm)
		if want == "" {
			t.Fatalf("%s: PutObject returned no checksum", algorithm)
		}

		got, err := client.GetObject(ctx, &s3.GetObjectInput{
			Bucket:       aws.String(bucket),
			Key:          aws.String(key),
			ChecksumMode: types.ChecksumModeEnabled,
		})
		if err != nil {
			t.Fatalf("%s: failed to get object: %v", algorithm, err)
		}

		data, err := io.ReadAll(got.Body)
		_ = got.Body.Close()

		if err != nil || string(data) != body {
			t.Fatalf("%s: body read failed or differs: %v %q", algorithm, err, data)
		}

		if returned := (objectChecksums{got.ChecksumCRC32, got.ChecksumCRC32C, got.ChecksumCRC64NVME, got.ChecksumSHA1, got.ChecksumSHA256, got.ChecksumSHA512}).get(algorithm); returned != want {
			t.Errorf("%s: GetObject checksum = %q, want %q", algorithm, returned, want)
		}

		if got.ChecksumType != types.ChecksumTypeFullObject {
			t.Errorf("%s: GetObject checksum type = %q, want %q", algorithm, got.ChecksumType, types.ChecksumTypeFullObject)
		}

		head, err := client.HeadObject(ctx, &s3.HeadObjectInput{
			Bucket:       aws.String(bucket),
			Key:          aws.String(key),
			ChecksumMode: types.ChecksumModeEnabled,
		})
		if err != nil {
			t.Fatalf("%s: failed to head object: %v", algorithm, err)
		}

		if returned := (objectChecksums{head.ChecksumCRC32, head.ChecksumCRC32C, head.ChecksumCRC64NVME, head.ChecksumSHA1, head.ChecksumSHA256, head.ChecksumSHA512}).get(algorithm); returned != want {
			t.Errorf("%s: HeadObject checksum = %q, want %q", algorithm, returned, want)
		}

		if _, err := client.DeleteObject(ctx, &s3.DeleteObjectInput{Bucket: aws.String(bucket), Key: aws.String(key)}); err != nil {
			t.Fatalf("%s: failed to delete object: %v", algorithm, err)
		}
	}

	if _, err := client.DeleteBucket(ctx, &s3.DeleteBucketInput{Bucket: aws.String(bucket)}); err != nil {
		t.Fatalf("failed to delete bucket: %v", err)
	}
}
