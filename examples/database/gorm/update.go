package gorm

import (
	"github.com/nathancastelein/go-introduction-api/examples/database/connect"
)

func UpdatePetName(id int, newName string) error {
	db, err := connect.Gorm()
	if err != nil {
		return err
	}

	var pet Pet
	if err := db.First(&pet, id).Error; err != nil {
		return err
	}

	return db.Model(&pet).Update("name", newName).Error
}
