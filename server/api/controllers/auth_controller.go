package controllers

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/RomainC75/todo2/api/dto/requests"
	"github.com/RomainC75/todo2/api/dto/responses"
	"github.com/RomainC75/todo2/api/services"
	"github.com/gin-gonic/gin"
)

type AuthCtrl struct {
	userSrv services.UserSrv
}

func NewAuthCtrl() *AuthCtrl {
	return &AuthCtrl{
		userSrv: *services.NewUserSrv(),
	}
}

func (authCtrl *AuthCtrl) HandleSignup(c *gin.Context) {
	var newUser requests.SignupRequest
	if err := c.ShouldBind(&newUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	err := authCtrl.userSrv.CreateUserSrv(newUser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"user created": newUser.Email})
}

func (authCtrl *AuthCtrl) HandleLogin(c *gin.Context) {
	var loginInfos requests.LoginRequest
	if err := c.ShouldBind(&loginInfos); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	loginResponse, err := authCtrl.userSrv.LoginUserSrv(loginInfos)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusUnprocessableEntity, gin.H{"authentication_details": loginResponse})
}

func (AuthCtrl *AuthCtrl) HandleVerify(c *gin.Context) {
	id, _ := c.Get("user_id")
	email, _ := c.Get("user_email")

	fmt.Println("==========>Hanfler", reflect.TypeOf(id))
	verifyResponse := responses.AuthVerifyResponse{
		Id:    id.(float64),
		Email: email.(string),
	}
	c.JSON(http.StatusAccepted, gin.H{"infos": verifyResponse})
}
