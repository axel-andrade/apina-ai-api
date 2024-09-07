package presenters

import (
	"encoding/json"
	"net/http"

	common_adapters "github.com/axel-andrade/opina-ai-api/internal/adapters/primary/http/common"
	common_ptr "github.com/axel-andrade/opina-ai-api/internal/adapters/primary/http/presenters/common"
	err_msg "github.com/axel-andrade/opina-ai-api/internal/core/domain/constants/errors"
	"github.com/axel-andrade/opina-ai-api/internal/core/usecases/voter/create_voter"
)

type CreateVoterPresenter struct {
	VoterPtr *common_ptr.VoterPresenter
}

func BuildCreateVoterPresenter() *CreateVoterPresenter {
	return &CreateVoterPresenter{VoterPtr: common_ptr.BuildVoterPresenter()}
}

func (p *CreateVoterPresenter) Show(result *create_voter.CreateVoterOutput, err error) common_adapters.OutputPort {
	if err != nil {
		return p.formatError(err)
	}

	fc := p.VoterPtr.Format(result.Voter)
	data, _ := json.Marshal(fc)

	return common_adapters.OutputPort{StatusCode: http.StatusCreated, Data: data}
}

func (p *CreateVoterPresenter) formatError(err error) common_adapters.OutputPort {
	errMsg := common_adapters.ErrorMessage{Message: err.Error()}

	switch err.Error() {
	case err_msg.CONTACT_FULL_NAME_REQUIRED, err_msg.CONTACT_CELLPHONE_REQUIRED, err_msg.INVALID_CELLPHONE:
		return common_adapters.OutputPort{StatusCode: http.StatusBadRequest, Data: errMsg}
	case err_msg.VOTER_ALREADY_EXISTS:
		return common_adapters.OutputPort{StatusCode: http.StatusConflict, Data: errMsg}
	default:
		return common_adapters.OutputPort{StatusCode: http.StatusBadRequest, Data: common_adapters.ErrorMessage{Message: err_msg.INTERNAL_SERVER_ERROR}}
	}
}
