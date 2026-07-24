package serverlessrepo

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "serverlessrepo"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Serverless Application Repository",
		Category:    "Compute",
		Description: "Serverless application catalog",
	}
}

func init() {
	service.Register(New())
}
