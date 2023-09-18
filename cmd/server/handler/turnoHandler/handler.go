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

func (t *TurnoHandler) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	turnoRequest := turno2.Turno{}
	err = ctx.BindJSON(&turnoRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	turno, err := t.turnoUpdater.Update(id, turnoRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, turno)
}

func (t *TurnoHandler) DeleteByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	err = t.turnoDeleter.DeleteByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "turno deleted"})
}

func (t *TurnoHandler) GetByPacienteDNI(ctx *gin.Context) {
	dni := ctx.Param("dni")
	turnos, err := t.turnoGetter.GetByPacienteDNI(dni)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "turno not found"})
		return
	}
	ctx.JSON(http.StatusOK, turnos)
}

func (t *TurnoHandler) SomeUpdate(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	turnoRequest := turno2.Turno{}
	err = ctx.BindJSON(&turnoRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	turno, err := t.turnoUpdater.SomeUpdate(id, turnoRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, turno)
}
