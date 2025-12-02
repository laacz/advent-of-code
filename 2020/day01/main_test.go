package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `1721
979
366
299
675
1456`

	expected := 514579
	result := part1(parse(input))

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := `1721
979
366
299
675
1456`

	expected := 241861950
	result := part2(parse(input))

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

}
