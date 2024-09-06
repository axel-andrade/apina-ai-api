package cockroach_repositories

import (
	cockroach_mappers "github.com/axel-andrade/opina-ai-api/internal/adapters/secondary/database/cockroach/mappers"
	"github.com/axel-andrade/opina-ai-api/internal/core/domain"
)

type VoterCockroachRepository struct {
	*BaseCockroachRepository
	VoterMapper cockroach_mappers.VoterMapper
}

func BuildVoterRepository() *VoterCockroachRepository {
	return &VoterCockroachRepository{BaseCockroachRepository: BuildBaseRepository()}
}

func (r *VoterCockroachRepository) ExistsVoter(cellphone string) (bool, error) {
	q := r.getQueryOrTx()

	var count int64
	if err := q.Model(&domain.Voter{}).Where("cellphone = ?", cellphone).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *VoterCockroachRepository) CreateVoter(transaction *domain.Voter) (*domain.Voter, error) {
	model := r.VoterMapper.ToPersistence(*transaction)

	q := r.getQueryOrTx()

	if err := q.Create(model).Error; err != nil {
		return nil, err
	}

	return r.VoterMapper.ToDomain(*model), nil
}
