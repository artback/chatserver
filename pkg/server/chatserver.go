package server

import (
	"chatserver/pkg/chat"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
)

type ChatServer struct {
	Service chat.Service
	Validator
}

func (c ChatServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	id := rand.Int63()
	recv := c.Service.Connect("messages", id)
	defer c.Service.Disconnect("messages", id)
	go func() {
		for msg := range recv {
			ws.WriteJSON(&msg)
		}
	}()
	for {
		var msg chat.Message
		if err := ws.ReadJSON(&msg); err != nil {
			fmt.Printf("ServeHTTP: %v", err)
			continue
		}
		if err = c.Validator.Struct(&msg); err != nil {
			fmt.Printf("ServeHTTP: %v", err)
			continue
		}
		c.Service.Broadcast("messages",msg)
	}
}

type Validator interface {
	Struct(s interface{}) error
}
