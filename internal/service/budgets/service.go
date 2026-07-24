package budgets

import (
	"github.com/kiro-aws/kiro-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "budgets"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Budgets",
		Category:    "Management & Configuration",
		Description: "AWS budget tracking and alerts",
	}
}

func init() {
	service.Register(New())
}
