package controllers

import (
	"api/pkg/models"
	"api/pkg/services"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func GetAllData(ctx *gin.Context) {
	body := services.SendGetRequest("http://localhost:8002/data/all")

	var response models.AuthResponseModel
	if err := json.Unmarshal(body, &response); err != nil {
		ctx.JSON(500, models.Error{Status: 500, Message: "Invalid response"})
		return
	}
	ctx.JSON(200, response)
}
