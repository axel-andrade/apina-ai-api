package domain

import (
	"fmt"

	err_msg "github.com/axel-andrade/opina-ai-api/internal/core/domain/constants/errors"
)

const (
	ContentStatusProcessing = "processing"
	ContentStatusReady      = "ready"
	ContentKindAudio        = "audio"
	ContentKindText         = "text"
)

type Content struct {
	Base
	Name        string `json:"name"`
	Description string `json:"description"`
	Kind        string `json:"kind"`
	File        []byte `json:"file"`
	Filename    string `json:"filename"`
	Status      string `json:"status"`
	Text        string `json:"text"`
	Language    string `json:"language"`
}

// ContentBuilder é responsável por construir uma instância de Content
type ContentBuilder struct {
	content *Content
}

// NewContentBuilder cria uma nova instância de ContentBuilder
func NewContentBuilder() *ContentBuilder {
	return &ContentBuilder{
		content: &Content{
			Status: ContentStatusProcessing,
		},
	}
}

func (cb *ContentBuilder) WithName(name string) *ContentBuilder {
	cb.content.Name = name
	return cb
}

func (cb *ContentBuilder) WithDescription(description string) *ContentBuilder {
	cb.content.Description = description
	return cb
}

func (cb *ContentBuilder) WithKind(kind string) *ContentBuilder {
	cb.content.Kind = kind
	return cb
}

func (cb *ContentBuilder) WithLanguage(language string) *ContentBuilder {
	cb.content.Language = language
	return cb
}

func (cb *ContentBuilder) WithText(text string) *ContentBuilder {
	cb.content.Text = text
	return cb
}

func (cb *ContentBuilder) WithFilename(filename string) *ContentBuilder {
	cb.content.Filename = filename
	return cb
}

func (cb *ContentBuilder) WithFile(file []byte) *ContentBuilder {
	cb.content.File = file
	return cb
}

func (cb *ContentBuilder) WithStatus(status string) *ContentBuilder {
	cb.content.Status = status
	return cb
}

// Build finaliza a construção do Content e realiza a validação
func (cb *ContentBuilder) Build() (*Content, error) {
	if err := cb.validate(); err != nil {
		return nil, err
	}
	return cb.content, nil
}

func (cb *ContentBuilder) validate() error {
	if cb.content.Name == "" {
		return fmt.Errorf(err_msg.NAME_IS_EMPTY)
	}

	if cb.content.Description == "" {
		return fmt.Errorf(err_msg.DESCRIPTION_IS_EMPTY)
	}

	if cb.content.Kind == "" {
		return fmt.Errorf(err_msg.KIND_IS_EMPTY)
	}

	if cb.content.Language == "" {
		return fmt.Errorf(err_msg.LANGUAGE_IS_EMPTY)
	}

	if cb.content.Kind == ContentKindText && cb.content.Text == "" {
		return fmt.Errorf(err_msg.TEXT_IS_EMPTY)
	}

	if cb.content.Kind == ContentKindAudio && len(cb.content.File) == 0 {
		return fmt.Errorf(err_msg.FILE_IS_EMPTY)
	}

	if cb.content.File != nil && cb.content.Filename == "" {
		return fmt.Errorf(err_msg.FILENAME_IS_EMPTY)
	}

	if cb.content.Kind != ContentKindText && cb.content.Kind != ContentKindAudio {
		return fmt.Errorf(err_msg.INVALID_KIND)
	}

	return nil
}
