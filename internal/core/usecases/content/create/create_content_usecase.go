package create_content

import (
	"log"

	"github.com/axel-andrade/opina-ai-api/internal/core/domain"
)

type CreateContentUC struct {
	Gateway CreateContentGateway
}

func BuildCreateContentUC(g CreateContentGateway) *CreateContentUC {
	return &CreateContentUC{g}
}

func (bs *CreateContentUC) Execute(input CreateContentInput) (*CreateContentOutput, error) {
	log.Println("Building content")
	builder := domain.NewContentBuilder().
		WithName(input.Name).
		WithDescription(input.Description).
		WithKind(input.Kind).
		WithLanguage(input.Language)

	if input.Kind == domain.ContentKindText {
		builder.
			WithText(domain.ContentStatusReady).
			WithStatus(domain.ContentStatusReady)
	} else {
		var fileBytes []byte
		var err error

		log.Println("Converting file to bytes")
		fileBytes, err = bs.Gateway.ConvertFileToBytes(input.File)
		if err != nil {
			return nil, err
		}

		builder.
			WithFile(fileBytes).
			WithFilename(input.Filename)
	}

	log.Println("Creating content")
	content, err := builder.Build()

	if err != nil {
		return nil, err
	}

	result, err := bs.Gateway.CreateContent(content)

	if err != nil {
		return nil, err
	}

	return &CreateContentOutput{Content: result}, nil
}
