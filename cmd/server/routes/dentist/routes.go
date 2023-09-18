package dentist

import (
	eeee "dentistas/cmd/server/handler/dentist/handler"
	"github.com/gin-gonic/gin"
)

func SetupDentistRouter(router *gin.Engine, dentistHandler *eeee.DentistHandler) {

	dentistGroup := router.Group("/dentist")
	{
		dentistGroup.GET("/:id", dentistHandler.GetDentistByID)
		dentistGroup.PUT("/:id", dentistHandler.Update)
		dentistGroup.PATCH("/:id", dentistHandler.Patch)
		dentistGroup.DELETE("/:id", dentistHandler.Delete)
		dentistGroup.POST("/", dentistHandler.Create)
	}

}
