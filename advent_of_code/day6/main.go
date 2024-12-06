package main

import (
	"fmt"
	"slices"

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
	startLoc  location
	startDir  direction
	steps     int
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
	fmt.Println(solvePart1("day6_sample"))
	fmt.Println(solvePart1("day6_data"))
	fmt.Println(solvePart2("day6_sample"))
	fmt.Println(solvePart2("day6_data"))
}
func solvePart1(fileName string) int {
	grid := utils.ReadData(fileName)
	start := findStartingLocation(grid)
	var g = guard{start, up, start, up, 0}
	visitedLocations := make(map[location][]direction)
	for !outOfBounds(grid, g.loc) {
		g = g.move(grid, visitedLocations)
	}
	return len(visitedLocations)
}
func solvePart2(fileName string) int {
	grid := utils.ReadData(fileName)
	start := findStartingLocation(grid)
	var g = guard{start, up, start, up, 0}
	return countClosed(grid, g)
}

func (g *guard) resetPosition() {
	g.loc = g.startLoc
	g.direction = g.startDir
	g.steps = 0
}
func countClosed(grid []string, g guard) int {
	count := 0
	for i := range grid {
		for j := range grid[i] {
			g.resetPosition()
			// can't use guards starting position
			if g.startLoc.row == i && g.startLoc.column == j {
				continue
			}
			if string(grid[i][j]) == blockedLocation {
				continue
			}
			grid[i] = grid[i][:j] + blockedLocation + grid[i][j+1:]
			if isClosedLoop(grid, g) {
				count++
			}
			grid[i] = grid[i][:j] + "." + grid[i][j+1:]

		}
	}
	return count
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
func isClosedLoop(grid []string, g guard) bool {
	visited := make(map[location][]direction)
	oob := outOfBounds(grid, g.loc)
	returned := g.hasReturned()
	for !oob && !returned {
		g = g.move(grid, visited)
		if outOfBounds(grid, g.loc) {
			return false
		}
		if g.hasReturned() {
			return true
		}
		if g.alreadyVisited(visited) {
			return true
		}
	}
	return false
}

func outOfBounds(grid []string, l location) bool {
	return l.column >= len(grid[0]) || l.column < 0 || l.row >= len(grid) || l.row < 0
}

func (g guard) alreadyVisited(visited map[location][]direction) bool {
	return slices.Contains(visited[g.loc], g.direction)
}
func (g guard) hasReturned() bool {
	return g.steps > 0 && g.loc == g.startLoc && g.direction == g.startDir
}

func (g guard) move(grid []string, visited map[location][]direction) guard {
	nextLocation := g.getNextLocation()
	g.steps++
	visited[g.loc] = append(visited[g.loc], g.direction)
	if !outOfBounds(grid, nextLocation) && string(grid[nextLocation.row][nextLocation.column]) == blockedLocation {
		g.direction = g.direction.rotate()
		return g
		// need to check if this is out of grid
	}
	g.loc.row = nextLocation.row
	g.loc.column = nextLocation.column
	return g
}
