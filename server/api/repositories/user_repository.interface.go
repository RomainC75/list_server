package repositories

import (
	"errors"

	"github.com/RomainC75/todo2/data/database"
	"github.com/RomainC75/todo2/data/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepo() *UserRepository {
	return &UserRepository{
		DB: database.GetConnection(),
	}
}

func (UserRepo *UserRepository) CreateUser(user models.User) (models.User, error) {
	var newUser models.User
	if result := UserRepo.DB.Create(&user).Scan(&newUser); result.RowsAffected == 0 {
		return models.User{}, errors.New("error trying to create a new user :-(")
	}
	return newUser, nil
}
