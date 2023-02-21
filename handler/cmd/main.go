package main

import (
	"fmt"
	"handler/pkg/app"
	"handler/pkg/handlers"
	"handler/pkg/repository"
	"handler/pkg/repository/postgres"
	"handler/pkg/services"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

	if err := postgres.Migrations(db); err != nil {
		fmt.Printf("Migration went successfull:\nMigration file -> handler%s\n", err.Error())
		return
	}

	repos := repository.InitRepository(db)
	service := services.InitServices(repos)
	handler := handlers.InitHandler(service)
	log.Println("Success")

	if err := server.Run(app.Public.ENV["PORT"], handler.InitRoutes()); err != nil {
		log.Fatalf("Error whilte running http server: %s", err.Error())
	}
}
