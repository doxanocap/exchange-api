package service

import (
	"auth/pkg/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(email, username, password string) models.AuthResponseModel {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	if _, err := models.FindUserByEmail(email); err.Status == 200 {
		return models.AuthResponseModel{Status: 400, Message: "user already exists"}
	}

	user, err := models.InsertUserToDB(email, username, string(hashed))
	if err.Status >= 300 {
		return models.AuthResponseModel{Status: err.Status, Message: err.Message}
	}

	tokens := GenerateTokens(user)
	_, err = models.InsertTokenToDB(user.Id, tokens[1])
	if err.Status >= 300 {
		return models.AuthResponseModel{Status: err.Status, Message: "Invalid request to DB"}
	}

	return models.AuthResponseModel{
		Status:       200,
		AccessToken:  tokens[0],
		RefreshToken: tokens[1],
		User:         user,
		Message:      ""}
}

func LoginUser(email, password string) models.AuthResponseModel {
	response, err := models.FindUserByEmail(email)
	if err.Status != 200 {
		return models.AuthResponseModel{Status: err.Status, Message: err.Message}
	}

	if err := bcrypt.CompareHashAndPassword(response.Password, []byte(password)); err != nil {
		log.Println("services -> user -> LoginUser -> ", err)
		return models.AuthResponseModel{Status: 400, Message: "Incorrect password"}
	}

	tokens := GenerateTokens(response)
	if err := SaveToken(response.Id, tokens[1]); err.Status != 200 {
		return models.AuthResponseModel{Status: err.Status, Message: err.Message}
	}
	return models.AuthResponseModel{
		Status:       200,
		AccessToken:  tokens[0],
		RefreshToken: tokens[1],
		User:         response,
		Message:      "",
	}
}

func RefreshUser(refreshToken string) models.AuthResponseModel {
	user, err := ValidateRToken(refreshToken)

	if err.Status != 200 {
		return models.AuthResponseModel{Status: err.Status, Message: err.Message}
	}

	tokens := GenerateTokens(user)
	err = SaveToken(user.Id, tokens[1])
	if err.Status != 200 {
		return models.AuthResponseModel{Status: err.Status, Message: "Undable to refresh"}
	}
	return models.AuthResponseModel{
		Status:       200,
		AccessToken:  tokens[0],
		RefreshToken: tokens[1],
		User:         user,
		Message:      ""}
}
