package codepipeline

import (
	"github.com/kiro-aws/kiro-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "codepipeline"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "CodePipeline",
		Category:    "Developer Tools",
		Description: "Continuous delivery workflow service",
	}
}

func init() {
	service.Register(New())
}
