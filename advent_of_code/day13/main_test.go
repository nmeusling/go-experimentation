package main

import (
	"testing"
)

func TestExtractButton(t *testing.T) {
	input := "Button A: X+94, Y+34"
	result := extractButton(input)
	expected := button{94, 34}
	if result != expected {
		t.Errorf("Expected %v but found %v", expected, result)
	}
}

func TestPrizeButton(t *testing.T) {
	input := "Prize: X=8400, Y=5400"
	result := extractPrize(input)
	expected := button{8400, 5400}
	if result != expected {
		t.Errorf("Expected %v but found %v", expected, result)
	}
}
