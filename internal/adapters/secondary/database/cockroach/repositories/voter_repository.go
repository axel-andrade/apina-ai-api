package cockroach_repositories

import (
	cockroach_mappers "github.com/axel-andrade/opina-ai-api/internal/adapters/secondary/database/cockroach/mappers"
	"github.com/axel-andrade/opina-ai-api/internal/core/domain"
)

type VoterRepository struct {
	*BaseRepository
	VoterMapper cockroach_mappers.VoterMapper
}

func BuildVoterRepository() *VoterRepository {
	return &VoterRepository{BaseRepository: BuildBaseRepository()}
}

func (r *VoterRepository) CreateVoter(transaction *domain.Voter) (*domain.Voter, error) {
	model := r.VoterMapper.ToPersistence(*transaction)

	q := r.getQueryOrTx()

	if err := q.Create(model).Error; err != nil {
		return nil, err
	}

	return r.VoterMapper.ToDomain(*model), nil
}
