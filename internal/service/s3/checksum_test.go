package s3

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestComputeChecksumCheckValues pins each algorithm to the published check
// value for "123456789" (CRC catalogue) or its standard digest.
func TestComputeChecksumCheckValues(t *testing.T) {
	t.Parallel()

	tests := map[string]string{
		"crc32":     "y/Q5Jg==",     // 0xCBF43926
		"crc32c":    "4waSgw==",     // 0xE3069283
		"crc64nvme": "rosUhgp5mIg=", // 0xAE8B14860A799888
		"sha1":      "98O8HYCOBHMq32eZZczDTKeuNEE=",
		"sha256":    "FeKw08M4keuw8e9gnsQZQgwg4yDOlMZfvIwzEkSOsiU=",
		"sha512":    "2eZ2LdHI6vbWGzxhkvxAjU1tXxF20MKRabwk5xw/J0rSf81YEbMT1oH35V7ALXPUmclUVba1u1A6z1dPuo/+hQ==",
	}

	for algorithm, want := range tests {
		if got := computeChecksum(algorithm, []byte("123456789")); got != want {
			t.Errorf("computeChecksum(%q): got %q, want %q", algorithm, got, want)
		}
	}

	if got := computeChecksum("md5", []byte("123456789")); got != "" {
		t.Errorf("computeChecksum(md5): got %q, want empty for an unknown algorithm", got)
	}
}

func newChecksumTestService(t *testing.T, bucket string) (*Service, Storage) {
	t.Helper()

	store := NewMemoryStorage()
	if err := store.CreateBucket(context.Background(), bucket); err != nil {
		t.Fatalf("CreateBucket: %v", err)
	}

	return New(store, ""), store
}

func objectRequest(method, bucket, key, body string) *http.Request {
	req := httptest.NewRequest(method, "/"+bucket+"/"+key, strings.NewReader(body))
	req.SetPathValue("bucket", bucket)
	req.SetPathValue("key", key)

	return req
}

// TestPutObjectStoresAndReturnsRequestChecksum covers the SDK default path:
// the upload carries x-amz-checksum-crc32, PutObject echoes it, and
// GetObject and HeadObject return it when x-amz-checksum-mode is ENABLED.
func TestPutObjectStoresAndReturnsRequestChecksum(t *testing.T) {
	t.Parallel()

	const (
		bucket = "checksum-roundtrip"
		key    = "greeting.txt"
		crc32  = "NhCmhg==" // CRC32 of "hello"
	)

	svc, _ := newChecksumTestService(t, bucket)

	put := objectRequest(http.MethodPut, bucket, key, "hello")
	put.Header.Set("X-Amz-Checksum-Crc32", crc32)
	put.Header.Set("X-Amz-Sdk-Checksum-Algorithm", "CRC32")

	putRec := httptest.NewRecorder()
	svc.PutObject(putRec, put)

	if putRec.Code != http.StatusOK {
		t.Fatalf("PutObject status: got %d, want %d (body=%s)", putRec.Code, http.StatusOK, putRec.Body.String())
	}

	if got := putRec.Header().Get("X-Amz-Checksum-Crc32"); got != crc32 {
		t.Fatalf("PutObject checksum header: got %q, want %q", got, crc32)
	}

	assertReturnedChecksum(t, "GetObject", svc.GetObject, objectRequest(http.MethodGet, bucket, key, ""), crc32)
	assertReturnedChecksum(t, "HeadObject", svc.HeadObject, objectRequest(http.MethodHead, bucket, key, ""), crc32)

	plain := httptest.NewRecorder()
	svc.GetObject(plain, objectRequest(http.MethodGet, bucket, key, ""))

	if got := plain.Header().Get("X-Amz-Checksum-Crc32"); got != "" {
		t.Errorf("GetObject without checksum mode returned a checksum: %q", got)
	}
}

// assertReturnedChecksum requests an object with x-amz-checksum-mode: ENABLED
// and checks the CRC32 full-object checksum and that no internal metadata
// leaks as user metadata.
func assertReturnedChecksum(t *testing.T, name string, call http.HandlerFunc, req *http.Request, want string) {
	t.Helper()

	req.Header.Set("X-Amz-Checksum-Mode", "ENABLED")

	rec := httptest.NewRecorder()
	call(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("%s status: got %d, want %d", name, rec.Code, http.StatusOK)
	}

	if got := rec.Header().Get("X-Amz-Checksum-Crc32"); got != want {
		t.Errorf("%s checksum: got %q, want %q", name, got, want)
	}

	if got := rec.Header().Get("X-Amz-Checksum-Type"); got != checksumTypeFullObject {
		t.Errorf("%s checksum type: got %q, want %q", name, got, checksumTypeFullObject)
	}

	for header := range rec.Header() {
		if strings.HasPrefix(strings.ToLower(header), "x-amz-meta-") {
			t.Errorf("%s leaked internal metadata as %s", name, header)
		}
	}
}

// TestPutObjectRejectsChecksumMismatch verifies S3's BadDigest response and
// that nothing is stored when the body does not match the declared checksum.
func TestPutObjectRejectsChecksumMismatch(t *testing.T) {
	t.Parallel()

	const bucket = "checksum-mismatch"

	svc, store := newChecksumTestService(t, bucket)

	req := objectRequest(http.MethodPut, bucket, "tampered.txt", "hello")
	req.Header.Set("X-Amz-Checksum-Sha256", computeChecksum("sha256", []byte("goodbye")))

	rec := httptest.NewRecorder()
	svc.PutObject(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("PutObject status: got %d, want %d", rec.Code, http.StatusBadRequest)
	}

	if !strings.Contains(rec.Body.String(), "BadDigest") {
		t.Fatalf("PutObject error body: got %s, want BadDigest", rec.Body.String())
	}

	if _, err := store.GetObject(context.Background(), bucket, "tampered.txt"); err == nil {
		t.Fatal("object was stored despite the checksum mismatch")
	}
}

// TestGetObjectDefaultsToCRC64NVME verifies that an object uploaded without a
// checksum reports the CRC64NVME default, and that ranged reads, which S3
// does not checksum, carry no checksum header.
func TestGetObjectDefaultsToCRC64NVME(t *testing.T) {
	t.Parallel()

	const (
		bucket = "checksum-default"
		key    = "data.bin"
		body   = "123456789"
	)

	svc, _ := newChecksumTestService(t, bucket)

	putRec := httptest.NewRecorder()
	svc.PutObject(putRec, objectRequest(http.MethodPut, bucket, key, body))

	if putRec.Code != http.StatusOK {
		t.Fatalf("PutObject status: got %d, want %d", putRec.Code, http.StatusOK)
	}

	get := objectRequest(http.MethodGet, bucket, key, "")
	get.Header.Set("X-Amz-Checksum-Mode", "ENABLED")

	rec := httptest.NewRecorder()
	svc.GetObject(rec, get)

	if got := rec.Header().Get("X-Amz-Checksum-Crc64nvme"); got != "rosUhgp5mIg=" {
		t.Fatalf("default checksum: got %q, want the CRC64NVME of %q", got, body)
	}

	ranged := objectRequest(http.MethodGet, bucket, key, "")
	ranged.Header.Set("X-Amz-Checksum-Mode", "ENABLED")
	ranged.Header.Set("Range", "bytes=0-3")

	rangedRec := httptest.NewRecorder()
	svc.GetObject(rangedRec, ranged)

	if rangedRec.Code != http.StatusPartialContent {
		t.Fatalf("ranged GetObject status: got %d, want %d", rangedRec.Code, http.StatusPartialContent)
	}

	if got := rangedRec.Header().Get("X-Amz-Checksum-Crc64nvme"); got != "" {
		t.Fatalf("ranged GetObject returned a full-object checksum: %q", got)
	}
}

// TestCompleteMultipartUploadRecordsChecksum verifies that a completed
// multipart object reports a full-object CRC64NVME, also through HeadObject.
func TestCompleteMultipartUploadRecordsChecksum(t *testing.T) {
	t.Parallel()

	const (
		bucket = "checksum-multipart"
		key    = "big.bin"
	)

	svc, store := newChecksumTestService(t, bucket)
	ctx := context.Background()

	upload, err := store.CreateMultipartUpload(ctx, bucket, key, nil)
	if err != nil {
		t.Fatalf("CreateMultipartUpload: %v", err)
	}

	part1, err := store.UploadPart(ctx, bucket, key, upload.UploadID, 1, strings.NewReader("12345"))
	if err != nil {
		t.Fatalf("UploadPart 1: %v", err)
	}

	part2, err := store.UploadPart(ctx, bucket, key, upload.UploadID, 2, strings.NewReader("6789"))
	if err != nil {
		t.Fatalf("UploadPart 2: %v", err)
	}

	if _, err := store.CompleteMultipartUpload(ctx, bucket, key, upload.UploadID, []PartRequest{
		{PartNumber: 1, ETag: part1.ETag},
		{PartNumber: 2, ETag: part2.ETag},
	}); err != nil {
		t.Fatalf("CompleteMultipartUpload: %v", err)
	}

	head := objectRequest(http.MethodHead, bucket, key, "")
	head.Header.Set("X-Amz-Checksum-Mode", "ENABLED")

	rec := httptest.NewRecorder()
	svc.HeadObject(rec, head)

	if got := rec.Header().Get("X-Amz-Checksum-Crc64nvme"); got != "rosUhgp5mIg=" {
		t.Fatalf("multipart HeadObject checksum: got %q, want the CRC64NVME of %q", got, "123456789")
	}
}
