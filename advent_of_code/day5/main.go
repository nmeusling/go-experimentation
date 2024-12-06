package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/nmeusling/go-experimentation/advent_of_code/utils"
)

func getInput(fileName string) (map[int][]int, [][]int) {
	lines := utils.ReadData(fileName)
	rules := make(map[int][]int)
	manuals := make([][]int, 0)
	firstSection := true
	for _, line := range lines {
		line := strings.TrimSpace(line)
		if line == "" {
			firstSection = false
			continue
		}
		if firstSection {
			rule := strings.Split(line, "|")
			before, _ := strconv.Atoi(rule[0])
			after, _ := strconv.Atoi(rule[1])
			rules[before] = append(rules[before], after)
			continue
		}
		manual := []int{}
		for _, entry := range strings.Split(line, ",") {
			value, _ := strconv.Atoi(entry)
			manual = append(manual, value)
		}
		manuals = append(manuals, manual)
	}
	return rules, manuals

}

func verifyOrder(rules map[int][]int, line []int) bool {
	for i := len(line) - 1; i >= 0; i-- {
		afterValues := rules[line[i]]
		beforeValues := line[:i]
		for _, value := range afterValues {
			if slices.Contains(beforeValues, value) {
				return false
			}
		}
	}
	return true
}

func sumInvalid(rules map[int][]int, lines [][]int) int {
	count := 0
	for _, line := range lines {
		if verifyOrder(rules, line) {
			median := line[len(line)/2]
			count += median
		}
	}
	return count
}

// func reOrder(rules map[int][]int, line []int){
// 	correct := []int{line[0]}
// 	for i := 1; i < len(line); i++{
// 		if
// 	}
// }

func main() {
	rulesSam, manualsSam := getInput("day5_sample")
	fmt.Printf("Sample sum: %v\n", sumInvalid(rulesSam, manualsSam))
	rules, manuals := getInput("day5_data")
	fmt.Printf("Sum: %v", sumInvalid(rules, manuals))
}
