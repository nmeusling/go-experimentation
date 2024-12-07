package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/nmeusling/go-experimentation/advent_of_code/utils"
)

type calculation struct {
	result   int
	operands []int
}

func main() {

	fmt.Println("Part 1")
	calcSam := getCalculations("day7_sample")
	fmt.Println(countValid(calcSam))
	calc := getCalculations("day7_data")
	fmt.Println(countValid(calc))

}

func getCalculations(fileName string) []calculation {
	calculations := []calculation{}
	input := utils.ReadData(fileName)
	for _, line := range input {
		split := strings.Split(line, ":")
		result, _ := strconv.Atoi(split[0])
		values := []int{}
		for _, operand := range strings.Fields(split[1]) {
			value, _ := strconv.Atoi(operand)
			values = append(values, value)
		}
		calculations = append(calculations, calculation{result, values})
	}
	return calculations
}

func getPossibleCalcs(c calculation) []int {
	possibleResults := []int{c.operands[0]}
	for i := 1; i < len(c.operands); i++ {
		newResults := []int{}
		for _, r := range possibleResults {
			newResults = append(newResults, r+c.operands[i])
			newResults = append(newResults, r*c.operands[i])
			// include this for pt 2 only
			newResults = append(newResults, combineDigits(r, c.operands[i]))
		}
		possibleResults = newResults
	}
	return possibleResults
}

func combineDigits(first, second int) int {
	val, _ := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(second))
	return val
}

func countValid(calcs []calculation) int {
	sum := 0
	for _, c := range calcs {
		possibleResults := getPossibleCalcs(c)
		if slices.Contains(possibleResults, c.result) {
			sum += c.result
		}
	}
	return sum
}
