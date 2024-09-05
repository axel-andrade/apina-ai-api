package presenters

import (
	"encoding/json"
	"net/http"

	common_adapters "github.com/axel-andrade/opina-ai-api/internal/adapters/primary/http/common"
	err_msg "github.com/axel-andrade/opina-ai-api/internal/core/domain/constants/errors"
)

type ProcessAudioPresenter struct{}

func BuildProcessAudioPresenter() *ProcessAudioPresenter {
	return &ProcessAudioPresenter{}
}

func (p *ProcessAudioPresenter) Show(err error) common_adapters.OutputPort {
	return common_adapters.OutputPort{StatusCode: http.StatusAccepted, Data: nil}
}

func (p *ProcessAudioPresenter) formatError(err error) common_adapters.OutputPort {
	errMsg := common_adapters.ErrorMessage{Message: err.Error()}
	data, _ := json.Marshal(errMsg)

	switch err.Error() {
	case err_msg.NAME_IS_EMPTY, err_msg.DESCRIPTION_IS_EMPTY, err_msg.KIND_IS_EMPTY, err_msg.LANGUAGE_IS_EMPTY, err_msg.FILE_IS_EMPTY, err_msg.TEXT_IS_EMPTY:
		return common_adapters.OutputPort{StatusCode: http.StatusBadRequest, Data: data}
	default:
		errMsg.Message = err_msg.INTERNAL_SERVER_ERROR
		data, _ = json.Marshal(errMsg)
		return common_adapters.OutputPort{StatusCode: http.StatusBadRequest, Data: data}
	}
}
