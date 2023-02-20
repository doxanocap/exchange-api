package app

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Public App

type App struct {
	ENV map[string]string
}

func InitEnvironment() {
	newApp := App{}
	if err := newApp.initConfigs(); err != nil {
		logrus.Fatalf("Error in init configs: %e", err)
	}

	if err := newApp.initEnvMap(); err != nil {
		logrus.Fatalf("Error in .env init: %e", err)
	}
	Public = newApp
}

func (a *App) initConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func (a *App) initEnvMap() error {
	err := *new(error)
	a.ENV, err = godotenv.Read(".env")
	return err
}
