// Package fsx emulates the AWS fsx API surface.
package fsx

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS fsx service.
type Service struct{}

// New creates a new fsx service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "fsx"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "FSx",
		Category:    "Storage",
		Description: "High-performance file systems",
	}
}

func init() {
	service.Register(New())
}
