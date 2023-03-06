package handlers

import (
	"api/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *Handler) SignIn(ctx *gin.Context) {
	var data []byte
	if _, err := ctx.Request.Body.Read(data); err != nil {
		handler.newErrorResponse(ctx, models.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	res, err := handler.services.Auth.SignIn(data)
	if err != nil {
		handler.newErrorResponse(ctx, models.ErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
	}
	ctx.JSON(200, res)
}

func (handler *Handler) SignUp(ctx *gin.Context) {
	var data []byte
	if _, err := ctx.Request.Body.Read(data); err != nil {
		handler.newErrorResponse(ctx, models.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	res, err := handler.services.Auth.SignUp(data)
	if err != nil {
		handler.newErrorResponse(ctx, models.ErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
	}
	ctx.JSON(200, res)
}

func (handler *Handler) SignOut(ctx *gin.Context) {
	var data []byte
	if _, err := ctx.Request.Body.Read(data); err != nil {
		handler.newErrorResponse(ctx, models.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	res := handler.services.Auth.SignOut()
	ctx.JSON(200, res)
}

func (handler *Handler) RefreshTokens(ctx *gin.Context) {
	var data []byte
	if _, err := ctx.Request.Body.Read(data); err != nil {
		handler.newErrorResponse(ctx, models.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	res := handler.services.Auth.RefreshTokens()
	ctx.JSON(200, res)
}
