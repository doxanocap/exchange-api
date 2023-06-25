package controllers

import (
	"github.com/gin-gonic/gin"
)

func Healthcheck(ctx *gin.Context) {
	ctx.JSON(200, "AUTH service is alive")
}
