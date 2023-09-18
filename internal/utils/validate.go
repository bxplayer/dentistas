package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ValidateId(ctx *gin.Context) (int, error) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, err
	}

	if id <= 0 {
		return 0, errors.New("ID no válido")
	}
	return id, nil
}
