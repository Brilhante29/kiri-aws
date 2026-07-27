// Package storagegateway emulates the AWS storagegateway API surface.
package storagegateway

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS storagegateway service.
type Service struct{}

// New creates a new storagegateway service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "storagegateway"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Storage Gateway",
		Category:    "Storage",
		Description: "Hybrid cloud storage connection",
	}
}

func init() {
	service.Register(New())
}
