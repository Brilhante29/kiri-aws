package lightsail

import (
	"github.com/kiro-aws/kiro-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "lightsail"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Lightsail",
		Category:    "Compute",
		Description: "Virtual private server service",
	}
}

func init() {
	service.Register(New())
}
