// Package serverlessrepo emulates the AWS serverlessrepo API surface.
package serverlessrepo

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS serverlessrepo service.
type Service struct{}

// New creates a new serverlessrepo service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "serverlessrepo"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Serverless Application Repository",
		Category:    "Compute",
		Description: "Serverless application catalog",
	}
}

func init() {
	service.Register(New())
}
