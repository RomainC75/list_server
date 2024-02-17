package services

import (
	"fmt"

	"github.com/RomainC75/todo2/api/dto/requests"
	"github.com/RomainC75/todo2/api/repositories"
)

type UserSrv struct {
	userRepository repositories.UserRepositoryInterface
}

func NewUserSrv() *UserSrv {
	return &UserSrv{
		userRepository: repositories.NewUserRepo(),
	}
}

func (userSrv *UserSrv) CreateUserSrv(user requests.SignupRequest) {
	fmt.Printf("==> service : create user service")
	fmt.Println("found dta : ", user)
}
