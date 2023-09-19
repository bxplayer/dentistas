package patientHandler

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


// GetPatientById godoc
// @Summary Get Patient By Id
// @Tags Patient
// @Description get patient by Id
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} patient.Patient
// @Router /patient/{id} [get]
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

// CreatePatient godoc
// @Summary Create Patient
// @Tags Patient
// @Description create patient
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body patient.Patient true "Patient"
// @Success 201 {object} patient.Patient
// @Router /patient [post]
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

// UpdatePatient godoc
// @Summary Update Patient
// @Tags Patient
// @Description update patient
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param patient body request true "Patient to update"
// @Success 201 {object} web.Response
// @Router /patient/ [put]
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

// PatchPatient godoc
// @Summary Patch Patient
// @Tags Patient
// @Description patch patient
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param patient body request true "Patient to patch"
// @Success 201 {object} web.Response
// @Router /patient [patch]
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

// DeletePatiend godoc
// @Summary Delete Patient
// @Tags Patient
// @Description delete patient
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id query string true "id"
// @Success 204 {object} web.Response
// @Router /patient [delete]
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
