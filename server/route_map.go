package main

import (
	"github.com/labstack/echo"
	"net/http"
	"encoding/json"
	"os"
)

func HandleMap(mapInstance *Map) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, mapInstance)
	}
}

func HandleInfo(mapInstance *Map) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, mapInstance.Info())
	}
}

func HandleReset(hub *Hub, mapInstance *Map) echo.HandlerFunc {
	return func(c echo.Context) error {
		*mapInstance = *LoadMap(os.Args[1])

		msg, _ := json.Marshal(Action{Action: "map", Map: mapInstance})
		hub.broadcast <- msg

		return c.JSON(http.StatusOK, mapInstance.Info())
	}
}
