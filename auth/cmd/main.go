package main

import (
	"auth/pkg/configs"
	"auth/pkg/database"
	"auth/pkg/routes"
	"database/sql"
)

var DB *sql.DB

func main() {
	configs.InitEnvironment()

	database.Connect(database.Config{
		Host:     configs.ENV("POSTGRES_HOST"),
		Port:     configs.ENV("POSTGRES_PORT"),
		Username: configs.ENV("POSTGRES_USER"),
		DBName:   configs.ENV("POSTGRES_DB"),
		SSLMode:  configs.ENV("POSTGRES_SSL"),
		Password: configs.ENV("POSTGRES_PASSWORD")})

	routes.SetupRoutes()
}
