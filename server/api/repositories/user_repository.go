package repositories

import (
	db "github.com/RomainC75/todo2/db/sqlc"
	"github.com/gin-gonic/gin"
)

type UserRepository struct {
	DB *db.Store
}

func NewUserRepo() *UserRepository {
	return &UserRepository{
		DB: db.GetConnection(),
	}
}

func (userRepo *UserRepository) CreateUser(ctx *gin.Context, arg db.CreateUserParams) (db.User, error) {
	user, err := (*userRepo.DB).CreateUser(ctx, arg)
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}

func (userRepo *UserRepository) FindUserByEmail(email string) (db.User, error) {
	var foundUser db.User
	// if result := userRepo.DB.Where("email = ?", email).First(&foundUser); result.RowsAffected == 0 {
	// 	return db.User{}, errors.New("no user found")
	// }
	return foundUser, nil
}
