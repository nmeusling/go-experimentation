package main

import (
	"fmt"
	"slices"

	"github.com/nmeusling/go-experimentation/advent_of_code/utils"
)

func main() {
	grid := utils.ReadData("day12_sample")
	r := getRegions(grid)
	fmt.Println(calcCost(r, grid))
	grid2 := utils.ReadData("day12_data")
	r2 := getRegions(grid2)
	fmt.Println(calcCost(r2, grid2))
}

func calcCost(regions []region, grid []string) int {
	cost := 0
	for _, r := range regions {
		cost += r.getArea() * r.getPerimeter(grid)
	}
	return cost
}

func getRegions(grid []string) []region {
	regions := []region{}
	assigned := []location{}
	for row := range grid {
		for column := range grid[row] {
			loc := location{row, column}
			if !slices.Contains(assigned, loc) {
				region := getRegion(grid, loc)
				assigned = append(assigned, region.locations...)
				regions = append(regions, region)
			}

		}
	}
	return regions
}
func getRegion(grid []string, start location) region {
	r := region{}
	r.locations = append(r.locations, start)
	regionType := grid[start.row][start.col]
	toVisit := start.getAdjacentLocations(grid)
	visited := []location{}
	for len(toVisit) > 0 {
		loc := toVisit[0]
		toVisit = toVisit[1:]
		visited = append(visited, loc)
		adjacent := loc.getAdjacentLocations(grid)
		if grid[loc.row][loc.col] == regionType {
			for _, a := range adjacent {
				if !slices.Contains(toVisit, a) && !slices.Contains(visited, a) {
					toVisit = append(toVisit, a)
				}
			}
			if !slices.Contains(r.locations, loc) {
				r.locations = append(r.locations, loc)
				continue
			}
		}
	}
	return r
}

func (l location) getAdjacentLocations(grid []string) []location {
	adjacent := []location{}
	if l.row > 0 {
		adjacent = append(adjacent, location{l.row - 1, l.col})
	}
	if l.row < len(grid)-1 {
		adjacent = append(adjacent, location{l.row + 1, l.col})

	}
	if l.col > 0 {
		adjacent = append(adjacent, location{l.row, l.col - 1})
	}
	if l.col < len(grid[0])-1 {
		adjacent = append(adjacent, location{l.row, l.col + 1})

	}
	return adjacent
}

type location struct {
	row, col int
}

func (l location) getPerimeter(grid []string) int {
	perimeter := 4
	regionType := grid[l.row][l.col]
	adjacent := l.getAdjacentLocations(grid)
	for _, adj := range adjacent {
		if grid[adj.row][adj.col] == regionType {
			perimeter -= 1
		}
	}
	return perimeter
}

type region struct {
	locations []location
}

func (r region) getArea() int {
	return len(r.locations)
}

func (r region) getPerimeter(grid []string) int {
	per := 0
	for _, r := range r.locations {
		per += r.getPerimeter(grid)
	}
	return per
}
