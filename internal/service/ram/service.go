// Package ram emulates the AWS ram API surface.
package ram

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS ram service.
type Service struct{}

// New creates a new ram service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "ram"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Resource Access Manager",
		Category:    "Security & Identity",
		Description: "Cross-account resource sharing",
	}
}

func init() {
	service.Register(New())
}
