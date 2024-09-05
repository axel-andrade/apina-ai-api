package common_ptr

import (
	"github.com/axel-andrade/opina-ai-api/internal/core/domain"
)

type ContentFormatted struct {
	ID          string  `json:"id" example:"123" description:"O ID do contéudo"`
	Status      string  `json:"status" example:"ready" description:"O status do contéudo"`
	Name        string  `json:"name" example:"Áudio 1" description:"O nome do contéudo"`
	Description string  `json:"description" example:"Áudio 1" description:"A descrição do contéudo"`
	Kind        string  `json:"kind" example:"audio" description:"O tipo do contéudo"`
	File        []byte  `json:"file" example:"audio.mp3" description:"O arquivo do contéudo"`
	Filename    *string `json:"file_name" example:"audio.mp3" description:"O nome do arquivo do contéudo"`
	Text        string  `json:"text" example:"Texto do contéudo" description:"O texto do contéudo"`
	CreatedAt   string  `json:"created_at" example:"2021-01-01T00:00:00Z" description:"Data de criação do contéudo"`
	UpdatedAt   string  `json:"updated_at" example:"2021-01-01T00:00:00Z" description:"Data de atualização do contéudo"`
}

type ContentPresenter struct{}

func BuildContentPresenter() *ContentPresenter {
	return &ContentPresenter{}
}

func (ptr *ContentPresenter) Format(content *domain.Content) ContentFormatted {
	return ContentFormatted{
		ID:          content.ID,
		Status:      content.Status,
		Name:        content.Name,
		Description: content.Description,
		Kind:        content.Kind,
		File:        content.File,
		Filename:    &content.Filename,
		Text:        content.Text,
		CreatedAt:   content.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   content.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}

func (ptr *ContentPresenter) FormatList(contents []domain.Content) []ContentFormatted {
	var contentsFormatted []ContentFormatted = make([]ContentFormatted, 0)

	for _, content := range contents {
		contentsFormatted = append(contentsFormatted, ptr.Format(&content))
	}

	return contentsFormatted
}
