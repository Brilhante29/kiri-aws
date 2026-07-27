// Package autoscaling emulates the AWS Auto Scaling API surface.
package autoscaling

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS Auto Scaling service.
type Service struct{}

// New creates a new Auto Scaling service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "autoscaling"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Auto Scaling",
		Category:    "Compute",
		Description: "EC2 auto scaling management",
	}
}

func init() {
	service.Register(New())
}
