package keyspaces

import (
	"github.com/kiro-aws/kiro-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "keyspaces"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Keyspaces",
		Category:    "Database",
		Description: "Apache Cassandra-compatible database",
	}
}

func init() {
	service.Register(New())
}
