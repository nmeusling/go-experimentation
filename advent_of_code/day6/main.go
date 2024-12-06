package main

import (
	"fmt"

	"github.com/nmeusling/go-experimentation/advent_of_code/utils"
)

var startingLocation = "^"
var blockedLocation = "#"

type location struct {
	row    int
	column int
}

type guard struct {
	loc       location
	direction direction
}

type direction struct {
	direction string
}

var up = direction{"UP"}
var right = direction{"RIGHT"}
var down = direction{"DOWN"}
var left = direction{"LEFT"}

func (d direction) rotate() direction {
	directions := map[direction]direction{up: right, right: down, down: left, left: up}
	return directions[d]
}

func main() {
	grid := utils.ReadData("day6_data")
	var g = guard{findStartingLocation(grid), up}
	visitedLocations := make(map[location]struct{})
	for !outOfBounds(grid, g.loc) {
		g = g.move(grid, visitedLocations)
	}
	fmt.Println(len(visitedLocations))
}

func findStartingLocation(grid []string) location {
	for i, row := range grid {
		for j, r := range row {
			if string(r) == startingLocation {
				return location{i, j}
			}
		}
	}
	return location{}
}

func (g guard) getNextLocation() location {
	if g.direction == up {
		return location{g.loc.row - 1, g.loc.column}
	}
	if g.direction == right {
		return location{g.loc.row, g.loc.column + 1}
	}
	if g.direction == down {
		return location{g.loc.row + 1, g.loc.column}
	}
	return location{g.loc.row, g.loc.column - 1}
}
func outOfBounds(grid []string, l location) bool {
	return l.column >= len(grid[0]) || l.column < 0 || l.row >= len(grid) || l.row < 0
}

func (g guard) move(grid []string, visited map[location]struct{}) guard {
	nextLocation := g.getNextLocation()
	visited[g.loc] = struct{}{}
	if !outOfBounds(grid, nextLocation) && string(grid[nextLocation.row][nextLocation.column]) == blockedLocation {
		g.direction = g.direction.rotate()
		return g
		// need to check if this is out of grid
	}
	g.loc.row = nextLocation.row
	g.loc.column = nextLocation.column
	return g
}
