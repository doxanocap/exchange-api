package configs

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var Public App

type App struct{}

var _m = map[string]string{}

func InitEnvironment() {
	_m = map[string]string{}

	if err := initEnvMap(); err != nil {
		logrus.Fatalf("Error in .env init: %e", err)
	}
}

func ENV(key string) string {
	osEnv := os.Getenv(key)
	if osEnv != "" {
		return osEnv
	}
	return _m[key]
}

func initEnvMap() error {
	err := *new(error)
	_m, err = godotenv.Read(".env")
	return err
}
