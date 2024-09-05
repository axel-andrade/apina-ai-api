package routes

import (
	"net/http"

	"github.com/axel-andrade/opina-ai-api/internal/infra"
)

func createContentHandler(d *infra.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			d.CreateContentController.Handle(w, r)
			return
		}
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func processAudioHandler(d *infra.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			d.ProcessAudioController.Handle(w, r)
			return
		}
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
