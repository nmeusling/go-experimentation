package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/nmeusling/go-experimentation/advent_of_code/utils"
)

func main() {
	machines := getButtonInput("day13_sample")
	fmt.Println("sample pt 1", calcTokensForWins(machines))
	machines2 := getButtonInput("day13_data")
	fmt.Println("data pt 1", calcTokensForWins(machines2))
}

func getButtonInput(fileName string) []machine {
	machines := []machine{}
	lines := utils.ReadData(fileName)
	for i := 0; i+3 <= len(lines); i += 4 {
		x := extractButton(lines[i])
		y := extractButton(lines[i+1])
		prize := extractPrize(lines[i+2])
		machines = append(machines, machine{x, y, prize})
	}
	return machines
}

func extractButton(line string) button {
	end := strings.Split(line, "X+")[1]
	sections := strings.Split(end, ", Y+")
	x, _ := strconv.Atoi(sections[0])
	y, _ := strconv.Atoi(sections[1])
	return button{x, y}
}

func extractPrize(line string) button {
	end := strings.Split(line, "X=")[1]
	sections := strings.Split(end, ", Y=")
	x, _ := strconv.Atoi(sections[0])
	y, _ := strconv.Atoi(sections[1])
	return button{x, y}

}

func calcTokensForWins(machines []machine) int {
	total := 0
	for _, m := range machines {
		total += m.calcTokensForWin()
	}
	return total
}

func (m machine) calcTokensForWin() int {
	possibleCosts := []int{}
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			if m.a.x*a+m.b.x*b == m.prize.x && m.a.y*a+m.b.y*b == m.prize.y {
				possibleCosts = append(possibleCosts, cost(a, b))
			}
		}
	}
	if len(possibleCosts) > 0 {
		return slices.Min(possibleCosts)
	}
	return 0
}

func cost(a, b int) int {
	return 3*a + b
}

type button struct {
	x, y int
}

type machine struct {
	a, b, prize button
}
