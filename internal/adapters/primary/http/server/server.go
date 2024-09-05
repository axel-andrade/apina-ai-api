package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/axel-andrade/opina-ai-api/internal/adapters/primary/http/middlewares"
	"github.com/axel-andrade/opina-ai-api/internal/adapters/primary/http/routes"
	"github.com/axel-andrade/opina-ai-api/internal/infra"
)

type Server struct {
	port   string
	server *http.Server
}

// NewServer cria e retorna uma nova instância do servidor HTTP com configurações padrão de middleware.
func NewServer(port string) Server {
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return Server{
		port:   port,
		server: srv,
	}
}

func (s *Server) AddRoutes(d *infra.Dependencies) {
	routes.ConfigRoutes(s.server.Handler.(*http.ServeMux), d)
}

func (s *Server) Run() {
	log.Printf("Server starting on port %s", s.port)

	// Adiciona middlewares padrão
	handler := middlewares.Gzip(s.server.Handler)
	handler = middlewares.Cors(handler)
	handler = middlewares.RequestID(handler)
	handler = middlewares.SecurityHeaders(handler)
	handler = middlewares.Cache(handler, time.Minute)
	handler = middlewares.Logging(handler)
	handler = middlewares.Recovery(handler)

	s.server.Handler = handler

	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error starting server: %v\n", err)
	}
}

func (s *Server) Shutdown() {
	if s.server == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		log.Printf("Error shutting down server: %v\n", err)
	}

	log.Println("Server shutdown completed")
}
