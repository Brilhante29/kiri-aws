// Package opensearch emulates the AWS OpenSearch API surface.
package opensearch

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS opensearch service.
type Service struct{}

// New creates a new opensearch service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "opensearch"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "OpenSearch",
		Category:    "Analytics & ML",
		Description: "Search and analytics suite",
	}
}

func init() {
	service.Register(New())
}
