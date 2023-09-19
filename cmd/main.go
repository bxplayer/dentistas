package main

import (
	"dentistas/cmd/server/external/database"
	"dentistas/cmd/server/external/database/turnoRepository"
	"dentistas/cmd/server/external/database/patientRepository"
	"dentistas/cmd/server/handler/patient"
	"dentistas/cmd/server/handler/turnoHandler"
	patientRoutes "dentistas/cmd/server/routes/patient"
	dentist2 "dentistas/cmd/server/handler/dentist"
	"dentistas/cmd/server/routes"
	"dentistas/docs"
	"dentistas/internal/dentist"
	"dentistas/internal/patient"
	"dentistas/internal/turno"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

// @title Certified Tech Developer - Dentistas
// @version 1.0
// @description This API Handle Products.

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.New()

	mysqlDatabase, err := database.NewMysqlDatabase(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	if err != nil {
		panic(err)
	}

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Configuracion de Dentista
	myDatabase := database.NewDatabase(mysqlDatabase)

	dentistService := dentist.NewService(myDatabase)

	dentistHandlerPo := dentist2.NewDentistHandler(dentistService, dentistService, dentistService, dentistService)

	routes.SetupDentistRouter(router, dentistHandlerPo)

	// Configuracion de Paciente

	patientDatabase := patientRepository.NewDatabase(mysqlDatabase)
	patientService := patient.NewService(patientDatabase)
	patientHandlerPo := patientHandler.NewPatientHandler(patientService, patientService, patientService, patientService)
	patientRoutes.SetupPatientRouter(router, patientHandlerPo)
	
	// Configuracion de Turno
	turnoDatabase := turnoRepository.NewDatabase(mysqlDatabase)
	turnoService := turno.NewService(turnoDatabase)
	turnoHandlerPro := turnoHandler.NewTurnoHandler(turnoService, turnoService, turnoService, turnoService)
	routes.SetupTurnRouter(router, turnoHandlerPro)

	err = router.Run()
	if err != nil {
		panic(err)
	}

	err = mysqlDatabase.Close()
	if err != nil {
		return
	}

}
