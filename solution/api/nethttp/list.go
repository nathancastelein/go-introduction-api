package nethttp

import (
	"net/http"
)

func (s *Server) ListPetsHandler(w http.ResponseWriter, r *http.Request) {
	pets, err := s.shelter.SelectAllPets()
	if err != nil {
		WriteJSONResult(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSONResult(w, http.StatusOK, pets)
}
