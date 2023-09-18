package handler

import (
	"dentistas/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PatientHandler struct {
	patientGetter  PatientGetters
	patientCreator PatientCreator
	patientDelete  PatientDelete
	patientUpdate  PatientUpdates
}

func NewPatientHandler(
	getter PatientGetters,
	creator PatientCreator,
	delete PatientDelete,
	update PatientUpdates,
) *PatientHandler {
	return &PatientHandler{
		patientGetter:  getter,
		patientCreator: creator,
		patientDelete:  delete,
		patientUpdate:  update,
	}
}

func (p *PatientHandler) GetPatientByID(ctx *gin.Context) {

	id, err := utils.ValidateId(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	Patient, err := p.patientGetter.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	ctx.JSON(http.StatusOK, Patient)
}

func (p *PatientHandler) Create(ctx *gin.Context) {

	newPatient, err := utils.ToPatient(ctx)
	if err != nil {
		return
	}

	response, err := p.patientCreator.Create(newPatient)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

func (p *PatientHandler) Update(ctx *gin.Context) {
	id, err := utils.ValidateId(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	newPatient, err := utils.ToPatient(ctx)
	if err != nil {
		return
	}

	response, err := p.patientUpdate.Update(id, newPatient)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

func (p *PatientHandler) Patch(ctx *gin.Context) {

	id, err := utils.ValidateId(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	newPatient, err := utils.ToPatientPatch(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response, err := p.patientUpdate.Patch(id, newPatient)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

func (p *PatientHandler) Delete(ctx *gin.Context) {
	id, err := utils.ValidateId(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = p.patientDelete.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
