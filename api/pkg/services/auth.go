package services

import (
	"api/pkg/configs"
	"api/pkg/dispatcher"
	"api/pkg/models"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

var (
	AUTH_URL = configs.ENV("DOMAINS_AUTH")
)

type AuthService struct {
	dp *dispatcher.Dispatcher
}

func NewAuthService(dp *dispatcher.Dispatcher) *AuthService {
	return &AuthService{dp}
}

func (service *AuthService) SignIn(body []byte) (models.AuthResponse, error) {
	res, err := service.dp.Post(AUTH_URL+"/sign-in", body)
	if err != nil {
		return models.AuthResponse{}, errors.Wrap(err, "response from AUTH/sign-in")
	}
	var authRes models.AuthResponse
	if err := json.Unmarshal(res, &authRes); err != nil {
		return models.AuthResponse{}, errors.Wrap(err, "trying to unmarshal AUTH/sign-in")
	}
	return authRes, nil
}

func (service *AuthService) SignUp(body []byte) (models.AuthResponse, error) {
	res, err := service.dp.Post(AUTH_URL+"/sign-up", body)
	if err != nil {
		return models.AuthResponse{}, errors.Wrap(err, "response from AUTH/sign-up response")
	}

	var authRes models.AuthResponse
	if err := json.Unmarshal(res, &authRes); err != nil {
		return models.AuthResponse{}, errors.Wrap(err, "trying to unmarshal AUTH/sign-up response")
	}

	return authRes, nil
}

func (service *AuthService) SignOut() models.ErrorResponse {
	res, err := service.dp.Get(AUTH_URL + "/sign-up")
	if err != nil {
		return models.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: errors.Wrap(err, "response from AUTH/sign-up").Error()}
	}

	var errRes models.ErrorResponse
	if err := json.Unmarshal(res, &errRes); err != nil {
		return models.ErrorResponse{
			Status:  http.StatusBadGateway,
			Message: errors.Wrap(err, "trying to unmarshal AUTH/sign-out response").Error()}
	}

	return errRes
}

func (service *AuthService) RefreshTokens() models.ErrorResponse {
	res, err := service.dp.Get(AUTH_URL + "/sign-up")
	if err != nil {
		return models.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: errors.Wrap(err, "response from AUTH/sign-up").Error()}
	}

	var errRes models.ErrorResponse
	if err := json.Unmarshal(res, &errRes); err != nil {
		return models.ErrorResponse{
			Status:  http.StatusBadGateway,
			Message: errors.Wrap(err, "trying to unmarshal AUTH/sign-out response").Error()}
	}

	return errRes
}
