// Package budgets emulates the AWS budgets API surface.
package budgets

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS budgets service.
type Service struct{}

// New creates a new budgets service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "budgets"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Budgets",
		Category:    "Management & Configuration",
		Description: "AWS budget tracking and alerts",
	}
}

func init() {
	service.Register(New())
}
