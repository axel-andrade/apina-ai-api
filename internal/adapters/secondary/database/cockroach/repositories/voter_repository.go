package cockroach_repositories

import (
	cockroach_mappers "github.com/axel-andrade/opina-ai-api/internal/adapters/secondary/database/cockroach/mappers"
	cockroach_models "github.com/axel-andrade/opina-ai-api/internal/adapters/secondary/database/cockroach/models"
	"github.com/axel-andrade/opina-ai-api/internal/core/domain"
)

type VoterCockroachRepository struct {
	*BaseCockroachRepository
	VoterMapper cockroach_mappers.VoterMapper
}

func BuildCockroachVoterRepository() *VoterCockroachRepository {
	return &VoterCockroachRepository{BaseCockroachRepository: BuildCockroachBaseRepository()}
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

func (r *VoterCockroachRepository) CreateVoters(voters []*domain.Voter) error {
	q := r.getQueryOrTx()

	// Create a list of persistence models
	var voterModels []cockroach_models.VoterModel

	// Iterating over the domain voters list and using the mapper to convert them
	for _, voter := range voters {
		model := r.VoterMapper.ToPersistence(*voter)
		voterModels = append(voterModels, *model)
	}

	// Use GORM's batch insert to insert all voters at once
	if err := q.Create(&voterModels).Error; err != nil {
		return err
	}

	return nil
}

func (r *VoterCockroachRepository) GetVotersByCellphones(cellphones []string) ([]*domain.Voter, error) {
	q := r.getQueryOrTx()

	// List to store the persistence models
	var models []cockroach_models.VoterModel

	// Use the IN clause to search voters based on multiple cellphones
	if err := q.Model(&cockroach_models.VoterModel{}).Where("cellphone IN ?", cellphones).Find(&models).Error; err != nil {
		return nil, err
	}

	// Use the mapper to convert persistence models into domain objects
	var voters []*domain.Voter
	for _, model := range models {
		voter := r.VoterMapper.ToDomain(model)
		voters = append(voters, voter)
	}

	return voters, nil
}
