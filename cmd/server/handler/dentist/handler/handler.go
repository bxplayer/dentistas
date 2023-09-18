package handler

import (
	"dentistas/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DentistHandler struct {
	dentistGetter  DentistGetters
	dentistCreator DentistCreator
	dentistDelete  DentistDelete
	dentistUpdate  DentistUpdates
}

func NewDentistHandler(
	getter DentistGetters,
	creator DentistCreator,
	delete DentistDelete,
	update DentistUpdates,
) *DentistHandler {
	return &DentistHandler{
		dentistGetter:  getter,
		dentistCreator: creator,
		dentistDelete:  delete,
		dentistUpdate:  update,
	}
}

// GetDentistByID godoc
// @Summary Get dentist by ID
// @Description Get dentist by ID
// @Tags dentists
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} dentist.Dentist
// @Router /dentist/{id} [get]
func (d *DentistHandler) GetDentistByID(ctx *gin.Context) {

	id, err := utils.ValidateId(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	dentist, err := d.dentistGetter.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "dentist not found"})
		return
	}
	ctx.JSON(http.StatusOK, dentist)
}

// AddTurno godoc
// @Summary Add a dentist
// @Description Add a dentist
// @Tags dentists
// @Accept  json
// @Produce  json
// @Param dentist body dentist.Dentist true "Dentist"
// @Success 201 {object} dentist.Dentist
// @Router /dentist [post]
func (d *DentistHandler) Create(ctx *gin.Context) {

	newDentist, err := utils.ToDentist(ctx)
	if err != nil {
		return
	}

	response, err := d.dentistCreator.Create(newDentist)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "dentist not found"})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

// Update godoc
// @Summary Update a dentist
// @Description Update a dentist
// @Tags dentists
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param turno body dentist.Dentist true "Dentist"
// @Success 200 {object} dentist.Dentist
// @Router /dentist/{id} [put]
func (d *DentistHandler) Update(ctx *gin.Context) {
	id, err := utils.ValidateId(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	newDentist, err := utils.ToDentist(ctx)
	if err != nil {
		return
	}

	response, err := d.dentistUpdate.Update(id, newDentist)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

// SomeUpdate godoc
// @Summary Update a dentist with a patch
// @Description Update a dentist with a patch
// @Tags dentists
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param turno body dentist.Dentist true "Dentist"
// @Success 200 {object} dentist.Dentist
// @Router /turno/{id} [patch]
func (d *DentistHandler) Patch(ctx *gin.Context) {

	id, err := utils.ValidateId(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	newDentist, err := utils.ToDentistPatch(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response, err := d.dentistUpdate.Patch(id, newDentist)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

// Delete godoc
// @Summary Delete a dentist
// @Description Delete a dentist
// @Tags dentists
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {string}
// @Router /dentist/{id} [delete]
func (d *DentistHandler) Delete(ctx *gin.Context) {
	id, err := utils.ValidateId(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = d.dentistDelete.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
