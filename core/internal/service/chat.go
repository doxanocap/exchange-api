package service

import "webchat/pkg/repository"

type ChatService struct {
	chatModels repository.Chat
}

func NewChatServices(repo *repository.Repository) *ChatService {
	return &ChatService{chatModels: repo.Chat}
}
