package service

import (
	"auth/pkg/models"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

func GenerateTokens(user models.User) []string {
	res, err := json.Marshal(user)
	if err != nil {
		log.Println("json marshal", err)
	}

	accessClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		Issuer:    string(res),
	})
	refreshClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 20)),
		Issuer:    string(res),
	})
	accessToken, _ := accessClaim.SignedString([]byte("JWT_SECRET_KEY"))
	refreshToken, _ := refreshClaim.SignedString([]byte("JWT_SECRET_KEY"))

	return []string{accessToken, refreshToken}
}

func ValidateRToken(refreshToken string) (models.User, models.Error) {
	token, err := models.FindToken(refreshToken)
	if err.Status != 200 {
		return models.User{}, err
	}

	res, err1 := jwt.ParseWithClaims(token.RefreshToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("JWT_SECRET_KEY"), nil
	})
	if err1 != nil {
		return models.User{}, models.Error{Status: 500, Message: "Error in handling claims"}
	}

	claims, ok := res.Claims.(*jwt.RegisteredClaims)
	user := models.User{}
	_ = json.Unmarshal([]byte(claims.Issuer), &user)
	if ok {
		return user, models.Error{Status: 200, Message: ""}
	}
	return models.User{}, models.Error{Status: 500, Message: "Invalid token"}
}

func ValidateAToken(accessToken string) (models.User, models.Error) {
	token, err := jwt.ParseWithClaims(accessToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("JWT_SECRET_KEY"), nil
	})

	if err != nil {
		return models.User{}, models.Error{Status: 500, Message: "Error in handling claims"}
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	user := models.User{}
	_ = json.Unmarshal([]byte(claims.Issuer), &user)
	if ok {
		return user, models.Error{Status: 200, Message: ""}
	}
	return models.User{}, models.Error{Status: 500, Message: "Invalid token"}
}

func DeleteToken(refreshToken string) models.Error {
	return models.DeleteToken(refreshToken)
}

func SaveToken(id int, refreshToken string) models.Error {
	_, err := models.FindTokenById(id)
	if err.Status == 200 {
		return models.UpdateToken(id, refreshToken)
	}
	if err.Status == 401 {
		if _, err := models.InsertTokenToDB(id, refreshToken); err.Status != 200 {
			return err
		}
		return models.Error{Status: 200, Message: ""}
	}
	return err
}
