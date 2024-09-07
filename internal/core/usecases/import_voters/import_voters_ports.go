package import_voters

import "github.com/axel-andrade/opina-ai-api/internal/core/domain"

type ImportVotersGateway interface {
	GetVotersByCellphones(cellphones []string) ([]*domain.Voter, error)
	CreateVoters(voters []*domain.Voter) error
}

type ImportVotersInput struct {
	UserID string
	Data   []byte
}

type ImportVotersOutput struct {
	Import *domain.Import
}
