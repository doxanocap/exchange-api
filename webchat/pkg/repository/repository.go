package repository

import (
	"github.com/jmoiron/sqlx"
)

type Client interface {
}

type Pool interface {
}

type Chat interface {
	Create(member1, member2 int) ChatData
	RecordMessage(chatId int, message []byte)
}

type Repository struct {
	Client
	Pool
	Chat *ChatModel
}

func InitRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Client: NewClientModel(db),
		Pool:   NewPoolModel(db),
		Chat:   NewChatModel(db),
	}
}
