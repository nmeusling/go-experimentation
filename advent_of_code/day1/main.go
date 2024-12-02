package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sampleLines := readData("day1_sample")
	leftSample, rightSample := getValues(sampleLines)
	sampleDistance := calcDistance(leftSample, rightSample)
	sampleSimilarity := calcSimilarity(leftSample, rightSample)
	fmt.Printf("Total Distance (Sample) %v\nTotal Similarity (sample) %v\n", sampleDistance, sampleSimilarity)

	lines := readData("day1_data")
	left, right := getValues(lines)
	totalDistance := calcDistance(left, right)
	totalSimilarity := calcSimilarity(left, right)
	fmt.Printf("Total Distance (Actual) %v\nTotal Similarity (actual) %v\n", totalDistance, totalSimilarity)

}

func calcSimilarity(left, right []int) int {
	totalSimilarity := 0
	for _, value := range left {
		totalSimilarity += value * countOccurrences(value, right)
	}
	return totalSimilarity
}

func countOccurrences(target int, values []int) int {
	count := 0
	for _, value := range values {
		if value == target {
			count++
		}
	}
	return count
}

func calcDistance(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)
	totalDistance := 0
	for i, l := range left {
		distance := math.Abs(float64(l) - float64(right[i]))
		totalDistance += int(distance)
	}
	return totalDistance

}

func getValues(lines []string) ([]int, []int) {
	left := make([]int, 0)
	right := make([]int, 0)
	for _, line := range lines {
		fields := strings.Fields(line)
		leftField, _ := strconv.Atoi(fields[0])
		rightField, _ := strconv.Atoi(fields[1])
		left = append(left, leftField)
		right = append(right, rightField)
	}
	return left, right
}

func readData(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}
