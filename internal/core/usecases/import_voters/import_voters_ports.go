package import_voters

import "github.com/axel-andrade/opina-ai-api/internal/core/domain"

type ImportVotersGateway interface {
	CreateImport(i *domain.Import) (*domain.Import, error)
	GetVotersByCellphones(cellphones []string) ([]*domain.Voter, error)
	CreateVoters(voters []*domain.Voter) error
	UpdateImport(i *domain.Import) (*domain.Import, error)
}

type ImportVotersInput struct {
	UserID   string
	Filename string
	Data     []byte
}

type ImportVotersOutput struct {
	Import *domain.Import
}
