package main

import (
	"strings"
	"testing"
)

var input = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func TestPartOne(t *testing.T) {
	expect := 3749
	actual := part1(strings.TrimSpace(input))

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 11387
	actual := part2(strings.TrimSpace(input))

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
