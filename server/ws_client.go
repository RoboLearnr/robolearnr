package main

import (
	"time"
	"golang.org/x/net/websocket"
	"github.com/labstack/gommon/log"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	// The connection Hub
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

func (c *Client) writePump() {
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.Close()
				return
			}
			err := websocket.Message.Send(c.conn, string(message))
			if err != nil {
				log.Debug("Warning: Connection closed", err)
				c.conn.Close()
				return
			}
		}
	}
}
