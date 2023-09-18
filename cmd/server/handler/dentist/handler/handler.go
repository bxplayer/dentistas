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
