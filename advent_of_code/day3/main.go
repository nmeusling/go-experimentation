package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/nmeusling/go-experimentation/advent_of_code/utils"
)

func main() {
	sample := strings.Join(utils.ReadData("day3_sample"), "")
	fmt.Printf("Sample Total = %v\n", calculate(getMatches(sample)))
	data := strings.Join(utils.ReadData("day3_data"), "")
	fmt.Printf("Total = %v\n", calculate(getMatches(data)))
	sample2 := strings.Join(utils.ReadData("day3_sample2"), "")
	fmt.Printf("Sample Total = %v\n", calculate(getMatches(removeDisabled(sample2))))
	data = strings.Join(utils.ReadData("day3_data"), "")
	fmt.Printf("Total = %v\n", calculate(getMatches(removeDisabled(data))))
}

func removeDisabled(data string) string {
	re := regexp.MustCompile(`don't[(][)].*?do[(][)]`)
	return re.ReplaceAllString(data, "")
}

func getMatches(data string) []string {
	re := regexp.MustCompile(`mul[(]\d+,\d+[)]`)
	return re.FindAllString(data, -1)
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
