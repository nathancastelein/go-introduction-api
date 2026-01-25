package bob

import (
	"context"

	"github.com/nathancastelein/go-introduction-api/examples/database/connect"
	"github.com/nathancastelein/go-introduction-api/examples/database/bob/models"
	"github.com/nathancastelein/go-introduction-api/shelter"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

func SelectAllPets() ([]shelter.Pet, error) {
	ctx := context.Background()
	db, err := connect.Bob()
	if err != nil {
		return nil, err
	}

	pets, err := models.Pets.Query().All(ctx, db)
	if err != nil {
		return nil, err
	}

	shelterPets := make([]shelter.Pet, len(pets))
	for i, pet := range pets {
		if err := pet.LoadCategory(ctx, db); err != nil {
			return nil, err
		}
		shelterPets[i] = shelter.Pet{
			Id:       int(pet.ID),
			Name:     pet.Name,
			Category: pet.R.Category.Name,
		}
	}

	return shelterPets, nil
}

func SelectOnePet(id int) (*shelter.Pet, error) {
	ctx := context.Background()
	db, err := connect.Bob()
	if err != nil {
		return nil, err
	}

	pet, err := models.FindPet(ctx, db, int32(id))
	if err != nil {
		return nil, err
	}

	if err := pet.LoadCategory(ctx, db); err != nil {
		return nil, err
	}

	return &shelter.Pet{
		Id:       int(pet.ID),
		Name:     pet.Name,
		Category: pet.R.Category.Name,
	}, nil
}

func SelectByCategories(categories ...string) ([]shelter.Pet, error) {
	ctx := context.Background()
	db, err := connect.Bob()
	if err != nil {
		return nil, err
	}

	databaseCategories, err := models.Categories.Query(
		models.SelectWhere.Categories.Name.In(categories...),
	).All(ctx, db)
	if err != nil {
		return nil, err
	}

	categoriesArgs := make([]bob.Expression, len(databaseCategories))
	for i, category := range databaseCategories {
		categoriesArgs[i] = psql.Arg(category.ID)
	}

	pets, err := models.Pets.Query(
		sm.Where(models.Pets.Columns.Category.In(categoriesArgs...)),
	).All(ctx, db)
	if err != nil {
		return nil, err
	}

	shelterPets := make([]shelter.Pet, len(pets))
	for i, pet := range pets {
		if err := pet.LoadCategory(ctx, db); err != nil {
			return nil, err
		}
		shelterPets[i] = shelter.Pet{
			Id:       int(pet.ID),
			Name:     pet.Name,
			Category: pet.R.Category.Name,
		}
	}

	return shelterPets, nil
}
