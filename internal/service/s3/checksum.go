package s3

import (
	"crypto/sha1" //nolint:gosec // SHA-1 is one of the additional checksum algorithms the S3 API defines.
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/binary"
	"hash/crc32"
	"hash/crc64"
	"io"
	"net/http"
	"strings"
)

// Additional checksums (x-amz-checksum-*). The AWS SDKs attach one to every
// upload and, with response checksum validation on by default, request it
// back on GetObject and HeadObject through x-amz-checksum-mode: ENABLED.
// Objects stored without a client checksum get CRC64NVME, the default S3
// applies to new objects.

const (
	// checksumAlgorithmMetadataKey is a reserved metadata key that carries the
	// upload's checksum algorithm from the handler into storage, where
	// applyChecksumMetadata moves it onto the object.
	checksumAlgorithmMetadataKey = "x-amz-checksum-algorithm"
	checksumModeHeader           = "X-Amz-Checksum-Mode"
	checksumTypeHeader           = "X-Amz-Checksum-Type"
	checksumModeEnabled          = "ENABLED"
	checksumTypeFullObject       = "FULL_OBJECT"
	defaultChecksumAlgorithm     = "crc64nvme"
)

// checksumAlgorithms lists the algorithms S3 accepts, in lookup order.
var checksumAlgorithms = []string{"crc32", "crc32c", "crc64nvme", "sha1", "sha256", "sha512"}

var (
	castagnoliTable = crc32.MakeTable(crc32.Castagnoli)
	// CRC-64/NVME uses polynomial 0xAD93D23594C935A9; hash/crc64 takes it in
	// reflected form and applies the all-ones initial value and final XOR the
	// variant specifies.
	nvmeTable = crc64.MakeTable(0x9a6c9329ac4bc9b5)
)

// computeChecksum returns the base64 checksum of data for an S3 algorithm
// name, or "" for an unknown algorithm. CRC values are big-endian, as S3
// encodes them.
func computeChecksum(algorithm string, data []byte) string {
	var sum []byte

	switch algorithm {
	case "crc32":
		sum = binary.BigEndian.AppendUint32(nil, crc32.ChecksumIEEE(data))
	case "crc32c":
		sum = binary.BigEndian.AppendUint32(nil, crc32.Checksum(data, castagnoliTable))
	case "crc64nvme":
		sum = binary.BigEndian.AppendUint64(nil, crc64.Checksum(data, nvmeTable))
	case "sha1":
		digest := sha1.Sum(data) //nolint:gosec // integrity checksum defined by the S3 API, not a security control
		sum = digest[:]
	case "sha256":
		digest := sha256.Sum256(data)
		sum = digest[:]
	case "sha512":
		digest := sha512.Sum512(data)
		sum = digest[:]
	default:
		return ""
	}

	return base64.StdEncoding.EncodeToString(sum)
}

// requestChecksum returns the algorithm and value of the first
// x-amz-checksum-<algorithm> header on a request, or empty strings.
func requestChecksum(header http.Header) (string, string) {
	for _, algorithm := range checksumAlgorithms {
		if value := header.Get("X-Amz-Checksum-" + algorithm); value != "" {
			return algorithm, value
		}
	}

	return "", ""
}

// checksumHeaderName returns the response header that carries a checksum.
func checksumHeaderName(algorithm string) string {
	return "x-amz-checksum-" + algorithm
}

// uploadChecksum is the additional checksum an upload declared, if any.
type uploadChecksum struct {
	algorithm string
	value     string
}

// writeHeader echoes the declared checksum on the upload response, as S3 does.
func (c uploadChecksum) writeHeader(w http.ResponseWriter) {
	if c.algorithm != "" {
		w.Header().Set(checksumHeaderName(c.algorithm), c.value)
	}
}

// readUploadBody reads an upload body and verifies any declared
// x-amz-checksum-* value against it, recording the algorithm in metadata for
// storage. On failure it writes the S3 error (BadDigest for a mismatch) and
// returns false.
func readUploadBody(w http.ResponseWriter, r *http.Request, metadata map[string]string) ([]byte, uploadChecksum, bool) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		writeS3Error(w, r, "InternalError", "Internal server error", http.StatusInternalServerError)

		return nil, uploadChecksum{}, false
	}

	algorithm, value := requestChecksum(r.Header)
	if algorithm == "" {
		return data, uploadChecksum{}, true
	}

	if computeChecksum(algorithm, data) != value {
		writeS3Error(w, r, "BadDigest",
			"The "+strings.ToUpper(algorithm)+" you specified did not match the calculated checksum.",
			http.StatusBadRequest)

		return nil, uploadChecksum{}, false
	}

	metadata[checksumAlgorithmMetadataKey] = algorithm

	return data, uploadChecksum{algorithm: algorithm, value: value}, true
}

// writeChecksumHeaders adds the object's full-object checksum when the
// request asks for it, as S3 does for GetObject and HeadObject. Part reads
// (partNumber) carry no full-object checksum. Objects stored before checksums
// were recorded fall back to the CRC64NVME of the body; a metadata-only copy
// (HeadObject) of such an object carries no body, so no checksum is reported
// rather than a wrong one.
func writeChecksumHeaders(w http.ResponseWriter, r *http.Request, obj *Object) {
	if !strings.EqualFold(r.Header.Get(checksumModeHeader), checksumModeEnabled) || r.URL.Query().Get("partNumber") != "" {
		return
	}

	algorithm, value := obj.ChecksumAlgorithm, obj.Checksum
	if value == "" {
		if obj.Body == nil && obj.Size > 0 {
			return
		}

		algorithm = defaultChecksumAlgorithm
		value = computeChecksum(algorithm, obj.Body)
	}

	w.Header().Set(checksumHeaderName(algorithm), value)
	w.Header().Set(checksumTypeHeader, checksumTypeFullObject)
}

// applyChecksumMetadata records the object's checksum: the algorithm carried
// by the reserved metadata key, or the CRC64NVME default. It removes the key,
// so it is never echoed as user metadata.
func applyChecksumMetadata(obj *Object, metadata map[string]string) {
	algorithm := metadata[checksumAlgorithmMetadataKey]
	delete(metadata, checksumAlgorithmMetadataKey)

	if algorithm == "" {
		algorithm = defaultChecksumAlgorithm
	}

	obj.ChecksumAlgorithm = algorithm
	obj.Checksum = computeChecksum(algorithm, obj.Body)
}
