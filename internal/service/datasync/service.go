// Package datasync emulates the AWS datasync API surface.
package datasync

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS datasync service.
type Service struct{}

// New creates a new datasync service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "datasync"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "DataSync",
		Category:    "Storage",
		Description: "Automated data transfer service",
	}
}

func init() {
	service.Register(New())
}
