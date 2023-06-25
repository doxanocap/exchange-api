package router

import (
	"api/pkg/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (router *Router) SignIn(ctx *gin.Context) {
	data, err := ctx.GetRawData()
	if err != nil {
		router.newErrorResponse(ctx, models.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	res, err := router.services.Auth.SignIn(data)
	if err != nil {
		router.newErrorResponse(ctx, models.ErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	ctx.SetCookie(
		"refreshToken",
		res.RefreshToken,
		30*24*60*60*1000,
		"/",
		"localhost",
		false,
		true)

	ctx.JSON(200, res)
}

func (router *Router) SignUp(ctx *gin.Context) {
	data, err := ctx.GetRawData()
	if err != nil {
		router.newErrorResponse(ctx, models.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	res, err := router.services.Auth.SignUp(data)
	if err != nil {
		router.newErrorResponse(ctx, models.ErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	ctx.SetCookie(
		"refreshToken",
		res.RefreshToken,
		30*24*60*60*1000,
		"/",
		"localhost",
		false,
		true)

	ctx.JSON(res.Status, res)
}

func (router *Router) SignOut(ctx *gin.Context) {
	header := ctx.Request.Header
	for k, v := range header {
		log.Println(k, v)
	}
	if len(header["Cookie"]) == 0 {
		ctx.JSON(401, models.ErrorResponse{Status: http.StatusUnauthorized, Message: "User is unauthorized"})
		ctx.Abort()
		return
	}

	err1 := router.services.Auth.SignOut(header)
	if err1.IsError() {
		router.newErrorResponse(ctx, err1)
		return
	}

	ctx.SetCookie(
		"refreshToken",
		"",
		0,
		"/",
		"localhost",
		false,
		true)

	ctx.JSON(200, gin.H{
		"message": "success",
	})
}

func (router *Router) RefreshTokens(ctx *gin.Context) {
	header := ctx.Request.Header
	if len(header["Cookie"]) == 0 {
		ctx.JSON(401, models.ErrorResponse{Status: http.StatusUnauthorized, Message: "User is unauthorized"})
		ctx.Abort()
		return
	}

	res := router.services.Auth.RefreshTokens(header)
	ctx.JSON(res.Status, res)
}

func (router *Router) UserValidation(ctx *gin.Context) {
	header := ctx.Request.Header
	if len(header["Authorization"]) == 0 {
		ctx.JSON(401, models.ErrorResponse{Status: http.StatusUnauthorized, Message: "User is unauthorized"})
		ctx.Abort()
		return
	}

	res, err := router.services.UserValidation(header)
	if err.IsError() {
		ctx.JSON(401, models.ErrorResponse{Status: 401, Message: "User is unauthorized"})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, res)
}
