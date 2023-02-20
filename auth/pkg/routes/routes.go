package routes

import (
	"auth/pkg/controllers"
	"auth/pkg/middlewares"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET", "PATCH", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Accept", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Authorization"},
		AllowCredentials: true,
	}))

	auth := r.Group("/auth")
	auth.POST("/sign-up", controllers.SignUp)
	auth.POST("/sign-in", controllers.SignIn)
	auth.GET("/refresh", controllers.RefreshUser)
	auth.GET("/sign-out", controllers.SignOut)

	user := r.Group("user")
	user.Use(middlewares.ValidateUserAuth)
	user.GET("/validate", controllers.AccountInformation)
	r.Run(":" + port)
}
