package main

import (
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

func HandleWs(hub *Hub) echo.HandlerFunc {
	return func(c echo.Context) error {
		serveWs(hub, c)

		return nil
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
