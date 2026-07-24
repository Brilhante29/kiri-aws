package qldb

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "qldb"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "QLDB",
		Category:    "Database",
		Description: "Quantum Ledger Database",
	}
}

func init() {
	service.Register(New())
}
