// Package textract emulates the AWS textract API surface.
package textract

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS textract service.
type Service struct{}

// New creates a new textract service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "textract"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Textract",
		Category:    "Analytics & ML",
		Description: "Document text and data extraction",
	}
}

func init() {
	service.Register(New())
}
