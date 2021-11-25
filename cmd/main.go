package main

import (
	"chatserver/pkg/channel"
	"chatserver/pkg/memory"
	"chatserver/pkg/server"
	"github.com/go-playground/validator/v10"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	http.Handle("/ws", server.ChatServer{
		Service:   channel.NewService(&memory.ChatRepository{}),
		Validator: validator.New(),
	})
	log.Fatal(http.ListenAndServe(":9090", nil))
}
