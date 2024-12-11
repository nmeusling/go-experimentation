package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nmeusling/go-experimentation/advent_of_code/utils"
)

func main() {
	sampleStones := getStones("day11_sample")
	sampleStones = stepStonesNTimes(25, sampleStones)
	fmt.Println(len(sampleStones))

	stones := getStones("day11_data")
	stones = stepStonesNTimes(25, stones)
	fmt.Println(len(stones))

	sampleStones2 := getInitialStones("day11_sample")
	processed := stepMappedStonesNTimes(sampleStones2, 75)
	fmt.Println(countStones(processed))

	stones2 := getInitialStones("day11_data")
	fmt.Println(countStones(stepMappedStonesNTimes(stones2, 75)))
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

func stepStonesNTimes(n int, stones []int) []int {
	newStones := stones
	for i := 0; i < n; i++ {
		newStones = stepStones(newStones)
	}
	return newStones
}

func stepStones(stones []int) []int {
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

func getInitialStones(fileName string) map[int]int {
	stones := make(map[int]int)
	for _, s := range getStones(fileName) {
		stones[s]++
	}
	return stones
}

func stepMappedStones(stones map[int]int) map[int]int {
	newStones := make(map[int]int)
	for s, count := range stones {
		results := stepStones([]int{s})
		for _, r := range results {
			newStones[r] += count
		}
	}
	return newStones
}

func stepMappedStonesNTimes(stones map[int]int, n int) map[int]int {
	newStones := stones
	for i := 0; i < n; i++ {
		newStones = stepMappedStones(newStones)
	}
	return newStones
}

func countStones(stones map[int]int) int {
	numRocks := 0
	for _, count := range stones {
		numRocks += count
	}
	return numRocks
}

// 0->1->2024->20, 24->2, 0, 2,4 ->
// 1->2024->20, 24->2, 0, 2,4 ->
//2->4048->40, 48-> 4, 0, 4, 8
//3 -> 6072 -> 60, 72 -> 6, 0, 7, 2
//4 -> 8096 -> 80, 96 -> 8,0,9,6
//5 -> 10120 - 20482880 -> 2048, 2880 -> 20, 48, 28, 80 ->2,0,4,8,2,8,8,0
//6 -> 12144 -> 24579456 -> 2457, 9456 -> 24, 57, 94, 56 -> 2,4,5,7,9,4,5,6
//7 -> 14168 -> 28676032 -> 2867, 6032 -> 28, 67, 60, 32 -> 2, 8, 6, 7, 6, 0, 3, 2
//8 -> 16192 -> 32772608 -> 3277, 2608 -> 32, 77, 26, 8 -> 3, 2, 7, 7, 2, 6, 16192
//9 -> 18216 -> 36869184 -> 3686, 9184 -> 36, 86, 91, 84 -> 3,6,8,6,9,1,8,4
