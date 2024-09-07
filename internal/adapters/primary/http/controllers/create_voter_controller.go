package controllers

import (
	"github.com/axel-andrade/opina-ai-api/internal/adapters/primary/http/presenters"
	"github.com/axel-andrade/opina-ai-api/internal/core/usecases/voter/create_voter"
	"github.com/gin-gonic/gin"
)

type CreateVoterController struct {
	Usecase   create_voter.CreateVoterUC
	Presenter presenters.CreateVoterPresenter
}

func BuildCreateVoterController(uc *create_voter.CreateVoterUC, ptr *presenters.CreateVoterPresenter) *CreateVoterController {
	return &CreateVoterController{Usecase: *uc, Presenter: *ptr}
}

// @Summary		Create a new voter
// @Description	Create a new voter
// @Tags			voters
// @Accept			json
// @Produce		json
// @Param			body	body		create_voter.CreateVoterInput	true	"Voter data"
// @Success		201		{object}	common_ptr.VoterFormatted
// @Router			/api/v1/voters [post]
func (ctrl *CreateVoterController) Handle(c *gin.Context) {
	inputMap := c.MustGet("body").(map[string]any)
	input := create_voter.CreateVoterInput{
		FullName:  inputMap["full_name"].(string),
		Cellphone: inputMap["cellphone"].(string),
	}

	result, err := ctrl.Usecase.Execute(input)
	output := ctrl.Presenter.Show(result, err)

	c.JSON(output.StatusCode, output.Data)
}
