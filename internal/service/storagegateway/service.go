package storagegateway

import (
	"github.com/kiro-aws/kiro-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "storagegateway"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Storage Gateway",
		Category:    "Storage",
		Description: "Hybrid cloud storage connection",
	}
}

func init() {
	service.Register(New())
}
