// Package qldb emulates the AWS QLDB (Quantum Ledger Database) API surface.
package qldb

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS QLDB service.
type Service struct{}

// New creates a new QLDB service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "qldb"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "QLDB",
		Category:    "Database",
		Description: "Quantum Ledger Database",
	}
}

func init() {
	service.Register(New())
}
