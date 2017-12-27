package main

import (
	"bufio"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"os"
	"strings"
)

func HandleMap() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, LoadMap(os.Args[1]))
	}
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
