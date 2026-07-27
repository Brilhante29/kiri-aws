// Package iot emulates the AWS iot API surface.
package iot

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS iot service.
type Service struct{}

// New creates a new iot service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "iot"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "IoT Core",
		Category:    "Messaging & Integration",
		Description: "IoT device connection and messaging",
	}
}

func init() {
	service.Register(New())
}
