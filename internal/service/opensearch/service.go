package opensearch

import (
	"github.com/kiro-aws/kiro-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "opensearch"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "OpenSearch",
		Category:    "Analytics & ML",
		Description: "Search and analytics suite",
	}
}

func init() {
	service.Register(New())
}
