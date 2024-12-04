package main

import (
	"fmt"

	"github.com/nmeusling/go-experimentation/advent_of_code/utils"
)

func main() {
	grid := utils.ReadData("day4_sample")
	fmt.Println(countXmas(grid))

}

func countXmas(grid []string) int {
	target := "XMAS"
	target_reversed := "SAMX"
	count := 0
	for i, row := range grid {
		for j := range row {
			// forward match
			if j+4 <= len(row) {
				forward := grid[i][j : j+4]
				if forward == target {
					count++
				}
			}

			// backward match
			if j-3 >= 0 {
				backward := grid[i][j-3 : j+1]
				if backward == target_reversed {
					count++
				}
			}

			//down match
			if i+4 <= len(grid) {
				down := ""
				for k := 0; k < 4; k++ {
					down += string(grid[i+k][j])
				}
				fmt.Println(down)
				if down == target {
					count++
				}
			}
			//up match
			if i-3 >= 0 {
				up := ""
				for k := 0; k < 4; k++ {
					up += string(grid[i-k][j])
				}
				if up == target {
					count++
				}
			}

			//down right
			if j+4 <= len(row) && i+4 <= len(grid) {
				down_right := ""
				for k := 0; k < 4; k++ {
					down_right += string(grid[i+k][j+k])
				}
				if down_right == target || down_right == target_reversed {
					count++
				}
			}

			//down left
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
