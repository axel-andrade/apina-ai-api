package process_audio

import (
	"github.com/axel-andrade/opina-ai-api/internal/core/domain"
)

type ProcessAudioGateway interface {
	FindContentByID(id string) (*domain.Content, error)
	TransformAudioToText(file []byte, filename string) (string, error)
	CorrectText(text string) (string, error)
	UpdateContent(id string, c *domain.Content) error
}

type ProcessAudioInput struct {
	ContentID string
}
