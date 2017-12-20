package main

//go:generate go-bindata -o bindata.go ../build/web/ ../build/web/static/js ../build/web/static/css ../build/web/sounds

import (
	"bufio"

	"log"
	"net/http"
	"os"

	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	static "github.com/Code-Hex/echo-static"
	assetfs "github.com/elazarl/go-bindata-assetfs"
)

type Map struct {
	Name string     `json:"name"`
	Grid [][]string `json:"grid"`
}

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:3000", "http://localhost:3000"},
		AllowMethods: []string{echo.GET},
	}))

	assets := NewAssets("../build/web/")

	e.Use(static.ServeRoot("/", assets))
	e.Use(static.ServeRoot("/static", assets))

	e.GET("/api/map", func(c echo.Context) error {
		file, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		var grid [][]string
		for scanner.Scan() {
			grid = append(grid, strings.Split(scanner.Text(), ","))
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		gameMap := Map{
			Name: "test",
			Grid: grid,
		}
		return c.JSON(http.StatusOK, gameMap)
	})
	e.Logger.Fatal(e.Start(":9000"))
}

func NewAssets(root string) *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    root,
	}
}