package main

import (
	"fmt"

	"github.com/nmeusling/go-experimentation/advent_of_code/utils"
)

func main() {
	sample := utils.ReadData("day4_sample")
	fmt.Println(countXmas(sample))
	grid := utils.ReadData("day4_data")
	fmt.Println(countXmas(grid))

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
