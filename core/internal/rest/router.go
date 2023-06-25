package rest

import (
	"core/internal/manager/interfaces"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Service *interfaces.Service
}

func InitRouter(srvs *interfaces.Service) *Router {
	return &Router{Service: srvs}
}

func (r *Router) AddRoutes(router *gin.Engine) {
	router.GET("/healthcheck", r.Service.Handler.healthcheck)
	{
		api := router.Group("/api")
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
			handler := api.Group("/rest")
			{
				data := handler.Group("/data")
				{
					data.POST("/exchangers", router.ExchangersData)
					data.POST("/currencies", router.CurrenciesData)
				}
			}
		}
	}

	webchat := r.Group("/web-chat")
	{
		webchat.GET("/pool", handler.webSocketConn)

		user := webchat.Group("/user")
		{
			user.GET("/online", func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{"yeah": "buddy"})
			})
		}
		{
			webchat.POST("/active-chats")
			webchat.GET("/messages/:chatUID")
		}
	}
	return r
}
