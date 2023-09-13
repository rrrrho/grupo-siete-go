package config

import (
	"errors"
	"os"
)

type Config struct {
	PublicConfig  PublicConfig
	PrivateConfig PrivateConfig
}

type PublicConfig struct {
	MySQLUser     string
	MySQLHost     string
	MySQLPort     string
	MySQLDatabase string
}

type PrivateConfig struct {
	MySQLPassword string
	// SecretKey string
}

var (
	ErrEnvNotExits       = errors.New("env not exits")
	ErrMysqlPassNotExits = errors.New("mysql password does not exits in env")
	// ErrSecretKeyNotExits = errors.New("secret key does not exits in env")
)

var (
	envs = map[string]PublicConfig{
		"local": {
			MySQLUser:     "local-clinica",
			MySQLHost:     "localhost",
			MySQLPort:     "3606",
			MySQLDatabase: "local-clinica-database",
		},
		"dev": {
			MySQLUser:     "dev-clinica",
			MySQLHost:     "localhost",
			MySQLPort:     "3607",
			MySQLDatabase: "dev-clinica-database",
		},
		"prod": {
			MySQLUser:     "prod-clinica",
			MySQLHost:     "localhost",
			MySQLPort:     "3608",
			MySQLDatabase: "prod-clinica-database",
		},
	}
)

func NewConfig(env string) (Config, error) {
	publicConfig, exists := envs[env]

	if !exists {
		return Config{}, ErrEnvNotExits
	}

	// secretKey := os.Getenv("SECRET_KEY")
	/*
		if secretKey == "" {
			return Config{}, ErrSecretKeyNotExits
		}
	*/
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	if mysqlPassword == "" {
		return Config{}, ErrMysqlPassNotExits
	}

	return Config{
		PublicConfig: publicConfig,
		PrivateConfig: PrivateConfig{
			// SecretKey: secretKey
			MySQLPassword: mysqlPassword,
		},
	}, nil
}
