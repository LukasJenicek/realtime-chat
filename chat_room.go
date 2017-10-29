package main

import "encoding/json"

type ChatRoom struct {
	clients   map[*Client]bool
	newClient chan *Client
	broadcast chan []byte
}

func (room *ChatRoom) send(message []byte, from *Client) {
	for conn := range room.clients {
		if conn != from {
			conn.send <- message
		}
	}
}

func (room *ChatRoom) newChatRoom() *ChatRoom {
	return &ChatRoom{
		clients:   make(map[*Client]bool),
		newClient: make(chan *Client),
		broadcast: make(chan []byte),
	}
}

func (room *ChatRoom) start() {
	for {
		select {
		case conn := <-room.newClient:
			room.clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "New socket has connected"})
			room.send(jsonMessage, conn)
		case message := <-room.broadcast:
			for conn := range room.clients {
				conn.send <- message
			}
		}
	}
}