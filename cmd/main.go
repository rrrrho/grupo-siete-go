package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"grupo-siete-go/cmd/server/config"
	"grupo-siete-go/cmd/server/external/database"
	"grupo-siete-go/cmd/server/handler"
	"grupo-siete-go/internal/turno"
	"os"
)

func main() {
	godotenv.Load()
	env := os.Getenv("env")

	if env == "" {
		panic(errors.New("env variable does not exits"))
	}

	cfg, err := config.NewConfig(env)

	if err != nil {
		panic(err)
	}

	mySqlDatabase, err := database.NewMySQLDatabase(cfg.PublicConfig.MySQLHost,
		cfg.PublicConfig.MySQLPort, cfg.PublicConfig.MySQLUser, cfg.PrivateConfig.MySQLPassword,
		cfg.PublicConfig.MySQLDatabase)

	if err != nil {
		panic(err)
	}

	router := gin.Default()

	turnoRepository := database.NewTurnoDatabase(mySqlDatabase)
	turnoService := turno.NewService(turnoRepository)
	turnoHandler := handler.NewTurnoHandler(turnoService, turnoService, turnoService, turnoService)

	turnoGroup := router.Group("/turnos")
	turnoGroup.GET("/:id", turnoHandler.GetByID)
	turnoGroup.GET("", turnoHandler.GetByDNI)
	turnoGroup.POST("", turnoHandler.Save)
	turnoGroup.PUT("", turnoHandler.Replace)
	turnoGroup.PATCH("/:id", turnoHandler.Update)
	turnoGroup.DELETE("/:id", turnoHandler.Delete)

	err = router.Run()

	if err != nil {
		panic(err)
	}
}
