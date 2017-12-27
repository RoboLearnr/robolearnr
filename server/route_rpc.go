package main

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
)

func HandleRpc(hub *Hub) echo.HandlerFunc {
	return func(c echo.Context) error {
		msg, _ := json.Marshal(Action{Action: c.Param("action")})

		hub.broadcast <- msg

		return c.JSON(http.StatusOK, nil)
	}
}
