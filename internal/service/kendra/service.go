// Package kendra emulates the AWS kendra API surface.
package kendra

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS kendra service.
type Service struct{}

// New creates a new kendra service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "kendra"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Kendra",
		Category:    "Analytics & ML",
		Description: "Intelligent enterprise search",
	}
}

func init() {
	service.Register(New())
}
