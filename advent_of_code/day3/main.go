package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/nmeusling/go-experimentation/advent_of_code/utils"
)

func main() {
	sample := utils.ReadData("day3_sample")
	fmt.Printf("Sample Total = %v\n", calculate(getMatches(sample)))
	data := utils.ReadData("day3_data")
	fmt.Printf("Total = %v\n", calculate(getMatches(data)))
}

func getMatches(fileLines []string) []string {
	re := regexp.MustCompile(`mul[(]\d+,\d+[)]`)
	matches := make([]string, 0)
	for _, line := range fileLines {
		matches = append(matches, re.FindAllString(line, -1)...)
	}
	return matches
}

func calculate(mulInstructions []string) int {
	total := 0
	for _, value := range mulInstructions {
		re := regexp.MustCompile(`\d+`)
		values := re.FindAllString(value, 2)
		left, _ := strconv.Atoi(values[0])
		right, _ := strconv.Atoi(values[1])
		total += left * right
	}
	return total
}
