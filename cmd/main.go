package main

import (
	"errors"
	"grupo-siete-go/cmd/server/config"
	"grupo-siete-go/cmd/server/external/database"
	"grupo-siete-go/cmd/server/handler"
	"grupo-siete-go/cmd/server/middlewares"
	"grupo-siete-go/internal/odontologo"
	"grupo-siete-go/internal/paciente"
	"grupo-siete-go/internal/turno"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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
	pacienteGroup.PUT("", pacienteHandler.Update)
	pacienteGroup.DELETE("/:id", pacienteHandler.Delete)

	err = router.Run()

	if err != nil {
		panic(err)
	}
}
