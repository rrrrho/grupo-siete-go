package main

import (
	"errors"
	"grupo-siete-go/cmd/server/config"
	"grupo-siete-go/cmd/server/external/database"
	"grupo-siete-go/cmd/server/handler"
	"grupo-siete-go/cmd/server/middlewares"
	"grupo-siete-go/docs"
	"grupo-siete-go/internal/odontologo"
	"grupo-siete-go/internal/paciente"
	"grupo-siete-go/internal/turno"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/swag/example/basic/docs"
)

// @title Clinica Odontologia Back End 3 - Grupo 7 - Certified Tech Developer - Digital House
// @version 1.0
// @description This API Handle Pacients, Dentist and Appointments.
// @termsOfService https://developers.ctd.com.ar/es_ar/terminos-y-condiciones
// @contact.name Rocio Belen Ghillino, Tomás Montivero, Agustin Damelio and Nicolás Gambino
// @contact.url https://github.com/rrrrho/grupo-siete-go
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	godotenv.Load()
	env := os.Getenv("env")

	if env == "" {
		panic(errors.New("env variable does not exits"))
	}

	// if env == "local" {
	// 	err := godotenv.Load()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	cfg, err := config.NewConfig(env)

	if err != nil {
		panic(err)
	}

	authMidd := middlewares.NewAuth(cfg.PublicConfig.PlubicKey, cfg.PrivateConfig.SecretKey)

	mySqlDatabase, err := database.NewMySQLDatabase(cfg.PublicConfig.MySQLHost,
		cfg.PublicConfig.MySQLPort, cfg.PublicConfig.MySQLUser, cfg.PrivateConfig.MySQLPassword,
		cfg.PublicConfig.MySQLDatabase)

	if err != nil {
		panic(err)
	}

	router := gin.Default()

	//docs endpoint
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"ok": "ok"})
	})

	// TURNOS
	turnoRepository := database.NewTurnoDatabase(mySqlDatabase)
	turnoService := turno.NewService(turnoRepository)
	turnoHandler := handler.NewTurnoHandler(turnoService, turnoService, turnoService, turnoService)

	turnoGroup := router.Group("/turnos")
	turnoGroup.GET("/:id", turnoHandler.GetByID)
	turnoGroup.GET("", turnoHandler.GetByDNI)
	turnoGroup.POST("", authMidd.AuthHeader, turnoHandler.Save)
	turnoGroup.PUT("", authMidd.AuthHeader, turnoHandler.Replace)
	turnoGroup.PATCH("/:id", authMidd.AuthHeader, turnoHandler.Update)
	turnoGroup.DELETE("/:id", authMidd.AuthHeader, turnoHandler.Delete)

	// ODONTOLOGOS
	odontologoRepository := database.NewOdontologoDatabase(mySqlDatabase)
	odontologoService := odontologo.NewService(odontologoRepository)
	odontologoHandler := handler.NewOdontologoHandler(odontologoService, odontologoService, odontologoService, odontologoService)

	odontologoGroup := router.Group("/odontologos")
	odontologoGroup.GET("/:id", odontologoHandler.GetByID)
	odontologoGroup.POST("", authMidd.AuthHeader, odontologoHandler.Save)
	odontologoGroup.PUT("", authMidd.AuthHeader, odontologoHandler.Replace)
	odontologoGroup.PATCH("/:id", authMidd.AuthHeader, odontologoHandler.Update)
	odontologoGroup.DELETE("/:id", authMidd.AuthHeader, odontologoHandler.Delete)

	// PACIENTES
	pacienteRepository := paciente.NewPacienteDatabase(mySqlDatabase)
	pacienteService := paciente.NewService(pacienteRepository)
	pacienteHandler := handler.NewPacienteHandler(*pacienteService)

	pacienteGroup := router.Group("/pacientes")
	pacienteGroup.GET("/:id", pacienteHandler.GetByID)
	pacienteGroup.POST("", pacienteHandler.Save)
	pacienteGroup.PATCH("/:id", pacienteHandler.Update)
	pacienteGroup.DELETE("/:id", pacienteHandler.Delete)

	err = router.Run()

	if err != nil {
		panic(err)
	}
}
