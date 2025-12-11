package main

import "testing"

var input = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10
`

var p1, p2 = parse(input)

func TestPart1(t *testing.T) {
	expected := 306
	actual := part1(p1, p2)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 291
	actual := part2(p1, p2)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
