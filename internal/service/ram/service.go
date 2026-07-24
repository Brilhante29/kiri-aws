package ram

import (
	"github.com/kiro-aws/kiro-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "ram"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Resource Access Manager",
		Category:    "Security & Identity",
		Description: "Cross-account resource sharing",
	}
}

func init() {
	service.Register(New())
}
