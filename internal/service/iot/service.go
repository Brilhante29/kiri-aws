package iot

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "iot"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "IoT Core",
		Category:    "Messaging & Integration",
		Description: "IoT device connection and messaging",
	}
}

func init() {
	service.Register(New())
}
