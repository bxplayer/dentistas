package utils

import (
	modelPatient "dentistas/internal/patient"
	"errors"
	"github.com/gin-gonic/gin"
)

func ToPatient(ctx *gin.Context) (modelPatient.Patient, error) {
	var parsePatient modelPatient.Patient

	if err := ctx.ShouldBindJSON(&parsePatient); err != nil {
		return modelPatient.Patient{}, errors.New("Objeto no valido")
	}

	return parsePatient, nil
}

func ToPatientPatch(ctx *gin.Context) (modelPatient.PatientPatch, error) {
	var parsePatient modelPatient.PatientPatch

	if err := ctx.ShouldBindJSON(&parsePatient); err != nil {
		return modelPatient.PatientPatch{}, errors.New("Objeto no valido")
	}

	return parsePatient, nil
}
