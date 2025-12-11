package main

import "testing"

var input = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`

func TestPart1(t *testing.T) {
	expected := 2
	actual := part1(parse(input))

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 1
	actual := part2(parse(input))

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
