// Package kinesisvideo emulates the AWS kinesisvideo API surface.
package kinesisvideo

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS kinesisvideo service.
type Service struct{}

// New creates a new kinesisvideo service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "kinesisvideo"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Kinesis Video Streams",
		Category:    "Analytics & ML",
		Description: "Video streaming and analytics",
	}
}

func init() {
	service.Register(New())
}
