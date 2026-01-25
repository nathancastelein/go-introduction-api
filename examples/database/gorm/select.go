package gorm

import (
	"github.com/nathancastelein/go-introduction-api/examples/database/connect"
	"github.com/nathancastelein/go-introduction-api/shelter"
)

func SelectAllPets() ([]shelter.Pet, error) {
	db, err := connect.Gorm()
	if err != nil {
		return nil, err
	}

	var pets []Pet
	if err := db.Preload("Category").Find(&pets).Error; err != nil {
		return nil, err
	}

	shelterPets := make([]shelter.Pet, len(pets))
	for i, pet := range pets {
		shelterPets[i] = shelter.Pet{
			Id:       int(pet.ID),
			Name:     pet.Name,
			Category: pet.Category.Name,
		}
	}

	return shelterPets, nil
}

func SelectOnePet(id int) (*shelter.Pet, error) {
	db, err := connect.Gorm()
	if err != nil {
		return nil, err
	}

	var pet Pet
	if err := db.Preload("Category").First(&pet, id).Error; err != nil {
		return nil, err
	}

	return &shelter.Pet{
		Id:       int(pet.ID),
		Name:     pet.Name,
		Category: pet.Category.Name,
	}, nil
}

func SelectByCategories(categories ...string) ([]shelter.Pet, error) {
	db, err := connect.Gorm()
	if err != nil {
		return nil, err
	}

	var dbCategories []Category
	if err := db.Where("name IN ?", categories).Find(&dbCategories).Error; err != nil {
		return nil, err
	}

	categoryIDs := make([]int32, len(dbCategories))
	for i, cat := range dbCategories {
		categoryIDs[i] = cat.ID
	}

	var pets []Pet
	if err := db.Preload("Category").Where("category IN ?", categoryIDs).Find(&pets).Error; err != nil {
		return nil, err
	}

	shelterPets := make([]shelter.Pet, len(pets))
	for i, pet := range pets {
		shelterPets[i] = shelter.Pet{
			Id:       int(pet.ID),
			Name:     pet.Name,
			Category: pet.Category.Name,
		}
	}

	return shelterPets, nil
}
