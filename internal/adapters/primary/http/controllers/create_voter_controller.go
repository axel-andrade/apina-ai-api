package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/axel-andrade/opina-ai-api/internal/adapters/primary/http/presenters"
	"github.com/axel-andrade/opina-ai-api/internal/core/usecases/voter/create_voter"
)

type CreateVoterController struct {
	Usecase   create_voter.CreateVoterUC
	Presenter presenters.CreateVoterPresenter
}

func BuildCreateVoterController(uc *create_voter.CreateVoterUC, ptr *presenters.CreateVoterPresenter) *CreateVoterController {
	return &CreateVoterController{Usecase: *uc, Presenter: *ptr}
}

func (ctrl *CreateVoterController) Handle(w http.ResponseWriter, r *http.Request) {
	var input create_voter.CreateVoterInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	result, err := ctrl.Usecase.Execute(input)
	output := ctrl.Presenter.Show(result, err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(output.StatusCode)
	if data, ok := output.Data.([]byte); ok {
		_, _ = w.Write(data)
	} else {
		http.Error(w, "Invalid data format", http.StatusInternalServerError)
	}
}
