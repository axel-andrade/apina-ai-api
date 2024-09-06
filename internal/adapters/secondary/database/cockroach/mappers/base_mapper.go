package cockroach_mappers

import (
	cockroach_models "github.com/axel-andrade/opina-ai-api/internal/adapters/secondary/database/cockroach/models"
	"github.com/axel-andrade/opina-ai-api/internal/core/domain"
)

type BaseMapper struct{}

func (m *BaseMapper) toDomain(model cockroach_models.BaseModel) *domain.Base {
	return &domain.Base{
		ID:        model.ID,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func (m *BaseMapper) toPersistence(e domain.Base) *cockroach_models.BaseModel {
	return &cockroach_models.BaseModel{
		ID:        e.ID,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
