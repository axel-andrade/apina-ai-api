package create_voter

import (
	"fmt"
	"log"

	"github.com/axel-andrade/opina-ai-api/internal/core/domain"
	err_msg "github.com/axel-andrade/opina-ai-api/internal/core/domain/constants/errors"
)

type CreateVoterUC struct {
	Gateway CreateVoterGateway
}

func BuildCreateVoterUC(g CreateVoterGateway) *CreateVoterUC {
	return &CreateVoterUC{g}
}

func (bs *CreateVoterUC) Execute(input CreateVoterInput) (*CreateVoterOutput, error) {
	log.Println("Building voter")
	voter, err := domain.BuildNewVoter(input.FullName, input.Cellphone)

	if err != nil {
		return nil, err
	}

	log.Printf("Checking if voter exists")
	if exists, err := bs.Gateway.ExistsVoter(voter.Cellphone); err != nil {
		return nil, err
	} else if exists {
		return nil, fmt.Errorf(err_msg.VOTER_ALREADY_EXISTS)
	}

	log.Printf("Creating voter")
	result, err := bs.Gateway.CreateVoter(voter)

	if err != nil {
		return nil, err
	}

	return &CreateVoterOutput{Voter: result}, nil
}
