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

// GetTurnoByID godoc
// @Summary Get turno by ID
// @Description Get turno by ID
// @Tags turnos
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} turno.Turno
// @Router /turno/{id} [get]
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

// AddTurno godoc
// @Summary Add a turno
// @Description Add a turno
// @Tags turnos
// @Accept  json
// @Produce  json
// @Param turno body turno.Turno true "Turno"
// @Param TOKEN header string false "Token"
// @Success 201 {object} turno.Turno
// @Router /turno [post]
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

// Update godoc
// @Summary Update a turno
// @Description Update a turno
// @Tags turnos
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param turno body turno.Turno true "Turno"
// @Param TOKEN header string false "Token"
// @Success 200 {object} turno.Turno
// @Router /turno/{id} [put]
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

// DeleteByID godoc
// @Summary Delete a turno
// @Description Delete a turno
// @Tags turnos
// @Produce  json
// @Param id path string true "ID"
// @Param TOKEN header string false "Token"
// @Success 200 {string} string "turno deleted"
// @Router /turno/{id} [delete]
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

// GetByPacienteDNI GetByDentistaID godoc
// @Summary Get turnos by Paciente ID
// @Description Get turnos by Paciente ID
// @Tags turnos
// @Produce json
// @Param dni query string true "DNI"
// @Success 200 {object} []turno.Turno
// @Router /turno [get]
func (t *TurnoHandler) GetByPacienteDNI(ctx *gin.Context) {
	dni := ctx.Query("dni")
	turnos, err := t.turnoGetter.GetByPacienteDNI(dni)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "turno not found"})
		return
	}
	ctx.JSON(http.StatusOK, turnos)
}

// SomeUpdate godoc
// @Summary Update a turno with a patch
// @Description Update a turno with a patch
// @Tags turnos
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param TOKEN header string false "Token"
// @Param turno body turno.Turno true "Turno"
// @Success 200 {object} turno.Turno
// @Router /turno/{id} [patch]
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
