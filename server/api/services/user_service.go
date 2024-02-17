package services

import (
	"errors"
	"fmt"

	"github.com/RomainC75/todo2/api/dto/requests"
	"github.com/RomainC75/todo2/api/repositories"
	"github.com/RomainC75/todo2/data/models"
)

type UserSrv struct {
	userRepository repositories.UserRepositoryInterface
}

func NewUserSrv() *UserSrv {
	return &UserSrv{
		userRepository: repositories.NewUserRepo(),
	}
}

func (userSrv *UserSrv) CreateUserSrv(user requests.SignupRequest) (models.User, error) {
	foundUser, err := userSrv.userRepository.FindUserByEmail(user.Email)
	fmt.Println("==> found user : ", foundUser, err)
	if err == nil {
		return models.User{}, errors.New("email already used")
	}

	userModel := models.User{
		Email:    user.Email,
		Password: user.Password,
	}

	fmt.Printf("............")

	createdUser, err := userSrv.userRepository.CreateUser(userModel)
	if err != nil {
		return models.User{}, err
	}
	return createdUser, nil
}
