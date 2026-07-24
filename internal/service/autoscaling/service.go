package autoscaling

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "autoscaling"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Auto Scaling",
		Category:    "Compute",
		Description: "EC2 auto scaling management",
	}
}

func init() {
	service.Register(New())
}
