package main

import (
	"github.com/gorilla/websocket"
	"fmt"
)

type Client struct {
	id     string
	socket *websocket.Conn
	send   chan []byte
}

func (client *Client) write() {
	defer func() {
		client.socket.Close()
	}()

	for {
		select {
		case message := <-client.send:
			client.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func (client *Client) read() {

	for {
		_, message, _ := client.socket.ReadMessage()
		fmt.Printf("\n From: %s, Message: %s", client.id, message)
		room.broadcast <- message
	}

}