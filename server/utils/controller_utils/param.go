package controller_utils

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetIdFromParam(c *gin.Context, paramName string) (int32, error) {
	id := c.Param(paramName)
	intId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return 0, errors.New("id not valid")
	}
	return int32(intId), nil
}

func GetUUIDFromParam(c *gin.Context, paramName string) (uuid.UUID, error) {
	id := c.Param(paramName)
	uuid, err := uuid.Parse(id)
	if err != nil {
		return uuid, errors.New("uuid not valid")
	}
	return uuid, nil
}
