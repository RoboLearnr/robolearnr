package main

import "os"

func main() {
	hub := newHub()
	go hub.run()

	mapInstance := LoadMap(os.Args[1])

	e := CreateServer(hub, mapInstance)
	e.Logger.Fatal(e.Start(":9000"))
}
