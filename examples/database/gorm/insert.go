package gorm

import (
	"github.com/nathancastelein/go-introduction-api/examples/database/connect"
	"github.com/nathancastelein/go-introduction-api/shelter"
)

func InsertNewPet(pet shelter.Pet) (shelter.Pet, error) {
	db, err := connect.Gorm()
	if err != nil {
		return shelter.Pet{}, err
	}

	var category Category
	if err := db.Where("name = ?", pet.Category).First(&category).Error; err != nil {
		return shelter.Pet{}, err
	}

	dbPet := Pet{
		Name:       pet.Name,
		CategoryID: &category.ID,
	}

	if err := db.Create(&dbPet).Error; err != nil {
		return shelter.Pet{}, err
	}

	return shelter.Pet{
		Id:       int(dbPet.ID),
		Name:     dbPet.Name,
		Category: category.Name,
	}, nil
}
