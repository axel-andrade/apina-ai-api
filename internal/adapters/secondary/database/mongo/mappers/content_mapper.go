package mappers

import (
	"github.com/axel-andrade/opina-ai-api/internal/adapters/secondary/database/mongo/models"
	"github.com/axel-andrade/opina-ai-api/internal/core/domain"
)

type ContentMapper struct {
	BaseMapper
}

func BuildContentMapper(baseMapper *BaseMapper) *ContentMapper {
	return &ContentMapper{BaseMapper: *baseMapper}
}

func (m *ContentMapper) ToDomain(model models.Content) *domain.Content {
	return &domain.Content{
		Base:        *m.BaseMapper.toDomain(model.Base),
		Name:        model.Name,
		Description: model.Description,
		Language:    model.Language,
		Kind:        model.Kind,
		File:        model.File,
		Filename:    model.Filename,
		Status:      model.Status,
		Text:        model.Text,
	}
}

func (m *ContentMapper) ToPersistence(entity domain.Content) models.Content {
	return models.Content{
		Base:        *m.BaseMapper.toPersistence(entity.Base),
		Name:        entity.Name,
		Description: entity.Description,
		Language:    entity.Language,
		Kind:        entity.Kind,
		File:        entity.File,
		Filename:    entity.Filename,
		Status:      entity.Status,
		Text:        entity.Text,
	}
}

func (m *ContentMapper) ToUpdate(entity domain.Content) map[string]interface{} {
	model := make(map[string]interface{})

	if entity.Name != "" {
		model["name"] = entity.Name
	}

	if entity.Description != "" {
		model["description"] = entity.Description
	}

	if entity.Language != "" {
		model["language"] = entity.Language
	}

	if entity.Kind != "" {
		model["kind"] = entity.Kind
	}

	if entity.File != nil {
		model["file"] = entity.File
	}

	if entity.Filename != "" {
		model["filename"] = entity.Filename
	}

	if entity.Status != "" {
		model["status"] = entity.Status
	}

	if entity.Text != "" {
		model["text"] = entity.Text
	}

	return model
}
