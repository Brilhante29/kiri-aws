// Package iotdata emulates the AWS iotdata API surface.
package iotdata

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS iotdata service.
type Service struct{}

// New creates a new iotdata service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "iotdata"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "IoT Data Plane",
		Category:    "Messaging & Integration",
		Description: "Real-time IoT message broker",
	}
}

func init() {
	service.Register(New())
}
