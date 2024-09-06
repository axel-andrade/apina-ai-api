package routes

import (
	"net/http"

	"github.com/axel-andrade/opina-ai-api/internal/infra"
)

func createVoterHandler(d *infra.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			d.CreateVoterController.Handle(w, r)
			return
		}
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
