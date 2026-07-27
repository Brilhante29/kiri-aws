// Package translate emulates the AWS translate API surface.
package translate

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS translate service.
type Service struct{}

// New creates a new translate service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "translate"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Translate",
		Category:    "Analytics & ML",
		Description: "Neural machine translation",
	}
}

func init() {
	service.Register(New())
}
