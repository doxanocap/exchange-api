package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"
	"time"
	"webchat/pkg/configs"
	"webchat/pkg/repository"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type ClientService struct {
	chatModels   repository.Chat
	clientModels repository.Client
}

type ClientType struct {
	mu   sync.RWMutex
	ctx  *gin.Context
	pool *PoolService
	conn *websocket.Conn
	send chan []byte
}

func NewClientService(repo *repository.Repository) *ClientService {
	return &ClientService{
		chatModels:   repo.Chat,
		clientModels: repo.Client,
	}
}

func (s *ClientService) NewClient(context *gin.Context, pool *PoolService, conn *websocket.Conn) *ClientType {
	return &ClientType{
		mu:   sync.RWMutex{},
		ctx:  context,
		pool: pool,
		conn: conn,
		send: make(chan []byte),
	}
}

func (c *ClientType) Reader() {
	defer func() {
		c.pool.unregister <- c
		if err := c.conn.Close(); err != nil {
			log.Println("reader -> closing websocket connection: ", err.Error())
			return
		}
	}()

	c.setReadParams()
	for {
		message, err := c.readMessage()
		if err != nil {
			return
		}
		fmt.Println("[ ws ]", c.ctx.Request.RemoteAddr, string(message))
		c.pool.broadcast <- message
	}
}

func (c *ClientType) Writer() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		if err := c.conn.Close(); err != nil {
			log.Println("writer -> closing websocket connection: ", err.Error())
			return
		}
	}()

	for {
		select {
		case message, ok := <-c.send:
			err := c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				log.Println("setting write deadline: ", err.Error())
				return
			}

			// if channel is closed then pool closed connection
			if !ok {
				err := c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					log.Println("closing chan: ", err.Error())
					return
				}
				return
			}

			err = c.writeMessage(message)
			if err != nil {
				log.Println("writing msg: ", err)
				return
			}

		case <-ticker.C:
			err := c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				log.Println("setting write deadline: ", err.Error())
				return
			}

			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *ClientType) ExtractUser() (repository.User, error) {
	cookie, err := c.ctx.Cookie("jwt")
	if err != nil {
		return repository.User{}, err
	}
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.ENV("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return repository.User{}, err
	}

	claims := token.Claims.(*jwt.RegisteredClaims)
	log.Println(0)
	var user repository.User
	if err := json.Unmarshal([]byte(claims.Issuer), &user); err != nil {
		return user, err
	}

	return user, nil
}

func (c *ClientType) readMessage() ([]byte, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, message, err := c.conn.ReadMessage()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
			log.Printf("error: %v, user-agent: %v", err, c.ctx.GetHeader("User-Agent"))
		}
		return nil, err
	}
	message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
	return message, nil
}

func (c *ClientType) writeMessage(message []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	w, err := c.conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return err
	}

	_, err = w.Write(message)
	if err != nil {
		return err
	}

	// add queued chat messages to the current websocket message.
	n := len(c.send)
	for i := 0; i < n; i++ {
		msg := <-c.send
		msg = append(newline, msg...)
		if _, err := w.Write(message); err != nil {
			return err
		}
	}

	err = w.Close()
	return err
}

func (c *ClientType) setReadParams() {
	c.conn.SetReadLimit(maxMessageSize)
	err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		log.Println("setReadDeadline:", err)
		return
	}
	c.conn.SetPongHandler(func(string) error {
		err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
		if err != nil {
			log.Println("setPongHandler:", err)
			return err
		}
		return nil
	})
}
