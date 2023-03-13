package services

import (
	"api/pkg/dispatcher"
	"api/pkg/models"
	"net/http"
)

type Auth interface {
	SignOut(header http.Header) models.ErrorResponse
	SignIn(body []byte) (models.AuthResponse, error)
	SignUp(body []byte) (models.AuthResponse, error)
	RefreshTokens(header http.Header) models.ErrorResponse
	UserValidation(header http.Header) (models.AuthUser, models.ErrorResponse)
}

type Handler interface {
	ExchangersData(body []byte) ([]models.ExchangersResponse, models.ErrorResponse)
	CurrenciesData(body []byte) ([]models.CurrenciesResponse, models.ErrorResponse)
}

type Services struct {
	Auth
	Handler
}

func InitServices(dp *dispatcher.Dispatcher) *Services {
	return &Services{
		NewAuthService(dp.AuthDispatcher),
		NewHandlerService(dp.HandlerDispatcher),
	}
}
