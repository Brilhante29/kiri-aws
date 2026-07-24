package polly

import (
	"github.com/kiro-aws/kiro-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "polly"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Polly",
		Category:    "Analytics & ML",
		Description: "Text-to-speech synthesis",
	}
}

func init() {
	service.Register(New())
}
