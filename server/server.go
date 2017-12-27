package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func CreateServer(hub *Hub) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT},
	}))

	e.Use(HandleStatic())

	e.GET("/api/map", HandleMap())
	e.GET("/api/:action", HandleRpc(hub))
	e.GET("/ws", HandleWs(hub))

	return e
}
