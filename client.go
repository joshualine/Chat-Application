package main

import (
	"github.com/gorilla/websocket"
)

// client represents a single chatting user.
type client struct {
	socket *websocket.Conn // socket is the web socket for this client.
	send   chan []byte     // send is a channel on which messages are sent.
	room   *room           // room is the room this client is chatting in.
}

func (c *client) read()  {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}
func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}