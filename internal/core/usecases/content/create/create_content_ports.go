package create_content

import (
	"mime/multipart"

	"github.com/axel-andrade/opina-ai-api/internal/core/domain"
)

type CreateContentGateway interface {
	CreateContent(c *domain.Content) (*domain.Content, error)
	ConvertFileToBytes(file multipart.File) ([]byte, error)
}

type CreateContentInput struct {
	Name        string
	Description string
	Kind        string
	File        multipart.File
	Filename    string
	Language    string
	Text        string
}

type CreateContentOutput struct {
	Content *domain.Content
}
