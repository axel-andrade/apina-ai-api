package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/axel-andrade/opina-ai-api/internal/adapters/primary/http/presenters"
	create_content "github.com/axel-andrade/opina-ai-api/internal/core/usecases/content/create"
)

type CreateContentController struct {
	Usecase   create_content.CreateContentUC
	Presenter presenters.CreateContentPresenter
}

func BuildCreateContentController(uc *create_content.CreateContentUC, ptr *presenters.CreateContentPresenter) *CreateContentController {
	return &CreateContentController{Usecase: *uc, Presenter: *ptr}
}

func (ctrl *CreateContentController) Handle(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	description := r.FormValue("description")
	kind := r.FormValue("kind")

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "unable to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	ext := filepath.Ext(header.Filename)
	if ext != ".wav" && ext != ".mp3" && ext != ".m4a" && ext != ".ogg" {
		http.Error(w, "Unsupported file format", http.StatusBadRequest)
		return
	}

	input := create_content.CreateContentInput{
		Name:        name,
		Description: description,
		Kind:        kind,
		Language:    "en",
		File:        file,
		Filename:    header.Filename,
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
