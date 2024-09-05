package presenters

import (
	"encoding/json"
	"net/http"

	common_adapters "github.com/axel-andrade/opina-ai-api/internal/adapters/primary/http/common"
	common_ptr "github.com/axel-andrade/opina-ai-api/internal/adapters/primary/http/presenters/common"
	err_msg "github.com/axel-andrade/opina-ai-api/internal/core/domain/constants/errors"
	create_content "github.com/axel-andrade/opina-ai-api/internal/core/usecases/content/create"
)

type CreateContentPresenter struct {
	ContentPtr *common_ptr.ContentPresenter
}

func BuildCreateContentPresenter() *CreateContentPresenter {
	return &CreateContentPresenter{ContentPtr: common_ptr.BuildContentPresenter()}
}

func (p *CreateContentPresenter) Show(result *create_content.CreateContentOutput, err error) common_adapters.OutputPort {
	if err != nil {
		return p.formatError(err)
	}

	fc := p.ContentPtr.Format(result.Content)
	data, _ := json.Marshal(fc)

	return common_adapters.OutputPort{StatusCode: http.StatusCreated, Data: data}
}

func (p *CreateContentPresenter) formatError(err error) common_adapters.OutputPort {
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
