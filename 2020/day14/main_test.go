package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
	mem[8] = 11
	mem[7] = 101
	mem[8] = 0`

	expected := 165
	result := part1(parse(input))

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := `mask = 000000000000000000000000000000X1001X
	mem[42] = 100
	mask = 00000000000000000000000000000000X0XX
	mem[26] = 1`

	expected := 208
	result := part2(parse(input))

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
