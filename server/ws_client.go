package main

import (
	"time"

	"fmt"
	"golang.org/x/net/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

func (c *Client) writePump() {
	for {
		select {
		case message := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			fmt.Println("Sending message", string(message))
			err := websocket.Message.Send(c.conn, string(message))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
