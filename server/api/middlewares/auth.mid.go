package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/RomainC75/todo2/utils"
	"github.com/gin-gonic/gin"
)

func IsAuth(isSocket bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		if isSocket {
			token = c.Request.URL.Query().Get("token")
			fmt.Printf("===========> SOCKETR")

		} else {
			var auth_header []string
			var ok bool
			auth_header, ok = c.Request.Header["Authorization"]
			if !ok || !strings.HasPrefix(auth_header[0], "Bearer") {
				c.JSON(http.StatusBadRequest, gin.H{"message": "token missing"})
				c.Abort()
				return
			}
			token = strings.Split(auth_header[0], " ")[1]
		}
		// claim, err := encrypt.GetClaimsFromToken(token)
		claims, err := utils.GetClaimsFromToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}
		utils.PrettyDisplay("CLAIMS", claims)

		c.Set("user_email", claims["email"])
		c.Set("user_id", claims["id"])
		c.Next()
	}
}
