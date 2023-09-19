package patient

import (
	handler "dentistas/cmd/server/handler/patient"
	"github.com/gin-gonic/gin"
)

func SetupPatientRouter(router *gin.Engine, patientHandler *handler.PatientHandler) {

	patientGroup := router.Group("/patient")
	{
		patientGroup.GET("/:id", patientHandler.GetPatientByID)
		patientGroup.PUT("/:id", patientHandler.Update)
		patientGroup.PATCH("/:id", patientHandler.Patch)
		patientGroup.DELETE("/:id", patientHandler.Delete)
		patientGroup.POST("/", patientHandler.Create)
	}

}
