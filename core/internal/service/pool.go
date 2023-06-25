package service

import (
	"encoding/json"
	"fmt"
	"time"
	"webchat/pkg/repository"

	log "github.com/sirupsen/logrus"
)

type PoolService struct {
	chatModel  repository.Chat
	broadcast  chan []byte
	register   chan *ClientType
	unregister chan *ClientType
	clients    map[*ClientType]repository.User
}

type Members struct {
	member1 int
	member2 int
}

var _m = map[Members]repository.ChatData{}

func NewPoolService(repo *repository.Repository) *PoolService {
	ps := &PoolService{
		chatModel:  repo.Chat,
		broadcast:  make(chan []byte),
		register:   make(chan *ClientType),
		unregister: make(chan *ClientType),
		clients:    map[*ClientType]repository.User{},
	}
	go ps.Start()
	return ps
}

type BroadcastMessage struct {
	SenderId   int    `json:"sender_id"`
	ReceiverId int    `json:"receiver_id"`
	Message    string `json:"message"`
}

func (p *PoolService) Start() {
	fmt.Println("[ ws ] Pool started succesfully!")
	for {
		select {
		case client := <-p.register:
			user, err := client.ExtractUser()
			if err != nil {
				log.Errorf("handling new client: ", err.Error())
				return
			}

			p.clients[client] = user

		case client := <-p.unregister:
			if _, ok := p.clients[client]; ok {
				delete(p.clients, client)
				close(client.send)
			}

		case message := <-p.broadcast:
			var data BroadcastMessage
			if err := json.Unmarshal(message, &data); err != nil {
				log.Errorf("invalid request body: ", err.Error())
				return
			}

			chData := p.ExtractData(data)
			chMsg := repository.ChatMessage{
				ChatId:   chData.ChatId,
				SenderId: data.SenderId,
				SentAt:   time.Now().Unix(),
				Message:  message,
			}

			if err := p.chatModel.RecordMessage(chMsg); err != nil {
				log.Errorf("writing message in the pool", err.Error())
				return

			}
			
			// record msg to database by chat id
			for client, user := range p.clients {
				if user.Id == data.SenderId || user.Id == data.ReceiverId {
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
