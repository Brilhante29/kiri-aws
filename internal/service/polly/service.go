// Package polly emulates the AWS polly API surface.
package polly

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS polly service.
type Service struct{}

// New creates a new polly service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "polly"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Polly",
		Category:    "Analytics & ML",
		Description: "Text-to-speech synthesis",
	}
}

func init() {
	service.Register(New())
}
