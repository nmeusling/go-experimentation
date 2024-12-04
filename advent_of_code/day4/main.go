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

func countXmas(grid []string) int {
	target := "XMAS"
	target_reversed := "SAMX"
	count := 0
	for i, row := range grid {
		for j := range row {
			// forward/backward match
			if j+4 <= len(row) {
				forward := grid[i][j : j+4]
				if forward == target || forward == target_reversed {
					count++
				}
			}

			//down/up match
			if i+4 <= len(grid) {
				down := ""
				for k := 0; k < 4; k++ {
					down += string(grid[i+k][j])
				}
				if down == target || down == target_reversed {
					count++
				}
			}

			//down right/up left
			if j+4 <= len(row) && i+4 <= len(grid) {
				down_right := ""
				for k := 0; k < 4; k++ {
					down_right += string(grid[i+k][j+k])
				}
				if down_right == target || down_right == target_reversed {
					count++
				}
			}

			//down left / up right
			if j-3 >= 0 && i+4 <= len(grid) {
				down_right := ""
				for k := 0; k < 4; k++ {
					down_right += string(grid[i+k][j-k])
				}
				// fmt.Println(down_right)
				if down_right == target || down_right == target_reversed {
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
