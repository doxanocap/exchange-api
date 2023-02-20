package router

import (
	"api/pkg/controllers"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.Default()

	router.Use(
		cors.New(
			cors.Config{
				AllowOrigins:     []string{"http://localhost:3000"},
				AllowMethods:     []string{"POST", "GET", "PATCH", "PUT", "DELETE"},
				AllowHeaders:     []string{"Content-Type", "Accept", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
				ExposeHeaders:    []string{"Authorization"},
				AllowCredentials: true,
			}))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", controllers.SingIn)
		auth.POST("/sign-in", controllers.SignUp)
		auth.GET("/sign-out", controllers.SignOut)
		auth.GET("/refresh", controllers.Refresh)
	}

	api := router.Group("/api")
	{
		api.Use(controllers.ValidateUser)
		data := api.Group("/data")
		{
			data.GET("/all", controllers.GetAllData)
		}
	}

	router.Run(":" + port)
}
