// Package quicksight emulates the AWS quicksight API surface.
package quicksight

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS quicksight service.
type Service struct{}

// New creates a new quicksight service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "quicksight"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "QuickSight",
		Category:    "Analytics & ML",
		Description: "Business intelligence service",
	}
}

func init() {
	service.Register(New())
}
