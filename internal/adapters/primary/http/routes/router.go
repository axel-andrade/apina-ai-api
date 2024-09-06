package routes

import (
	"net/http"

	"github.com/axel-andrade/opina-ai-api/internal/infra"
)

func ConfigRoutes(mux *http.ServeMux, d *infra.Dependencies) {
	mux.HandleFunc("/healthcheck", healthcheckHandler)
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("path/to/swagger/files"))))

	mux.HandleFunc("/api/v1/voters", createVoterHandler(d))
}
