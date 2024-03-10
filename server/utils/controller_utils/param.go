package controller_utils

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIdFromParam(c *gin.Context, paramName string) (int32, error) {
	id := c.Param(paramName)
	intId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return 0, errors.New("id not valid")
	}
	return int32(intId), nil
}
