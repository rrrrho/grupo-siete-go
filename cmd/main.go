package main

import (
	"errors"
	"github.com/joho/godotenv"
	"grupo-siete-go/cmd/server/config"
	"grupo-siete-go/cmd/server/external/database"
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

}
