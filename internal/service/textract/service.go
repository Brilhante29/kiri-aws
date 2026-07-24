package textract

import (
	"github.com/kiro-aws/kiro-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "textract"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Textract",
		Category:    "Analytics & ML",
		Description: "Document text and data extraction",
	}
}

func init() {
	service.Register(New())
}
