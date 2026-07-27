// Package inspector2 emulates the AWS inspector2 API surface.
package inspector2

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS inspector2 service.
type Service struct{}

// New creates a new inspector2 service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "inspector2"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Inspector v2",
		Category:    "Security & Identity",
		Description: "Automated vulnerability management",
	}
}

func init() {
	service.Register(New())
}
