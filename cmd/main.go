package main

import (
	"dentistas/cmd/server/external/database"
	"dentistas/cmd/server/handler/dentist/handler"
	dentistRoutes "dentistas/cmd/server/routes/dentist"
	"dentistas/internal/dentist"
	patientRoutes "dentistas/cmd/server/routes/patient"
	"dentistas/internal/patient"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

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

	myDatabase := database.NewDatabase(mysqlDatabase)

	dentistService := dentist.NewService(myDatabase)

	dentistHandlerPo := handler.NewDentistHandler(dentistService, dentistService, dentistService, dentistService)

	dentistRoutes.SetupDentistRouter(router, dentistHandlerPo)

	patientDatabase := NewDatabase(mysqlDatabase)

	patientService := patient.NewService(patientDatabase)

	patientHandlerPo := handler.NewPatientHandler(patientService, patientService, patientService, patientService)

	patientRoutes.SetupPatientRouter(router, patientHandlerPo)

	err = router.Run()

	if err != nil {
		panic(err)
	}
}
