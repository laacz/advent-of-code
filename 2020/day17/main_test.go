package main

import "testing"

var example = `.#.
..#
###`

func TestPart1(t *testing.T) {
	expected := 112
	result := part1(example)
	if result != expected {
		t.Errorf("part1() = %d, expected %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 848
	result := part2(example)
	if result != expected {
		t.Errorf("part2() = %d, expected %d", result, expected)
	}
}
