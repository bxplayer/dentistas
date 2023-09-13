package handler

import (
	"dentistas/internal/domain/dentista"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DentistGetter interface {
	GetByID(id int) (dentista.Dentista, error)
}

type DentistCreator interface {
	Update(id int, dentista dentista.Dentista) (dentista.Dentista, error)
	Add(dentista dentista.Dentista) (dentista.Dentista, error)
}

type DentistDelete interface {
	Delete(id int) error
}

type DentistHandler struct{
	dentistGetter DentistGetter
	dentistCreator DentistCreator
	dentistDelete DentistDelete
}

func newDentistHandler(getter DentistGetter, creator DentistCreator, delete DentistDelete) *DentistHandler {
	return &DentistHandler{
		dentistGetter:  getter,
		dentistCreator: creator,
		dentistDelete : delete,
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
		ctx.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	ctx.JSON(http.StatusOK, dentist)
}

func (d *DentistHandler) PutProduct(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	dentistRequest := dentista.Dentista{}
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
}