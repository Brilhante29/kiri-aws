package kendra

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "kendra"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Kendra",
		Category:    "Analytics & ML",
		Description: "Intelligent enterprise search",
	}
}

func init() {
	service.Register(New())
}
