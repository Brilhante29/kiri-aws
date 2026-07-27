// Package codebuild emulates the AWS codebuild API surface.
package codebuild

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS codebuild service.
type Service struct{}

// New creates a new codebuild service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "codebuild"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "CodeBuild",
		Category:    "Developer Tools",
		Description: "Continuous integration build service",
	}
}

func init() {
	service.Register(New())
}
