// Package lightsail emulates the AWS lightsail API surface.
package lightsail

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS lightsail service.
type Service struct{}

// New creates a new lightsail service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "lightsail"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Lightsail",
		Category:    "Compute",
		Description: "Virtual private server service",
	}
}

func init() {
	service.Register(New())
}
