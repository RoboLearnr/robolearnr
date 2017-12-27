package main

//go:generate go-bindata -o bindata.go ../build/web/ ../build/web/static/js ../build/web/static/css ../build/web/sounds

import (
	"bufio"

	"log"
	"net/http"
	"os"
	"encoding/json"

	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	static "github.com/Code-Hex/echo-static"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"golang.org/x/net/websocket"
)

type Map struct {
	Name string     `json:"name"`
	Grid [][]string `json:"grid"`
}

type Action struct {
	Action string `json:"action"`
}

var messageBus chan string = make(chan string)


func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT},
	}))

	assets := NewAssets("../build/web/")

	e.Use(static.ServeRoot("/", assets))
	e.Use(static.ServeRoot("/static", assets))

	e.GET("/api/map", func(c echo.Context) error {

		return c.JSON(http.StatusOK, LoadMap(os.Args[1]))
	})
	hub := newHub()
	go hub.run()
	e.GET("/api/forward", func(c echo.Context) error {
		msg, _ := json.Marshal(Action{Action: "forward",})

		hub.broadcast <- msg

		return c.JSON(http.StatusOK, nil)
	})
	e.GET("/api/rotate", func(c echo.Context) error {
		msg, _ := json.Marshal(Action{Action: "rotate",})

		hub.broadcast <- msg

		return c.JSON(http.StatusOK, nil)
	})

	e.GET("/ws", func (c echo.Context) error {
		serveWs(hub, c)


		return nil
	})

	e.Logger.Fatal(e.Start(":9000"))

}

func LoadMap(filename string) Map {
	file, err := os.Open(filename)
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

	return Map{
		Name: "test",
		Grid: grid,
	}
}

func NewAssets(root string) *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    root,
	}
}

var connectionPool []websocket.Conn

//func socket(c echo.Context) error {
//	websocket.Handler(func(conn *websocket.Conn) {
//		defer conn.Close()
//
//		for {
//			msg := <- messageBus
//			fmt.Println("send", msg)
//			// Write
//			err := websocket.Message.Send(conn, msg)
//			if err != nil {
//				conn.Close()
//				c.Logger().Error(err)
//			}
//
//		}
//	}).ServeHTTP(c.Response(), c.Request())
//
//	return nil
//}