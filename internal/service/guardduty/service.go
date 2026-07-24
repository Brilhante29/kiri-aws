package guardduty

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "guardduty"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "GuardDuty",
		Category:    "Security & Identity",
		Description: "Threat detection service",
	}
}

func init() {
	service.Register(New())
}
