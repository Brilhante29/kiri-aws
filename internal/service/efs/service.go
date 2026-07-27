// Package efs emulates the AWS efs API surface.
package efs

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS efs service.
type Service struct{}

// New creates a new efs service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "efs"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "EFS",
		Category:    "Storage",
		Description: "Elastic file system",
	}
}

func init() {
	service.Register(New())
}
