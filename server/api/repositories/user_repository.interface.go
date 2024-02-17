package repositories

import "github.com/RomainC75/todo2/data/models"

type UserRepositoryInterface interface {
	CreateUser(user models.User) (models.User, error)
	FindUserByEmail(email string) (models.User, error)
}
