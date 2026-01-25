package gorm

import (
	"github.com/nathancastelein/go-introduction-api/examples/database/connect"
)

func DeleteOnePet(id int) error {
	db, err := connect.Gorm()
	if err != nil {
		return err
	}

	var pet Pet
	if err := db.First(&pet, id).Error; err != nil {
		return err
	}

	return db.Delete(&pet).Error
}
