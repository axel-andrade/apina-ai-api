package common_ptr

import (
	"github.com/axel-andrade/opina-ai-api/internal/core/domain"
)

type VoterFormatted struct {
	ID        string `json:"id" example:"123" description:"O ID do contéudo"`
	FullName  string `json:"full_name" example:"John Doe" description:"O nome do eleitor"`
	Cellphone string `json:"cellphone" example:"+5511999999999" description:"O telefone do eleitor"`
	CreatedAt string `json:"created_at" example:"2021-01-01T00:00:00Z" description:"Data de criação do contéudo"`
	UpdatedAt string `json:"updated_at" example:"2021-01-01T00:00:00Z" description:"Data de atualização do contéudo"`
}

type VoterPresenter struct{}

func BuildVoterPresenter() *VoterPresenter {
	return &VoterPresenter{}
}

func (ptr *VoterPresenter) Format(voter *domain.Voter) VoterFormatted {
	return VoterFormatted{
		ID:        voter.ID,
		FullName:  voter.FullName,
		Cellphone: voter.Cellphone,
		CreatedAt: voter.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: voter.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}

func (ptr *VoterPresenter) FormatList(voters []domain.Voter) []VoterFormatted {
	var votersFormatted []VoterFormatted = make([]VoterFormatted, 0)

	for _, voter := range voters {
		votersFormatted = append(votersFormatted, ptr.Format(&voter))
	}

	return votersFormatted
}
