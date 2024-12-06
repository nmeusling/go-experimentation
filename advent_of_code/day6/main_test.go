package main

import (
	"testing"
)

func TestNotClosedLoop(t *testing.T) {
	var g = guard{location{6, 4}, up, location{6, 4}, up, 0}
	grid := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}
	result := isClosedLoop(grid, g)
	expected := false
	if result != expected {
		t.Errorf("Expected %v, but found %v", expected, result)
	}
}

func TestClosedLoopTwo(t *testing.T) {
	var g = guard{location{6, 4}, up, location{6, 4}, up, 0}
	grid := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"......#.#.",
		"#.........",
		"......#...",
	}
	result := isClosedLoop(grid, g)
	expected := true
	if result != expected {
		t.Errorf("Expected %v, but found %v", expected, result)
	}
}

func TestClosedLoopThree(t *testing.T) {
	var g = guard{location{6, 4}, up, location{6, 4}, up, 0}
	grid := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		".......##.",
		"#.........",
		"......#...",
	}
	result := isClosedLoop(grid, g)
	expected := true
	if result != expected {
		t.Errorf("Expected %v, but found %v", expected, result)
	}
}

func TestClosedLoopFour(t *testing.T) {
	var g = guard{location{6, 4}, up, location{6, 4}, up, 0}
	grid := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"##........",
		"......#...",
	}
	result := isClosedLoop(grid, g)
	expected := true
	if result != expected {
		t.Errorf("Expected %v, but found %v", expected, result)
	}
}

func TestClosedLoopFive(t *testing.T) {
	var g = guard{location{6, 4}, up, location{6, 4}, up, 0}
	grid := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#..#......",
		"......#...",
	}
	result := isClosedLoop(grid, g)
	expected := true
	if result != expected {
		t.Errorf("Expected %v, but found %v", expected, result)
	}
}

func TestClosedLoopSix(t *testing.T) {
	var g = guard{location{6, 4}, up, location{6, 4}, up, 0}
	grid := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......##..",
	}
	result := isClosedLoop(grid, g)
	expected := true
	if result != expected {
		t.Errorf("Expected %v, but found %v", expected, result)
	}
}

func TestClosedLoopOne(t *testing.T) {
	var g = guard{location{6, 4}, up, location{6, 4}, up, 0}
	grid := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#.#^.....",
		"........#.",
		"#.........",
		"......#...",
	}
	result := isClosedLoop(grid, g)
	expected := true
	if result != expected {
		t.Errorf("Expected %v, but found %v", expected, result)
	}
}
