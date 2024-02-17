package repositories

import (
	"errors"
	"fmt"

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

func (userRepo *UserRepository) CreateUser(user models.User) (models.User, error) {
	fmt.Println("try to find")
	var newUser models.User
	if result := userRepo.DB.Create(&user).Scan(&newUser); result.RowsAffected == 0 {
		fmt.Println("affected rows : ", result.RowsAffected)
		return models.User{}, errors.New("error trying to create a new user :-(")
	}
	fmt.Println("affected rows : ", newUser)
	return newUser, nil
}

func (userRepo *UserRepository) FindUserByEmail(email string) (models.User, error) {
	var foundUser models.User
	if result := userRepo.DB.Where("email = ?", email).First(&foundUser); result.RowsAffected == 0 {
		return models.User{}, errors.New("no user found")
	}
	return foundUser, nil
}
