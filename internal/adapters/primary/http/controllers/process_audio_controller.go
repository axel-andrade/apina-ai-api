package controllers

import (
	"net/http"

	"github.com/axel-andrade/opina-ai-api/internal/adapters/primary/http/presenters"
	"github.com/axel-andrade/opina-ai-api/internal/core/usecases/content/process_audio"
)

type ProcessAudioController struct {
	Usecase   process_audio.ProcessAudioUC
	Presenter presenters.ProcessAudioPresenter
}

func BuildProcessAudioController(uc *process_audio.ProcessAudioUC, ptr *presenters.ProcessAudioPresenter) *ProcessAudioController {
	return &ProcessAudioController{Usecase: *uc, Presenter: *ptr}
}

func (ctrl *ProcessAudioController) Handle(w http.ResponseWriter, r *http.Request) {
	contentID := r.PathValue("contentID")

	input := process_audio.ProcessAudioInput{
		ContentID: contentID,
	}

	err := ctrl.Usecase.Execute(input)
	output := ctrl.Presenter.Show(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(output.StatusCode)

	if output.Data != nil {
		if data, ok := output.Data.([]byte); ok {
			_, _ = w.Write(data)
		} else {
			http.Error(w, "Invalid data format", http.StatusInternalServerError)
		}
	}
}
