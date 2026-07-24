package codebuild

import (
	"github.com/kiro-aws/kiro-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "codebuild"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "CodeBuild",
		Category:    "Developer Tools",
		Description: "Continuous integration build service",
	}
}

func init() {
	service.Register(New())
}
