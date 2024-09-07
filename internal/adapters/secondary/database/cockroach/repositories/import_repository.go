package cockroach_repositories

import (
	cockroach_mappers "github.com/axel-andrade/opina-ai-api/internal/adapters/secondary/database/cockroach/mappers"
	cockroach_models "github.com/axel-andrade/opina-ai-api/internal/adapters/secondary/database/cockroach/models"
	"github.com/axel-andrade/opina-ai-api/internal/core/domain"
)

type ImportCockroachRepository struct {
	*BaseCockroachRepository
	ImportMapper cockroach_mappers.ImportMapper
}

func BuildCockroachImportRepository() *ImportCockroachRepository {
	return &ImportCockroachRepository{BaseCockroachRepository: BuildCockroachBaseRepository()}
}

func (r *ImportCockroachRepository) CreateImport(i *domain.Import) (*domain.Import, error) {
	model := r.ImportMapper.ToPersistence(*i)

	q := r.getQueryOrTx()

	if err := q.Create(model).Error; err != nil {
		return nil, err
	}

	return r.ImportMapper.ToDomain(*model), nil
}

func (r *ImportCockroachRepository) UpdateImport(i *domain.Import) (*domain.Import, error) {
	q := r.getQueryOrTx()

	model := r.ImportMapper.ToPersistence(*i)
	if err := q.Model(&cockroach_models.ImportModel{}).Where("id = ?", i.ID).Updates(model).Error; err != nil {
		return nil, err
	}

	return r.ImportMapper.ToDomain(*model), nil
}
