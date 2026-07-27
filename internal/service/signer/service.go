// Package signer emulates the AWS signer API surface.
package signer

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS signer service.
type Service struct{}

// New creates a new signer service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "signer"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Signer",
		Category:    "Security & Identity",
		Description: "Code signing service",
	}
}

func init() {
	service.Register(New())
}
