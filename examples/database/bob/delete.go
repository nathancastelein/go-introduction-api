package bob

import (
	"context"

	"github.com/nathancastelein/go-introduction-api/examples/database/connect"
	"github.com/nathancastelein/go-introduction-api/examples/database/bob/models"
)

func DeleteOnePet(id int) error {
	ctx := context.Background()
	db, err := connect.Bob()
	if err != nil {
		return err
	}

	databasePet, err := models.FindPet(ctx, db, int32(id))
	if err != nil {
		return err
	}

	return databasePet.Delete(ctx, db)
}
