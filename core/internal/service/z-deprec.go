package service

import (
	"api/pkg/dispatcher"
	"api/pkg/repository"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type AuthService struct {
	dp *dispatcher.NewDispatcher
}

func NewAuthService(dp *dispatcher.NewDispatcher) *AuthService {
	return &AuthService{dp}
}

func (service *AuthService) SignIn(body []byte) (repository.AuthResponse, error) {
	res, err := service.dp.Post("/auth/sign-in", body)
	if err.Message != "" {
		return repository.AuthResponse{}, errors.New("response from AUTH/sign-in: " + err.Message)
	}
	var authRes repository.AuthResponse
	if err := json.Unmarshal(res, &authRes); err != nil {
		return repository.AuthResponse{}, errors.Wrap(err, "trying to unmarshal AUTH/sign-in")
	}
	return authRes, nil
}

func (service *AuthService) SignUp(body []byte) (repository.AuthResponse, error) {
	res, err := service.dp.Post("/auth/sign-up", body)
	if err.Message != "" {
		return repository.AuthResponse{}, errors.New("response from AUTH/sign-up response: " + err.Message)
	}

	var authRes repository.AuthResponse
	if err := json.Unmarshal(res, &authRes); err != nil {
		return repository.AuthResponse{}, errors.Wrap(err, "trying to unmarshal AUTH/sign-up response")
	}

	return authRes, nil
}

func (service *AuthService) SignOut(header http.Header) repository.ErrorResponse {
	res, err := service.dp.GetRequest(
		"/auth/sign-out",
		header)

	if err.Message != "" {
		return err
	}

	var errRes repository.ErrorResponse
	if err := json.Unmarshal(res, &errRes); err != nil {
		return repository.ErrorResponse{
			Status:  http.StatusBadGateway,
			Message: errors.Wrap(err, "trying to unmarshal AUTH/sign-out response").Error()}
	}

	return errRes
}

func (service *AuthService) RefreshTokens(header http.Header) repository.ErrorResponse {
	res, err := service.dp.GetRequest(
		"/auth/refresh-token",
		header)

	if err.Message != "" {
		return err
	}

	var errRes repository.ErrorResponse
	if err := json.Unmarshal(res, &errRes); err != nil {
		return repository.ErrorResponse{
			Status:  http.StatusBadGateway,
			Message: errors.Wrap(err, "trying to unmarshal AUTH/refresh-token response").Error()}
	}

	return errRes
}

func (service *AuthService) UserValidation(header http.Header) (repository.AuthUser, repository.ErrorResponse) {
	res, err := service.dp.PostRequest(
		"/auth/user/validate",
		[]byte{},
		header)

	if err.IsError() {
		return repository.AuthUser{}, err
	}

	var authResponse repository.AuthUser
	var errResponse repository.ErrorResponse
	err1 := json.Unmarshal(res, &authResponse)
	err2 := json.Unmarshal(res, &errResponse)

	if err1 != nil && err2 == nil {
		return repository.AuthUser{}, errResponse
	} else if err1 == nil && err2 != nil {
		return authResponse, repository.ErrorResponse{}
	} else {
		return repository.AuthUser{}, repository.ErrorResponse{
			Status:  http.StatusBadGateway,
			Message: "trying to unmarshal AUTH/validate response: 1-> " + err1.Error() + "| 2 -> " + err2.Error(),
		}
	}
}
