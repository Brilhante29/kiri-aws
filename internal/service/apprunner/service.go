// Package apprunner emulates the AWS apprunner API surface.
package apprunner

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS apprunner service.
type Service struct{}

// New creates a new apprunner service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "apprunner"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "App Runner",
		Category:    "Compute",
		Description: "Containerized web application service",
	}
}

func init() {
	service.Register(New())
}
