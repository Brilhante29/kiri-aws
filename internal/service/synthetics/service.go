package synthetics

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "synthetics"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "CloudWatch Synthetics",
		Category:    "Monitoring & Logging",
		Description: "Synthetic canary monitoring",
	}
}

func init() {
	service.Register(New())
}
