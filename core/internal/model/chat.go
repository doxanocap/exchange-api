package model

type ChatData struct {
	ChatId     int  `json:"chat_id"`
	SenderId   int  `json:"sender_id"`
	ReceiverId int  `json:"receiver_id"`
	Blocked    bool `json:"blocked"`
}

type ChatMessage struct {
	ChatId   int    `json:"chat_id"`
	SenderId int    `json:"sender_id"`
	SentAt   int64  `json:"sent_at"`
	Message  []byte `json:"message"`
}
