package create_voter

import (
	"github.com/axel-andrade/opina-ai-api/internal/core/domain"
)

type CreateVoterGateway interface {
	ExistsVoter(cellphone string) (bool, error)
	CreateVoter(v *domain.Voter) (*domain.Voter, error)
}

type CreateVoterInput struct {
	FullName  string
	Cellphone string
}

type CreateVoterOutput struct {
	Voter *domain.Voter
}
