package cockroach_mappers

import (
	cockroach_models "github.com/axel-andrade/opina-ai-api/internal/adapters/secondary/database/cockroach/models"
	"github.com/axel-andrade/opina-ai-api/internal/core/domain"
)

type VoterMapper struct {
	BaseMapper
}

func (v *VoterMapper) ToDomain(model cockroach_models.VoterModel) *domain.Voter {
	return &domain.Voter{
		Base:      *v.BaseMapper.toDomain(model.BaseModel),
		FullName:  model.FullName,
		Cellphone: model.Cellphone,
	}
}

func (v *VoterMapper) ToPersistence(e domain.Voter) *cockroach_models.VoterModel {
	return &cockroach_models.VoterModel{
		BaseModel: *v.BaseMapper.toPersistence(e.Base),
		FullName:  e.FullName,
		Cellphone: e.Cellphone,
	}
}

func (v *VoterMapper) ToUpdate(model cockroach_models.VoterModel, e domain.Voter) *cockroach_models.VoterModel {
	if e.FullName != "" {
		model.FullName = e.FullName
	}

	if e.Cellphone != "" {
		model.Cellphone = e.Cellphone
	}

	return &model
}
