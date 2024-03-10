package services

import (
	"github.com/RomainC75/todo2/api/dto/requests"
	"github.com/gin-gonic/gin"
)

type UserServiceInterface interface {
	CreateUserSrv(ctx *gin.Context, user requests.SignupRequest) error
	LoginUserSrv(user requests.LoginRequest) error
}
