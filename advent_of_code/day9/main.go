package main

import (
	"fmt"
	"strconv"

	"github.com/nmeusling/go-experimentation/advent_of_code/utils"
)

var empty = -1

func main() {
	discMap := utils.ReadData("day9_sample")[0]
	formattedDisc := rearrangeBlocks(makeBlocks(discMap))
	fmt.Printf("(sample) Pt 1 Disc: %v\nChecksum: %v\n", formattedDisc, calcChecksum(formattedDisc))
	formatTwo := rearrangeFiles(makeBlocks(discMap))
	fmt.Printf("(sample) Pt 2 Disc: %v\nChecksum:%v\n", formatTwo, calcChecksum(formatTwo))
	discMap = utils.ReadData("day9_data")[0]
	fmt.Printf("Pt 1 checksum: %v\n", calcChecksum(rearrangeBlocks(makeBlocks(discMap))))
	fmt.Printf("Pt 2 checksum: %v\n", calcChecksum(rearrangeFiles(makeBlocks(discMap))))

}

func makeBlocks(discMap string) []int {
	isEmpty := false
	id := 0
	blocks := []int{}
	for _, char := range discMap {
		num, _ := strconv.Atoi(string(char))
		for j := 0; j < num; j++ {
			if !isEmpty {
				blocks = append(blocks, id)
			}
			if isEmpty {
				blocks = append(blocks, empty)
			}
		}
		if !isEmpty {
			id++
		}
		isEmpty = !isEmpty
	}
	return blocks
}

func rearrangeBlocks(blocks []int) []int {
	i := 0
	j := len(blocks) - 1

	for i < j {
		if blocks[i] != empty {
			i++
			continue
		}
		if blocks[j] == empty {
			j--
			continue
		}
		blocks[i] = blocks[j]
		blocks[j] = empty
		i++
		j--
	}
	return blocks
}

// 00...111...2...333.44.5555.6666.777.888899
// 00992111777.44.333....5555.6666.....8888..
func rearrangeFiles(blocks []int) []int {
	currentIndex := len(blocks) - 1
	finalID := blocks[currentIndex]
	for currentID := finalID; currentID > 0; currentID-- {
		for blocks[currentIndex] == empty || blocks[currentIndex] > currentID {
			currentIndex--
		}
		start := currentIndex
		end := currentIndex
		for ; blocks[currentIndex] == currentID; currentIndex-- {
			start = currentIndex
		}
		blockLength := end - start + 1
		newStart, newEnd := findFirstGap(blockLength, blocks)
		if newStart >= 0 && newStart < start && newEnd >= 0 && newEnd < end {
			for i := newStart; i <= newEnd; i++ {
				blocks[i] = currentID
			}
			for i := start; i <= end; i++ {
				blocks[i] = empty
			}
		}

	}
	return blocks
}

func findFirstGap(size int, blocks []int) (int, int) {
	start, end := 0, 0
	for i := 0; i < len(blocks); i++ {
		if blocks[i] != empty {
			continue
		}
		start = i
		for j := i; j < len(blocks) && blocks[j] == empty; j++ {
			end = j
		}
		if (end - start + 1) >= size {
			return start, start + size - 1
		}
	}
	return -1, -1
}

func calcChecksum(blocks []int) int {
	sum := 0
	for i, val := range blocks {
		if val != empty {
			sum += i * val
		}
	}
	return sum
}
