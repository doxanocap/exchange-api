package middlewares

import (
	"auth/pkg/models"
	"auth/pkg/services"
	"github.com/gin-gonic/gin"
)

func ValidateUserAuth(ctx *gin.Context) {
	auth := ctx.Request.Header["Authorization"]
	if len(auth) == 0 {
		ctx.JSON(401, models.Error{Status: 401, Message: "User is unauthorized"})
		ctx.Abort()
		return
	}

	accessToken := auth[0]
	user, err := services.ValidateAToken(accessToken)
	if err.Status != 200 {
		ctx.JSON(err.Status, err)
		ctx.Abort()
		return
	}
	ctx.Set("user", user)
	ctx.Next()
}