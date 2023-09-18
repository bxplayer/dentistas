package routes

import (
	handlerDentist "dentistas/cmd/server/handler/dentist/handler"
	handlerTurn "dentistas/cmd/server/handler/turnoHandler"
	"dentistas/cmd/server/middleware"
	"github.com/gin-gonic/gin"
)

func SetupDentistRouter(router *gin.Engine, dentistHandler *handlerDentist.DentistHandler) {

	dentistGroup := router.Group("/dentist")
	dentistGroup.GET("/:id", dentistHandler.GetDentistByID)

	dentistGroup.Use(middleware.Authenticate())
	{
		dentistGroup.PUT("/:id", dentistHandler.Update)
		dentistGroup.PATCH("/:id", dentistHandler.Patch)
		dentistGroup.DELETE("/:id", dentistHandler.Delete)
		dentistGroup.POST("/", dentistHandler.Create)
	}

}

func SetupTurnRouter(router *gin.Engine, turnHandler *handlerTurn.TurnoHandler) {

	turnoGroup := router.Group("/turno")
	turnoGroup.GET("/", turnHandler.GetByPacienteDNI)
	turnoGroup.GET("/:id", turnHandler.GetTurnoByID)

	turnoGroup.Use(middleware.Authenticate())
	{
		turnoGroup.POST("/", turnHandler.AddTurno)
		turnoGroup.PUT("/:id", turnHandler.Update)
		turnoGroup.DELETE("/:id", turnHandler.DeleteByID)
		turnoGroup.PATCH("/:id", turnHandler.SomeUpdate)
	}

}
