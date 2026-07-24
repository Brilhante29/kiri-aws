package bedrock

import (
	"github.com/kiro-aws/kiro-aws/internal/service"
)

// Service implements the AWS Bedrock service.
type Service struct{}

// New creates a new Bedrock service.
func New() *Service {
	return &Service{}
}

// Name returns the service name.
func (s *Service) Name() string {
	return "bedrock"
}

// RegisterRoutes registers routes with the router.
func (s *Service) RegisterRoutes(_ service.Router) {}

// Meta returns the service's documentation metadata.
func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Bedrock",
		Category:    "Analytics & ML",
		Description: "Generative AI foundation models and runtime",
	}
}

func init() {
	service.Register(New())
}
