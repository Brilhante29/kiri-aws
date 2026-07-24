package wafv2

import (
	"github.com/kiro-aws/kiro-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "wafv2"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "WAFv2",
		Category:    "Security & Identity",
		Description: "Web application firewall v2",
	}
}

func init() {
	service.Register(New())
}
