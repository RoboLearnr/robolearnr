package main

import (
	"time"

	"golang.org/x/net/websocket"
	"fmt"
	"github.com/labstack/echo"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	for {
		select {
		case message := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			fmt.Println("Sending message" , string(message))
			err := websocket.Message.Send(c.conn, string(message))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func serveWs(hub *Hub, c echo.Context) {
	websocket.Handler(func(conn *websocket.Conn) {
		defer conn.Close()

		client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
		client.hub.register <- client
		hub.register <- client

		client.writePump()

	}).ServeHTTP(c.Response(), c.Request())


}