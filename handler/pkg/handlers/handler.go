package handlers

import "handler/pkg/services"

type Handler struct {
	services *services.Services
}

func InitHandler(services *services.Services) *Handler {
	return &Handler{services: services}
}
