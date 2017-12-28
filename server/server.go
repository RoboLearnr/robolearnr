package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func CreateServer(hub *Hub, mapInstance *Map) *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT},
	}))

	e.Use(HandleStatic())

	e.GET("/api/map", HandleMap(mapInstance))
	e.GET("/api/info", HandleInfo(mapInstance))
	e.GET("/api/reset", HandleReset(hub, mapInstance))
	e.GET("/api/:action", HandleRpc(hub, mapInstance))
	e.GET("/ws", HandleWs(hub))

	return e
}

