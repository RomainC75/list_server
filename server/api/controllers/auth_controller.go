package controllers

import (
	"net/http"

	"github.com/RomainC75/todo2/api/dto/requests"
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
	}

	authCtrl.userSrv.CreateUserSrv(newUser)

	c.JSON(http.StatusAccepted, gin.H{"user": newUser})
}
