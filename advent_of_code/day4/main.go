package main

import (
	"fmt"
	"strings"

	"github.com/nmeusling/go-experimentation/advent_of_code/utils"
)

func main() {
	sample := utils.ReadData("day4_sample")
	fmt.Println(countXmas(sample))
	fmt.Println(countMasCrosses(sample))
	grid := utils.ReadData("day4_data")
	fmt.Println(countXmas(grid))
	fmt.Println(countMasCrosses(grid))

}

func getPotentialMatches(grid []string, row, col int) []string {
	potentialMatches := make([]string, 0)
	if col+4 <= len(grid[row]) {
		forward := grid[row][col : col+4]
		potentialMatches = append(potentialMatches, forward)
		if row+4 <= len(grid) {
			down_right := ""
			for k := 0; k < 4; k++ {
				down_right += string(grid[row+k][col+k])
			}
			potentialMatches = append(potentialMatches, down_right)
		}
	}

	if row+4 <= len(grid) {
		down := ""
		for k := 0; k < 4; k++ {
			down += string(grid[row+k][col])
		}
		potentialMatches = append(potentialMatches, down)
		if col-3 >= 0 {
			down_left := ""
			for k := 0; k < 4; k++ {
				down_left += string(grid[row+k][col-k])
			}
			potentialMatches = append(potentialMatches, down_left)
		}
	}
	return potentialMatches

}
func countXmas(grid []string) int {
	target := "XMAS"
	target_reversed := "SAMX"
	count := 0
	for i, row := range grid {
		for j := range row {
			potentialMatches := getPotentialMatches(grid, i, j)
			for _, potentialMatch := range potentialMatches {
				if potentialMatch == target || potentialMatch == target_reversed {
					count++
				}
			}

		}
	}
	return count
}

// look for a; get 4 diagonals, 2 s , 2 m, same row or column
func countMasCrosses(grid []string) int {
	count := 0
	for i, row := range grid {
		for j := range row {
			if string(grid[i][j]) != "A" {
				continue
			}
			if i-1 >= 0 && i+1 < len(grid) && j-1 >= 0 && j+1 < len(row) {
				corners := string(grid[i-1][j-1]) + string(grid[i-1][j+1]) + string(grid[i+1][j-1]) + string(grid[i+1][j+1])
				if strings.Count(corners, "S") != 2 || strings.Count(corners, "M") != 2 || corners[1] == corners[2] {
					continue
				}
				count += 1
			}
		}
	}
	return count
}
