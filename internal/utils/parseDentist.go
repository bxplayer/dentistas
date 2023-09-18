package utils

import (
	modelDentist "dentistas/internal/dentist"
	"errors"
	"github.com/gin-gonic/gin"
)

func ToDentist(ctx *gin.Context) (modelDentist.Dentist, error) {
	var parseDentist modelDentist.Dentist

	if err := ctx.ShouldBindJSON(&parseDentist); err != nil {
		return modelDentist.Dentist{}, errors.New("Objeto no valido")
	}

	return parseDentist, nil
}

func ToDentistPatch(ctx *gin.Context) (modelDentist.DentistPatch, error) {
	var parseDentist modelDentist.DentistPatch

	if err := ctx.ShouldBindJSON(&parseDentist); err != nil {
		return modelDentist.DentistPatch{}, errors.New("Objeto no valido")
	}

	return parseDentist, nil
}
