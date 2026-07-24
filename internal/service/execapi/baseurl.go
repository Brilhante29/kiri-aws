package execapi

import (
	"fmt"
	"os"
)

// DefaultBaseURL is the kiro server URL used for in-process self-calls
// (execute-api -> Lambda invoke) when KIRO_HOST/KIRO_PORT are not set.
const DefaultBaseURL = "http://localhost:4566"

// ResolveBaseURL builds the kiro server base URL from KIRO_HOST / KIRO_PORT,
// mirroring the convention used by other services that self-call kiro
// endpoints (e.g. eventbridge -> lambda).
func ResolveBaseURL() string {
	if host := os.Getenv("KIRO_HOST"); host != "" {
		port := os.Getenv("KIRO_PORT")
		if port == "" {
			port = "4566"
		}

		return fmt.Sprintf("http://%s:%s", host, port)
	}

	if port := os.Getenv("KIRO_PORT"); port != "" {
		return fmt.Sprintf("http://localhost:%s", port)
	}

	return DefaultBaseURL
}
