package repositories

import (
	"fmt"
	"time"

	db "github.com/RomainC75/todo2/db/sqlc"
	"github.com/gin-gonic/gin"
)

type UserRepository struct {
	Store *db.Store
}

func NewUserRepo() *UserRepository {
	return &UserRepository{
		Store: db.GetConnection(),
	}
}

// sql.ErrNoRows

func (userRepo *UserRepository) CreateUser(ctx *gin.Context, arg db.CreateUserParams) (db.User, error) {
	arg.CreatedAt = time.Now()
	arg.UpdatedAt = arg.CreatedAt
	fmt.Println("=> pre insert db", arg)
	user, err := (*userRepo.Store).CreateUser(ctx, arg)
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}

func (userRepo *UserRepository) FindUserByEmail(ctx *gin.Context, email string) (db.User, error) {
	foundUser, err := (*userRepo.Store).GetUserByEmail(ctx, email)
	if err != nil {
		return db.User{}, err
	}
	return foundUser, nil
}
