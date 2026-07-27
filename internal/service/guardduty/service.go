// Package guardduty emulates the AWS guardduty API surface.
package guardduty

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS guardduty service.
type Service struct{}

// New creates a new guardduty service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "guardduty"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "GuardDuty",
		Category:    "Security & Identity",
		Description: "Threat detection service",
	}
}

func init() {
	service.Register(New())
}
