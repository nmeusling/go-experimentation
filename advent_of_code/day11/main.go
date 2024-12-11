package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nmeusling/go-experimentation/advent_of_code/utils"
)

func main() {
	sampleStones := getStones("day11_sample")
	sampleStones = stepStoneNTimes(25, sampleStones)
	fmt.Println(len(sampleStones))

	stones := getStones("day11_data")
	stones = stepStoneNTimes(25, stones)
	fmt.Println(len(stones))

	stones2 := getStones("day11_data")
	stones2 = stepStoneNTimes(50, stones2)
	fmt.Println(len(stones2))
}

func getStones(fileName string) []int {
	stones := []int{}
	lines := utils.ReadData(fileName)
	for _, field := range strings.Fields(lines[0]) {
		val, _ := strconv.Atoi(field)
		stones = append(stones, val)
	}
	return stones
}

func stepStoneNTimes(n int, stones []int) []int {
	newStones := stones
	for i := 0; i < n; i++ {
		newStones = stepStone(newStones)
	}
	return newStones
}

func stepStone(stones []int) []int {
	newStones := []int{}
	for _, stone := range stones {
		if stone == 0 {
			newStones = append(newStones, 1)
			continue
		}

		if len(strconv.Itoa(stone))%2 == 0 {
			split := splitStone(stone)
			newStones = append(newStones, split...)
			continue
		}
		newStones = append(newStones, stone*2024)
	}
	return newStones
}

func splitStone(stone int) []int {
	digits := strconv.Itoa(stone)
	mid := len(digits) / 2
	first, _ := strconv.Atoi(digits[:mid])
	second, _ := strconv.Atoi(digits[mid:])
	return []int{first, second}
}
