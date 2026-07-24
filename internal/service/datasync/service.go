package datasync

import (
	"github.com/kiro-aws/kiro-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "datasync"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "DataSync",
		Category:    "Storage",
		Description: "Automated data transfer service",
	}
}

func init() {
	service.Register(New())
}
