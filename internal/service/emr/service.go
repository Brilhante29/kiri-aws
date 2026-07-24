package emr

import (
	"github.com/kiro-aws/kiro-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "emr"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "EMR",
		Category:    "Analytics & ML",
		Description: "Big data processing platform",
	}
}

func init() {
	service.Register(New())
}
