package fsx

import (
	"github.com/kiro-aws/kiro-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "fsx"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "FSx",
		Category:    "Storage",
		Description: "High-performance file systems",
	}
}

func init() {
	service.Register(New())
}
