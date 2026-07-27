// Package codepipeline emulates the AWS codepipeline API surface.
package codepipeline

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS codepipeline service.
type Service struct{}

// New creates a new codepipeline service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "codepipeline"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "CodePipeline",
		Category:    "Developer Tools",
		Description: "Continuous delivery workflow service",
	}
}

func init() {
	service.Register(New())
}
