package codecommit

import (
	"github.com/kiro-aws/kiro-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "codecommit"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "CodeCommit",
		Category:    "Developer Tools",
		Description: "Source control repository service",
	}
}

func init() {
	service.Register(New())
}
