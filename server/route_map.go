package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func HandleMap(mapInstance *Map) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, mapInstance)
	}
}
