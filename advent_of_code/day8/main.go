package main

import (
	"fmt"

	"github.com/nmeusling/go-experimentation/advent_of_code/utils"
)

func main() {
	fmt.Println(countAntinodes("day8_sample"))
	fmt.Println(countAntinodes("day8_data"))
	fmt.Println(countAllAntinodes("day8_sample"))
	fmt.Println(countAllAntinodes("day8_data"))
}

func countAntinodes(fileName string) int {
	grid := utils.ReadData(fileName)
	antennas := getInput(grid)
	antinodes := make(map[location]struct{}, 0)
	for _, a := range antennas {
		combos := getCombinations(a)
		for _, combo := range combos {
			antis := combo.getAntinodes(len(grid), len(grid[0]))
			for _, anti := range antis {
				antinodes[anti] = struct{}{}
			}
		}
	}
	return len(antinodes)
}

func countAllAntinodes(fileName string) int {
	grid := utils.ReadData(fileName)
	antennas := getInput(grid)
	antinodes := make(map[location]struct{}, 0)
	for _, a := range antennas {
		combos := getCombinations(a)
		for _, combo := range combos {
			antis := combo.getAllPointsOnLine(len(grid), len(grid[0]))
			for _, anti := range antis {
				antinodes[anti] = struct{}{}
			}
		}
	}
	return len(antinodes)
}

type location struct {
	row    int
	column int
}

type line struct {
	point1 location
	point2 location
}

func getInput(grid []string) map[string][]location {
	antennas := make(map[string][]location, 0)
	for i, row := range grid {
		for j, col := range row {
			char := string(col)
			if char != "." {
				antennas[char] = append(antennas[char], location{i, j})
			}
		}
	}
	return antennas
}

func getCombinations(locs []location) []line {
	lines := []line{}
	for i, loc := range locs {
		for j := i + 1; j < len(locs); j++ {
			lines = append(lines, line{loc, locs[j]})
		}
	}
	return lines
}

func (l1 location) minus(l2 location) location {
	return location{l1.row - l2.row, l1.column - l2.column}
}

func (l1 location) plus(l2 location) location {
	return location{l2.row + l1.row, l2.column + l1.column}
}

func (l line) isDouble(point location) bool {
	if l.point1.row-point.row == 2*(l.point2.row-point.row) {
		if l.point1.column-point.column == 2*(l.point2.column-point.column) {
			return true
		}
	}
	if 2*(l.point1.row-point.row) == l.point2.row-point.row {
		if 2*(l.point1.column-point.column) == l.point2.column-point.column {
			return true
		}
	}
	return false
}

func (l location) isInBounds(numRows, numCols int) bool {
	if l.column < 0 || l.column >= numCols {
		return false
	}
	if l.row < 0 || l.row >= numRows {
		return false
	}
	return true
}

func (l line) getPointsOnLine(numRows, numCols int) []location {
	points := []location{}
	slope := l.point2.minus(l.point1)
	for testPoint := l.point1.minus(slope); testPoint.isInBounds(numRows, numCols); testPoint = testPoint.minus(slope) {
		points = append(points, testPoint)
	}
	for testPoint := l.point2.plus(slope); testPoint.isInBounds(numRows, numCols); testPoint = testPoint.plus(slope) {
		points = append(points, testPoint)
	}
	return points
}

func (l line) getAllPointsOnLine(numRows, numCols int) []location {
	points := []location{l.point1}
	slope := l.point2.minus(l.point1)
	for testPoint := l.point1.minus(slope); testPoint.isInBounds(numRows, numCols); testPoint = testPoint.minus(slope) {
		points = append(points, testPoint)
	}
	for testPoint := l.point1.plus(slope); testPoint.isInBounds(numRows, numCols); testPoint = testPoint.plus(slope) {
		points = append(points, testPoint)
	}
	return points
}

func (l line) getAntinodes(numRows, numCols int) []location {
	possibilities := l.getPointsOnLine(numRows, numCols)
	antinodes := []location{}
	for _, possibility := range possibilities {
		if l.isDouble(possibility) {
			antinodes = append(antinodes, possibility)
		}
	}
	return antinodes
}
