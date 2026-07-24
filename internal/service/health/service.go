package health

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "health"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "AWS Health",
		Category:    "Management & Configuration",
		Description: "Service status and health events",
	}
}

func init() {
	service.Register(New())
}
