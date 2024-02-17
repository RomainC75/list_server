package services

import "github.com/RomainC75/todo2/api/dto/requests"

type UserServiceInterface interface {
	CreateUserSrv(user requests.SignupRequest)
}
