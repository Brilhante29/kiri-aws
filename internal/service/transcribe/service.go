package transcribe

import (
	"github.com/Brilhante29/kiri-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "transcribe"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Transcribe",
		Category:    "Analytics & ML",
		Description: "Speech-to-text recognition",
	}
}

func init() {
	service.Register(New())
}
