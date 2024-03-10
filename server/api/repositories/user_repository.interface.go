package repositories

import (
	db "github.com/RomainC75/todo2/db/sqlc"
	"github.com/gin-gonic/gin"
)

type UserRepositoryInterface interface {
	CreateUser(ctx *gin.Context, arg db.CreateUserParams) (db.User, error)
	FindUserByEmail(ctx *gin.Context, email string) (db.User, error)
}
