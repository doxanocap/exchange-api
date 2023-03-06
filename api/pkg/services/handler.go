package services

import "api/pkg/dispatcher"

type HandlerService struct {
}

func NewHandlerService(dp *dispatcher.Dispatcher) *HandlerService {
	return &HandlerService{}
}
