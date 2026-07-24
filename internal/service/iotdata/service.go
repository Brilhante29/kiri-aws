package iotdata

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "iotdata"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "IoT Data Plane",
		Category:    "Messaging & Integration",
		Description: "Real-time IoT message broker",
	}
}

func init() {
	service.Register(New())
}
