package quicksight

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "quicksight"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "QuickSight",
		Category:    "Analytics & ML",
		Description: "Business intelligence service",
	}
}

func init() {
	service.Register(New())
}
