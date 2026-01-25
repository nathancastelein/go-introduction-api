package nethttp

import (
	"log/slog"
	"net/http"

	"github.com/nathancastelein/go-introduction-api/shelter"
)

type Server struct {
	shelter shelter.PetRepository
	*http.ServeMux
}

func NewServer(shelter shelter.PetRepository) *Server {
	server := &Server{
		shelter: shelter,
	}

	server.ServeMux = http.NewServeMux()

	server.HandleFunc("GET /pet", server.ListPetsHandler)
	server.HandleFunc("GET /pet/{id}", server.GetOnePetHandler)
	server.HandleFunc("POST /pet/{id}/rename", server.UpdatePetHandler)

	return server
}

func (s *Server) Run(port string) error {
	if port == "" {
		port = "8080"
	}
	server := &http.Server{
		Addr:    ":" + port,
		Handler: s,
	}

	slog.Info("starting server...")
	return server.ListenAndServe()
}
