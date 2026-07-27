// Package codecommit emulates the AWS codecommit API surface.
package codecommit

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS codecommit service.
type Service struct{}

// New creates a new codecommit service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "codecommit"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "CodeCommit",
		Category:    "Developer Tools",
		Description: "Source control repository service",
	}
}

func init() {
	service.Register(New())
}
