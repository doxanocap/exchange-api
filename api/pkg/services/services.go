package services

import (
	"api/pkg/dispatcher"
	"api/pkg/models"
)

type Auth interface {
	SignIn(body []byte) (models.AuthResponse, error)
	SignUp(body []byte) (models.AuthResponse, error)
	SignOut() models.ErrorResponse
	RefreshTokens() models.ErrorResponse
}

type Handler interface {
}

type Services struct {
	Auth
	Handler
}

func InitServices(dp *dispatcher.Dispatcher) *Services {
	return &Services{
		NewAuthService(dp),
		NewHandlerService(dp),
	}
}
