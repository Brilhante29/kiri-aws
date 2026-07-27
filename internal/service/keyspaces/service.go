// Package keyspaces emulates the AWS keyspaces API surface.
package keyspaces

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS keyspaces service.
type Service struct{}

// New creates a new keyspaces service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "keyspaces"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Keyspaces",
		Category:    "Database",
		Description: "Apache Cassandra-compatible database",
	}
}

func init() {
	service.Register(New())
}
