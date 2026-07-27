// Package codeartifact emulates the AWS codeartifact API surface.
package codeartifact

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS codeartifact service.
type Service struct{}

// New creates a new codeartifact service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "codeartifact"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "CodeArtifact",
		Category:    "Developer Tools",
		Description: "Artifact and package repository",
	}
}

func init() {
	service.Register(New())
}
