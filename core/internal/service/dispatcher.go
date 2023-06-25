package service 

import (
	"api/pkg/configs"
	"api/pkg/repository"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)
type Dispatch interface {
	Get(endpoint string) ([]byte, repository.ErrorResponse)
	Post(endpoint string, postBody []byte) ([]byte, repository.ErrorResponse)
	PostRequest(endpoint string, body []byte, header http.Header) ([]byte, repository.ErrorResponse)
	GetRequest(endpoint string, header http.Header) ([]byte, repository.ErrorResponse)
}

type NewDispatcher struct {
	client http.Client
	url    string
}

type Dispatcher struct {
	AuthDispatcher    *NewDispatcher
	HandlerDispatcher *NewDispatcher
}

func InitDispatcher() *Dispatcher {
	var AUTH_URL = configs.ENV("DOMAINS_AUTH")
	var HANDLER_URL = configs.ENV("DOMAINS_HANDLER")

	return &Dispatcher{
		NewAuthDispatcher(AUTH_URL),
		NewHandlerDispatcher(HANDLER_URL),
	}
}

func (dispatcher *NewDispatcher) Get(endpoint string) ([]byte, repository.ErrorResponse) {
	resp, err := http.Get(dispatcher.url + endpoint)
	code := resp.StatusCode
	
	if err != nil {
		return nil, repository.ErrorResponse{
			Status:  code,
			Message: "sending get request: " + err.Error()}
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, repository.ErrorResponse{
			Status:  code,
			Message: "sending get request: " + err.Error()}
	}

	return body, repository.ErrorResponse{Status: http.StatusOK, Message: ""}
}

// Post Method for sending POST request
func (dispatcher *NewDispatcher) Post(endpoint string, postBody []byte) ([]byte, repository.ErrorResponse) {
	resp, err := http.Post(
		dispatcher.url+endpoint,
		"application/json",
		bytes.NewBufferString(string(postBody)))
	code := resp.StatusCode

	if err != nil {
		return nil, repository.ErrorResponse{
			Status:  code,
			Message: "sending get request: " + err.Error()}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, repository.ErrorResponse{
			Status:  code,
			Message: "sending get request: " + err.Error()}
	}

	return body, repository.ErrorResponse{Status: http.StatusOK, Message: ""}
}

func (dispatcher *NewDispatcher) PostRequest(endpoint string, body []byte, header http.Header) ([]byte, repository.ErrorResponse) {
	reader := bytes.NewReader(body)
	req, err := http.NewRequest(http.MethodPost, dispatcher.url+endpoint, reader)
	if err != nil {
		return nil, repository.ErrorResponse{
			Status:  500,
			Message: fmt.Sprintf("invalid request to the servers: %s%s -> %s", dispatcher.url, endpoint, err.Error()),
		}
	}
	req.Header = header
	res, err := dispatcher.client.Do(req)
	defer res.Body.Close()
	code := res.StatusCode

	if err != nil {
		return nil, repository.ErrorResponse{
			Status:  code,
			Message: "sending get request: " + err.Error()}
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, repository.ErrorResponse{
			Status:  code,
			Message: "sending get request: " + err.Error()}
	}

	return body, repository.ErrorResponse{Status: http.StatusOK, Message: ""}
}

func (dispatcher *NewDispatcher) GetRequest(endpoint string, header http.Header) ([]byte, repository.ErrorResponse) {
	req, err := http.NewRequest(http.MethodGet, dispatcher.url+endpoint, nil)
	if err != nil {
		return nil, repository.ErrorResponse{
			Status:  500,
			Message: fmt.Sprintf("invalid request to the servers: %s%s -> %s", dispatcher.url, endpoint, err.Error()),
		}
	}

	req.Header = header
	res, err := dispatcher.client.Do(req)
	defer res.Body.Close()
	code := res.StatusCode

	if err != nil {
		return nil, repository.ErrorResponse{
			Status:  code,
			Message: "sending get request: " + err.Error()}
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, repository.ErrorResponse{
			Status:  code,
			Message: "sending get request: " + err.Error()}
	}

	return body, repository.ErrorResponse{Status: http.StatusOK, Message: ""}
}
