package main

import (
	"os"
	"bufio"
	"strings"
	"log"
)

type Map struct {
	Name string     `json:"name"`
	Grid [][]string `json:"grid"`
	Car *Car        `json:"car"`
}



func (m *Map) Forward() {
	newPosition := m.findNewForwardPosition()
	x, y := newPosition[0], newPosition[1]

	if x < 0 || x >= len(m.Grid) || y < 0 || y >= len(m.Grid[x]) {
		log.Printf("Out of bound: %v\n", newPosition)
		return
	}
	if m.Grid[x][y] == "w" {
		log.Printf("Hitting a wall: %s\n", newPosition)
		return
	}

	m.swapTiles(m.Car.Position, newPosition)
}

func (m *Map) findNewForwardPosition() []int {
	x, y := m.Car.Position[0], m.Car.Position[1]

	switch m.Car.Rotation {
	case 0: // right
		return []int{x, y+1}
	case 90: // down
		return []int{x+1, y}
	case 180: // left
		return []int{x, y-1}
	case 270:
		return []int{x-1, y}
	}

	return m.Car.Position
}

func (m *Map) swapTiles(oldPosition []int, newPosition []int) {
	oldX, oldY := oldPosition[0], oldPosition[1]
	newX, newY := newPosition[0], newPosition[1]

	oldTile := m.Grid[oldX][oldY]+""

	m.Grid[oldX][oldY] = m.Grid[newX][newY]+""
	m.Grid[newX][newY] = oldTile+""

	m.Car.Position = []int{newX, newY}
}

func (m *Map) Rotate() {
	m.Car.Rotate()
}

func (m *Map) InitCar() {
	var position = []int{}
	for x, row := range m.Grid {
		for y, e := range row {
			if e == "c" {
				position = []int{x, y}
			}
		}
	}

	m.Car = &Car{Position: position, Rotation: 0}
}

type Car struct {
	Position []int
	Rotation int `json:"rotation"`
}

func (c *Car) Rotate() {
	c.Rotation += 90
	if c.Rotation >= 360 {
		c.Rotation = 0
	}
}

func NewMap(name string, grid [][]string) *Map {
	m := new(Map)

	m.Name = name
	m.Grid = grid
	m.InitCar()

	return m
}

func LoadMap(filename string) *Map {
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



	mapI := NewMap(filename, grid)

	return mapI
}
