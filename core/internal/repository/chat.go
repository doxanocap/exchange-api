package repository

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"webchat/pkg/repository/postgres"
)

type ChatModel struct {
	pool *pgxpool.Pool
	log  *zap.Logger
	builder
}

func NewChatModel(db *sqlx.DB) *ChatModel {
	return &ChatModel{psql: db}
}

func (ch *ChatModel) Create(member1, member2 int) ChatData {
	query := fmt.Sprintf(`
		INSERT INTO %s 
		(member1_id, member2_id, blocked)
		VALUES (%d,%d,false)
    `, postgres.ChatListTable, member1, member2)
	var data ChatData
	if err := ch.psql.Select(data, query); err != nil {
		log.Println(err)
		return ChatData{}
	}
	return data
}

func (ch *ChatModel) RecordMessage(msg ChatMessage) error {
	query := fmt.Sprintf(`
		INSERT INTO %s 
		(chat_id, sender_id, sent_at, message)
		VALUES (%d,%d,%d,'%s')
    `, postgres.ChatMessages, msg.ChatId, msg.SenderId, msg.SentAt, msg.Message)

	var data ChatData
	if err := ch.psql.Select(data, query); err != nil {
		return err
	}
	return nil
}
