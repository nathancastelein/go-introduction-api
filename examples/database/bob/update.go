package bob

import (
	"context"

	"github.com/aarondl/opt/omit"
	"github.com/nathancastelein/go-introduction-api/examples/database/connect"
	"github.com/nathancastelein/go-introduction-api/examples/database/bob/models"
)

func UpdatePetName(id int, newName string) error {
	ctx := context.Background()
	db, err := connect.Bob()
	if err != nil {
		return err
	}

	databasePet, err := models.FindPet(ctx, db, int32(id))
	if err != nil {
		return err
	}

	setter := &models.PetSetter{
		Name: omit.From(newName),
	}

	return databasePet.Update(ctx, db, setter)
}
