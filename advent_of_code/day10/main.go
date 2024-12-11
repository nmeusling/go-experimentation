package main

import (
	"fmt"

	"github.com/nmeusling/go-experimentation/advent_of_code/utils"
)

type location struct {
	row, col, height int
}
type topMap struct {
	heights [][]location
}

func main() {
	grid := utils.ReadData("day10_sample")
	fmt.Println(grid)
}
