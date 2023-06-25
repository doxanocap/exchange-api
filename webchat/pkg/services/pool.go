package services

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"webchat/pkg/repository"
)

type PoolService struct {
	chatModel  repository.Chat
	broadcast  chan []byte
	register   chan *ClientType
	unregister chan *ClientType
	clients    map[*ClientType]bool
}

type Members struct {
	member1 int
	member2 int
}

var _m = map[Members]repository.ChatData{}

func NewPoolService(repo *repository.Repository) *PoolService {
	return &PoolService{
		chatModel:  repo.Chat,
		broadcast:  make(chan []byte),
		register:   make(chan *ClientType),
		unregister: make(chan *ClientType),
		clients:    map[*ClientType]bool{},
	}
}

type BroadcastMessage struct {
	SenderId   int    `json:"sender_id"`
	ReceiverId int    `json:"receiver_id"`
	Message    string `json:"message"`
}

func (p *PoolService) Start() {
	for {
		select {
		case client := <-p.register:
			p.clients[client] = true

		case client := <-p.unregister:
			if _, ok := p.clients[client]; ok {
				delete(p.clients, client)
				close(client.send)
			}

		case message := <-p.broadcast:
			var data BroadcastMessage
			if err := json.Unmarshal(message, &data); err != nil {
				log.Println("invalid request: ", err.Error())
				continue
			}

			chData := p.ExtractData(data)
			p.chatModel.RecordMessage(chData.ChatId, message)
			// record msg to database by chat id

			for client := range p.clients {
				if client.id == data.SenderId || client.id == data.ReceiverId {
					select {
					case client.send <- message:
					default:
						close(client.send)
						delete(p.clients, client)
					}
				}
			}
		}
	}
}

func (p *PoolService) ExtractData(data BroadcastMessage) repository.ChatData {
	m1 := Members{data.SenderId, data.ReceiverId}
	m2 := Members{data.ReceiverId, data.SenderId}

	var chData repository.ChatData
	val1, ok1 := _m[m1]
	val2, ok2 := _m[m2]
	if ok1 {
		chData = val1
	} else if ok2 {
		chData = val2
	} else {
		p.chatModel.Create(data.SenderId, data.ReceiverId)
	}
	return chData
}
