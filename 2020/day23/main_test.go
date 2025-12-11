package main

import "testing"

var cups = parse("389125467")

func TestPart1(t *testing.T) {
	expected := "67384529"
	actual := part1(cups)

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 149245887792
	actual := part2(cups)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
