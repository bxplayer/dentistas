package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	dentist, err := d.dentistGetter.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "dentist not found"})
		return
	}
	ctx.JSON(http.StatusOK, dentist)
}

/*func (d *Handler) PutDentist(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	dentistRequest := dentist.Dentist{}
	err = ctx.BindJSON(&dentistRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dentist, err := d.dentistCreator.Update(id, dentistRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}
	ctx.JSON(200, dentist)
}*/
