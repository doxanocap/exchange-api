package controllers

import (
	"auth/pkg/models"
	"auth/pkg/services"
	"github.com/gin-gonic/gin"
	"log"
)

func SignUp(ctx *gin.Context) {
	var data map[string]string

	if err := ctx.BindJSON(&data); err != nil {
		panic(err)
	}

	response := services.CreateUser(
		data["email"],
		data["username"],
		data["password"])

	if response.Status != 200 {
		ctx.JSON(response.Status, response)
		return
	}

	ctx.SetCookie(
		"refreshToken",
		response.RefreshToken,
		30*24*60*60*1000,
		"/",
		"localhost",
		false,
		true)

	ctx.JSON(200, response)
}

func SignIn(ctx *gin.Context) {
	var data map[string]string
	if err := ctx.BindJSON(&data); err != nil {
		log.Println("controllers -> user -> signIn -> 1")
	}

	response := services.LoginUser(data["email"], data["password"])
	if response.Status != 200 {
		log.Println("controllers -> user -> signIn -> 2", response.Message)
		ctx.JSON(response.Status, models.Error{Status: response.Status, Message: response.Message})
		return
	}

	ctx.SetCookie(
		"refreshToken",
		response.RefreshToken,
		30*24*60*60*1000,
		"/",
		"localhost",
		false,
		true)

	ctx.JSON(response.Status, response)
}

func SignOut(ctx *gin.Context) {
	token, err := ctx.Cookie("refreshToken")
	if err != nil {
		ctx.JSON(500, models.Error{Status: 500, Message: "Something went wrong: Token <-> Client"})
		return
	}
	if token == "" {
		ctx.JSON(400, models.Error{Status: 400, Message: "You are not authorized to sign-out"})
		return
	}
	if err := services.DeleteToken(token); err.Status != 200 {
		ctx.JSON(err.Status, err)
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

	ctx.JSON(200, models.Error{Status: 200, Message: "Successfully signed out"})
}

func RefreshUser(ctx *gin.Context) {
	token, err := ctx.Cookie("refreshToken")
	if err != nil || token == "" {
		ctx.JSON(500, models.Error{Status: 500, Message: "Something went wrong: Token <-> Client"})
		return
	}
	response := services.RefreshUser(token)
	ctx.JSON(response.Status, response)
}

func AccountInformation(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(500, gin.H{"error": "unhandled error"})
	}

	ctx.JSON(200, user)
}
