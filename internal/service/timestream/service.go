// Package timestream emulates the AWS timestream API surface.
package timestream

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS timestream service.
type Service struct{}

// New creates a new timestream service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "timestream"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Timestream",
		Category:    "Database",
		Description: "Time-series database service",
	}
}

func init() {
	service.Register(New())
}
