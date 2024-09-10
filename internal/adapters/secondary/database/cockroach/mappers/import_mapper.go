package cockroach_mappers

import (
	cockroach_models "github.com/axel-andrade/opina-ai-api/internal/adapters/secondary/database/cockroach/models"
	"github.com/axel-andrade/opina-ai-api/internal/core/domain"
)

type ImportMapper struct {
	BaseMapper
}

func (i *ImportMapper) ToDomain(model cockroach_models.ImportModel) *domain.Import {
	return &domain.Import{
		Base:         *i.BaseMapper.toDomain(model.BaseModel),
		UserID:       model.UserID,
		Filename:     model.Filename,
		Status:       model.Status,
		TotalRecords: model.TotalRecords,
		ErrorMessage: model.ErrorMessage,
	}
}

func (i *ImportMapper) ToPersistence(e domain.Import) *cockroach_models.ImportModel {
	return &cockroach_models.ImportModel{
		BaseModel:    *i.BaseMapper.toPersistence(e.Base),
		UserID:       e.UserID,
		Filename:     e.Filename,
		Status:       e.Status,
		TotalRecords: e.TotalRecords,
		ErrorMessage: e.ErrorMessage,
	}
}

func (i *ImportMapper) ToUpdate(model cockroach_models.ImportModel, e domain.Import) *cockroach_models.ImportModel {
	if e.UserID != "" {
		model.UserID = e.UserID
	}

	if e.Filename != "" {
		model.Filename = e.Filename
	}

	if e.Status != "" {
		model.Status = e.Status
	}

	if e.TotalRecords != 0 {
		model.TotalRecords = e.TotalRecords
	}

	if e.ErrorMessage != "" {
		model.ErrorMessage = e.ErrorMessage
	}

	return &model
}
