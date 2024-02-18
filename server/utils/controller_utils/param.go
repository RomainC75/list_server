package controller_utils

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIdFromParam(c *gin.Context, paramName string) (int, error) {
	id := c.Param(paramName)
	intId, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.New("id not valid")
	}
	return intId, nil
}
