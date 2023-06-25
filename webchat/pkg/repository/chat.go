package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"webchat/pkg/repository/postgres"
)

type ChatModel struct {
	psql *sqlx.DB
}

func NewChatModel(db *sqlx.DB) *ChatModel {
	return &ChatModel{psql: db}
}

type ChatData struct {
	ChatId     int  `json:"chat_id"`
	SenderId   int  `json:"sender_id"`
	ReceiverId int  `json:"receiver_id"`
	Blocked    bool `json:"blocked"`
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

func (ch *ChatModel) RecordMessage(chatId int, message []byte) {

}