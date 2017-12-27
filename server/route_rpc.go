package main

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
)

func HandleRpc(hub *Hub, mapInstance *Map) echo.HandlerFunc {
	return func(c echo.Context) error {
		switch c.Param("action") {
		case "forward":
			mapInstance.Forward()
			break
		case "rotate":
			mapInstance.Rotate()
		}

		msg, _ := json.Marshal(Action{Action: "map", Map: mapInstance})

		hub.broadcast <- msg

		return c.JSON(http.StatusOK, nil)
	}
}

type Action struct {
	Action string `json:"action"`
	Map    *Map   `json:"map"`
}