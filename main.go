package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"fmt"
	"chat/utils"
)

type Message struct {
	Sender    string
	Recipient string
	Content   string
}

// In Go, top-level variable assignments must be prefixed with the var keyword. Omitting the var keyword is only allowed within blocks.
var room = ChatRoom{
	clients:   make(map[*Client]bool),
	newClient: make(chan *Client),
	broadcast: make(chan []byte),
}

func main() {

	go room.start()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	})

	http.HandleFunc("/ws", handleWebSocketCommunication)
	http.ListenAndServe(":5000", nil)
}

func handleWebSocketCommunication(res http.ResponseWriter, req *http.Request) {
	uuid := utils.GenerateUUID(32)

	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)

	if err != nil {
		fmt.Printf("something bad happened")
		return
	}

	client := &Client{id: uuid, socket: conn, send: make(chan []byte)}

	room.newClient <- client

	go client.write()
	go client.read()
}

