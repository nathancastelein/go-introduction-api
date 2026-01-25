package nethttp

import (
	"github.com/nathancastelein/go-introduction-api/examples/database/bob"
	"github.com/nathancastelein/go-introduction-api/shelter"
)

type petRepository struct{}

func (p petRepository) SelectAllPets() ([]shelter.Pet, error) {
	return bob.SelectAllPets()
}

func (p petRepository) SelectOnePet(id int) (*shelter.Pet, error) {
	return bob.SelectOnePet(id)
}

func (p petRepository) UpdatePetName(id int, newName string) error {
	return bob.UpdatePetName(id, newName)
}

func newPetRepository() shelter.PetRepository {
	return &petRepository{}
}
