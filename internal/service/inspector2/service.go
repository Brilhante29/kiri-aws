package inspector2

import (
	"github.com/kiro-aws/kiro-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "inspector2"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Inspector v2",
		Category:    "Security & Identity",
		Description: "Automated vulnerability management",
	}
}

func init() {
	service.Register(New())
}
