package router

import (
	"api/pkg/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	services *services.Services
}

func InitRouter(services *services.Services) *Router {
	return &Router{services: services}
}

func (router *Router) InitRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET", "PATCH", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Accept", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Authorization"},
		AllowCredentials: true,
	}))

	r.GET("/healthcheck", router.healthcheck)
	{
		api := r.Group("/api")
		{
			auth := api.Group("/auth")
			{
				auth.POST("/sign-in", router.SignIn)
				auth.POST("/sign-up", router.SignUp)
				auth.GET("/sign-out", router.SignOut)
				auth.GET("/refresh-token", router.RefreshTokens)
			}

			api.Use(router.UserValidation)
			// service called router
			handler := api.Group("/handler")
			{
				data := handler.Group("/data")
				{
					data.POST("/exchangers", router.ExchangersData)
					data.POST("/currencies", router.CurrenciesData)
				}
			}
		}
	}
	return r
}
