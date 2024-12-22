package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	input := `1
10
100
2024`
	expected := 37327623
	if actual := part1(parseInput(input)); actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	input := `1
2
3
2024`
	// input = "123"
	expected := 23
	if actual := part2(parseInput(input)); actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}
