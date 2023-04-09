package repositories

import (
	"go-delivery/app/models"
	"go-delivery/config"
)

type UserRepository struct {
}

var db = config.GetDB()

func (ur *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
