package kinesisvideo

import (
	"github.com/kiro-aws/kiro-aws/internal/service"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Name() string {
	return "kinesisvideo"
}

func (s *Service) RegisterRoutes(_ service.Router) {}

func (s *Service) Meta() service.Meta {
	return service.Meta{
		Display:     "Kinesis Video Streams",
		Category:    "Analytics & ML",
		Description: "Video streaming and analytics",
	}
}

func init() {
	service.Register(New())
}
