package main

import (
	"os"
	"fmt"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	hub := newHub()
	go hub.run()

	mapInstance := LoadMap(os.Args[1])

	e := CreateServer(hub, mapInstance)

	fmt.Println("Starting...")
	err := open.Run("http://127.0.0.1:9000/")
	if err != nil {
		fmt.Println("Open http://127.0.0.1:9000 in your favorite browser.")
	}

	e.Logger.Fatal(e.Start(":9000"))
}
