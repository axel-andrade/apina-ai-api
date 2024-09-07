package create_voter

import (
	"github.com/axel-andrade/opina-ai-api/internal/core/domain"
)

type CreateVoterGateway interface {
	ExistsVoter(cellphone string) (bool, error)
	CreateVoter(v *domain.Voter) (*domain.Voter, error)
}

type CreateVoterInput struct {
	FullName  string `json:"full_name" example:"John Doe" description:"The full name of the voter"`
	Cellphone string `json:"cellphone" example:"553199999999" description:"The cellphone number of the voter"`
}

type CreateVoterOutput struct {
	Voter *domain.Voter
}
