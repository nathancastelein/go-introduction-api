package bob

import (
	"context"

	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
	"github.com/nathancastelein/go-introduction-api/examples/database/connect"
	"github.com/nathancastelein/go-introduction-api/examples/database/bob/models"
	"github.com/nathancastelein/go-introduction-api/shelter"
)

func InsertNewPet(pet shelter.Pet) (shelter.Pet, error) {
	ctx := context.Background()
	db, err := connect.Bob()
	if err != nil {
		return shelter.Pet{}, err
	}

	category, err := models.Categories.Query(
		models.SelectWhere.Categories.Name.EQ(pet.Category),
	).One(ctx, db)
	if err != nil {
		return shelter.Pet{}, err
	}

	setter := &models.PetSetter{
		Name:     omit.From(pet.Name),
		Category: omitnull.From(category.ID),
	}

	databasePet, err := models.Pets.Insert(setter).One(ctx, db)
	if err != nil {
		return shelter.Pet{}, err
	}

	return shelter.Pet{
		Id:       int(databasePet.ID),
		Name:     databasePet.Name,
		Category: category.Name,
	}, nil
}
