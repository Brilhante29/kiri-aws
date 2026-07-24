package translate

import (
	"github.com/kiro-aws/kiro-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "translate"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Translate",
		Category:    "Analytics & ML",
		Description: "Neural machine translation",
	}
}

func init() {
	service.Register(New())
}
