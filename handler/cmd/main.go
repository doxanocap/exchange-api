package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"handler/pkg/app"
	"handler/pkg/handlers"
	"handler/pkg/repository"
	"handler/pkg/repository/postgres"
	"handler/pkg/services"
)

func main() {
	server := app.Server{}
	app.InitEnvironment()

	db := postgres.InitDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: app.Public.ENV["PSQL_PASSWORD"]})

	repository := repository.InitRepository(db)
	service := services.InitServices(repository)
	handler := handlers.InitHandler(service)

	if err := server.Run(app.Public.ENV["PORT"], handler.InitRoutes()); err != nil {
		logrus.Fatalf("Error whilte running http server: %s", err.Error())
	}
}
