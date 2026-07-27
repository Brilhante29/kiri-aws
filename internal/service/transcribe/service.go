// Package transcribe emulates the AWS transcribe API surface.
package transcribe

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS transcribe service.
type Service struct{}

// New creates a new transcribe service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "transcribe"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Transcribe",
		Category:    "Analytics & ML",
		Description: "Speech-to-text recognition",
	}
}

func init() {
	service.Register(New())
}
