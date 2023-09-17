package turnoHandler

import (
	turno2 "dentistas/internal/turno"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
)

type TurnoHandler struct {
	turnoGetter  TurnoGetter
	turnoUpdater TurnoUpdater
	turnoCreator TurnoCreator
	turnoDeleter TurnoDeleter
}

func NewTurnoHandler(
	turnoGetter TurnoGetter,
	turnoUpdater TurnoUpdater,
	turnoCreator TurnoCreator,
	turnoDeleter TurnoDeleter) *TurnoHandler {
	return &TurnoHandler{
		turnoGetter:  turnoGetter,
		turnoUpdater: turnoUpdater,
		turnoCreator: turnoCreator,
		turnoDeleter: turnoDeleter,
	}
}

func (t *TurnoHandler) GetTurnoByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	turno, err := t.turnoGetter.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "turno not found"})
		return
	}
	ctx.JSON(http.StatusOK, turno)
}

func (t *TurnoHandler) AddTurno(ctx *gin.Context) {
	turnoRequest := turno2.Turno{}
	if err := ctx.ShouldBindBodyWith(&turnoRequest, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	turno, err := t.turnoCreator.AddTurno(turnoRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, turno)
}
