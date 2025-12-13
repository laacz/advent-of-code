package main

import "testing"

var input = `16
10
15
5
1
11
7
19
6
12
4`

func TestPart1(t *testing.T) {
	expected := 7 * 5
	actual := part1(parse(input))

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 8
	actual := part2(parse(input))

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
