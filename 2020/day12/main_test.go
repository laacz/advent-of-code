package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `F10
N3
F7
R90
F11`

	expected := 25
	result := part1(input)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := `F10
N3
F7
R90
F11`

	expected := 286
	result := part2(input)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
