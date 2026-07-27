// Package wafv2 emulates the AWS WAFv2 API surface.
package wafv2

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS wafv2 service.
type Service struct{}

// New creates a new wafv2 service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "wafv2"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "WAFv2",
		Category:    "Security & Identity",
		Description: "Web application firewall v2",
	}
}

func init() {
	service.Register(New())
}
