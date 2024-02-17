package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/RomainC75/todo2/api/dto/requests"
	"github.com/RomainC75/todo2/api/dto/responses"
	"github.com/RomainC75/todo2/api/repositories"
	"github.com/RomainC75/todo2/config"
	"github.com/RomainC75/todo2/data/models"
	"github.com/RomainC75/todo2/utils"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserSrv struct {
	userRepository repositories.UserRepositoryInterface
}

func NewUserSrv() *UserSrv {
	return &UserSrv{
		userRepository: repositories.NewUserRepo(),
	}
}

func (userSrv *UserSrv) CreateUserSrv(user requests.SignupRequest) error {
	foundUser, err := userSrv.userRepository.FindUserByEmail(user.Email)
	fmt.Println("==> found user : ", foundUser, err)
	if err == nil {
		return errors.New("email already used")
	}

	b, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error trying to encrypt the password")
	}

	userModel := models.User{
		Email:    user.Email,
		Password: string(b),
	}

	createdUser, err := userSrv.userRepository.CreateUser(userModel)
	utils.PrettyDisplay("createdUser", createdUser)
	if err != nil {
		return err
	}
	return nil
}

func (UserSrv *UserSrv) LoginUserSrv(user requests.LoginRequest) (responses.AuthLoginResponse, error) {
	foundUser, err := UserSrv.userRepository.FindUserByEmail(user.Email)
	if err != nil {
		return responses.AuthLoginResponse{}, errors.New("email not valid")
	}
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	utils.PrettyDisplay("password : ", foundUser.Password)
	if err != nil {
		return responses.AuthLoginResponse{}, errors.New("password not valid")
	}

	validityTime := config.Get().Jwt.JwtValidityTime
	claim := utils.UserClaims{
		Id:    foundUser.ID,
		Email: foundUser.Email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(validityTime)).Unix(),
		},
	}

	signedAccessToken, err := utils.NewAccessToken(claim)
	if err != nil {
		fmt.Println(err.Error())
		return responses.AuthLoginResponse{}, errors.New("error trying to create token")
	}

	return responses.AuthLoginResponse{
		Id:    foundUser.ID,
		Email: foundUser.Email,
		Token: signedAccessToken,
	}, nil
}
