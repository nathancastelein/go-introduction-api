package nethttp

import (
	"net/http"
)

func (s *Server) GetOnePetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := PathInt(r, "id")
	if err != nil {
		WriteJSONResult(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
		return
	}

	pet, err := s.shelter.SelectOnePet(id)
	if err != nil {
		WriteJSONResult(w, http.StatusNotFound, map[string]string{"error": "pet not found"})
		return
	}

	WriteJSONResult(w, http.StatusOK, pet)
}
