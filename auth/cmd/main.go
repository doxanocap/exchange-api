package main

import (
	"auth/pkg/database"
	"auth/pkg/routes"
)

func main() {
	database.Connect()
	routes.SetupRoutes()
}
