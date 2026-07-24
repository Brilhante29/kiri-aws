package execapi

import (
	"fmt"
	"os"
)

// DefaultBaseURL is the kiri server URL used for in-process self-calls
// (execute-api -> Lambda invoke) when KIRI_HOST/KIRI_PORT are not set.
const DefaultBaseURL = "http://localhost:4566"

// ResolveBaseURL builds the kiri server base URL from KIRI_HOST / KIRI_PORT,
// mirroring the convention used by other services that self-call kiri
// endpoints (e.g. eventbridge -> lambda).
func ResolveBaseURL() string {
	if host := os.Getenv("KIRI_HOST"); host != "" {
		port := os.Getenv("KIRI_PORT")
		if port == "" {
			port = "4566"
		}

		return fmt.Sprintf("http://%s:%s", host, port)
	}

	if port := os.Getenv("KIRI_PORT"); port != "" {
		return fmt.Sprintf("http://localhost:%s", port)
	}

	return DefaultBaseURL
}
