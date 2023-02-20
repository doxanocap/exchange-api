package controllers

import (
	"api/pkg/models"
	"api/pkg/services"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func ValidateUser(ctx *gin.Context) {
	body := services.SendGetRequest("http://localhost:8000/auth/user/validate")

	var response models.AuthResponseModel
	if err := json.Unmarshal(body, &response); err != nil {
		ctx.JSON(500, models.Error{Status: 500, Message: "Invalid response"})
		return
	}
	ctx.JSON(200, response)
}

func SingIn(ctx *gin.Context) {
	var data models.SignInInput

	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(500, models.Error{Status: 500, Message: "Invalid input"})
		return
	}
	postBody, _ := json.Marshal(data)

	body := services.SendPostRequest("http://localhost:8000/auth/sign-in", postBody)

	var response models.AuthResponseModel
	if err := json.Unmarshal(body, &response); err != nil {
		ctx.JSON(500, models.Error{Status: 500, Message: "Invalid response"})
		return
	}

	ctx.JSON(200, response)
}

func SignUp(ctx *gin.Context) {
	var data models.SignUpInput

	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(500, models.Error{Status: 500, Message: "Invalid input"})
		return
	}
	postBody, _ := json.Marshal(data)

	body := services.SendPostRequest("http://localhost:8000/auth/sign-up", postBody)

	var response models.AuthResponseModel
	if err := json.Unmarshal(body, &response); err != nil {
		ctx.JSON(500, models.Error{Status: 500, Message: "Invalid response"})
		return
	}

	ctx.JSON(200, response)
}

func SignOut(ctx *gin.Context) {
	body := services.SendGetRequest("http://localhost:8000/auth/sign-out")

	var response models.AuthResponseModel
	if err := json.Unmarshal(body, &response); err != nil {
		ctx.JSON(500, models.Error{Status: 500, Message: "Invalid response"})
		return
	}
	ctx.JSON(200, response)
}

func Refresh(ctx *gin.Context) {
	body := services.SendGetRequest("http://localhost:8000/auth/sign-out")

	var response models.AuthResponseModel
	if err := json.Unmarshal(body, &response); err != nil {
		ctx.JSON(500, models.Error{Status: 500, Message: "Invalid response"})
		return
	}
	ctx.JSON(200, response)
}
