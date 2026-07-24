package apprunner

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "apprunner"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "App Runner",
		Category:    "Compute",
		Description: "Containerized web application service",
	}
}

func init() {
	service.Register(New())
}
