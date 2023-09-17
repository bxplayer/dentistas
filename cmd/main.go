package main

import (
	"dentistas/cmd/server/external/database"
	"dentistas/cmd/server/external/database/turnoRepository"
	"dentistas/cmd/server/handler/dentist/handler"
	"dentistas/cmd/server/handler/turnoHandler"
	"dentistas/internal/dentist"
	"dentistas/internal/turno"
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

	// Configuracion de DentistaId
	myDatabase := database.NewDatabase(mysqlDatabase)

	dentistService := dentist.NewService(myDatabase)

	dentistHandlerPo := handler.NewDentistHandler(dentistService, dentistService, dentistService, dentistService)

	dentistGroup := router.Group("/dentist")

	dentistGroup.GET("/:id", dentistHandlerPo.GetDentistByID)

	// Configuracion de Turno
	turnoDatabase := turnoRepository.NewDatabase(mysqlDatabase)
	turnoService := turno.NewService(turnoDatabase)
	turnoHandlerPro := turnoHandler.NewTurnoHandler(turnoService, turnoService, turnoService, turnoService)
	turnoGroup := router.Group("/turno")
	turnoGroup.GET("/:id", turnoHandlerPro.GetTurnoByID)
	turnoGroup.POST("/", turnoHandlerPro.AddTurno)
	turnoGroup.PUT("/:id", turnoHandlerPro.Update)

	err = router.Run()

	if err != nil {
		panic(err)
	}
}
