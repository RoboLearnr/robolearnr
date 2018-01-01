package main

import (
	"os"
	"bufio"
	"strings"
	"log"
	"net/http"
)

type Map struct {
	Name string     `json:"name"`
	Grid [][]string `json:"grid"`
	Car *Car        `json:"car"`

	originalGrid [][]string
	goal []int
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

	m.moveCar(m.Car.Position, newPosition)
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

func (m *Map) moveCar(oldPosition []int, newPosition []int) {
	oldX, oldY := m.Car.Position[0], m.Car.Position[1]
	newX, newY := newPosition[0], newPosition[1]

	carTile := m.Grid[oldX][oldY]+""
	m.Grid[oldX][oldY] = m.originalGrid[oldX][oldY]+""
	m.Grid[newX][newY] = carTile+""

	m.Car.Position = []int{newX, newY}
}

func (m *Map) Rotate() {
	m.Car.Rotate()
}

func (m *Map) Init() {
	var carPosition, goalPosition = []int{}, []int{}
	for x, row := range m.Grid {
		for y, e := range row {
			if e == "c" {
				carPosition = []int{x, y}
			}
		}
	}

	m.Car = &Car{Position: carPosition, Rotation: 270}
	m.goal = goalPosition

	m.originalGrid = make([][]string, len(m.Grid))
	for i := range m.Grid {
		m.originalGrid[i] = make([]string, len(m.Grid[i]))
		copy(m.originalGrid[i], m.Grid[i])
	}
	m.originalGrid[carPosition[0]][carPosition[1]] = "e"
}

func (m *Map) Info() MapInfo {
	x, y := m.Car.Position[0], m.Car.Position[1]
	onGoal := m.originalGrid[x][y] == "g"

	newPosition := m.findNewForwardPosition()
	newX, newY := newPosition[0], newPosition[1]
	nextTile := m.originalGrid[newX][newY]
	beforeObstacle := nextTile == "w"

	return MapInfo{BeforeObstacle: beforeObstacle, OnGoal: onGoal}
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

type MapInfo struct {
	BeforeObstacle bool `json:"before_obstacle"`
	OnGoal         bool `json:"on_goal"`
}

func NewMap(name string, grid [][]string) *Map {
	m := new(Map)

	m.Name = name
	m.Grid = grid
	m.Init()

	return m
}

func LoadMap(path string) *Map {
	var grid [][]string

	scanner := bufio.NewScanner(GetMapReader(path))
	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ","))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return NewMap(path, grid)
}

func GetMapReader(path string) *bufio.Reader {
	if strings.HasPrefix(path, "http") {
		response, err := http.Get(path)
		if err != nil {
			log.Fatal(err)
		}
		return bufio.NewReader(response.Body)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewReader(file)
}
