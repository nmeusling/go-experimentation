package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nmeusling/go-experimentation/advent_of_code/utils"
)

func main() {
	reports := getValues(utils.ReadData("day2_sample"))
	numSafe := countSafe(reports)
	fmt.Printf("Num safe sample: %v\n", numSafe)

	reportsFinal := getValues(utils.ReadData("day2_data"))
	numSafeFinal := countSafe(reportsFinal)
	fmt.Printf("Num safe sample: %v\n", numSafeFinal)
}

func getValues(lines []string) [][]int {
	reports := make([][]int, 0)
	for i, line := range lines {
		reports = append(reports, make([]int, 0))
		levels := strings.Fields(line)
		for _, level := range levels {
			levelInt, _ := strconv.Atoi(level)
			reports[i] = append(reports[i], levelInt)
		}
	}
	return reports
}

func isSafe(levels []int) bool {
	ascending := levels[1]-levels[0] > 0
	difference := 0
	for i := 0; i < len(levels)-1; i++ {
		if ascending {
			difference = levels[i+1] - levels[i]
		} else {
			difference = levels[i] - levels[i+1]
		}
		if difference < 1 || difference > 3 {
			return false
		}
	}
	return true

}

func countSafe(reports [][]int) int {
	count := 0
	for _, report := range reports {
		if isSafe(report) {
			count++
		}
	}
	return count
}
