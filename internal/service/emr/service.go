// Package emr emulates the AWS emr API surface.
package emr

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

// Service implements the AWS emr service.
type Service struct{}

// New creates a new emr service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "emr"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "EMR",
		Category:    "Analytics & ML",
		Description: "Big data processing platform",
	}
}

func init() {
	service.Register(New())
}
