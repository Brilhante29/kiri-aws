// Package synthetics emulates the AWS synthetics API surface.
package synthetics

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS synthetics service.
type Service struct{}

// New creates a new synthetics service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "synthetics"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "CloudWatch Synthetics",
		Category:    "Monitoring & Logging",
		Description: "Synthetic canary monitoring",
	}
}

func init() {
	service.Register(New())
}
