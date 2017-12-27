package main

type Map struct {
	Name string     `json:"name"`
	Grid [][]string `json:"grid"`
}

type Action struct {
	Action string `json:"action"`
	Map    Map    `json:"map"`
}

func main() {
	hub := newHub()
	go hub.run()

	e := CreateServer(hub)
	e.Logger.Fatal(e.Start(":9000"))
}
