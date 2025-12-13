package main

import "testing"

var input = `abc

a
b
c

ab
ac

a
a
a
a

b`

func TestPart1(t *testing.T) {
	expected := 11
	actual := part1(parse(input))

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 6
	actual := part2(parse(input))

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
