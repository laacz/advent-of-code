package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	input := `aa bb cc dd ee
aa bb cc dd aa
aa bb cc dd aaa
`
	if actual := part1(parseInput(input)); actual != 2 {
		t.Errorf("Expected 2, got %d", actual)
	}
}

func TestPartTwo(t *testing.T) {
	input := `abcde fghij
abcde xyz ecdab
a ab abc abd abf abj
iiii oiii ooii oooi oooo
oiii ioii iioi iiio`

	if actual := part2(parseInput(input)); actual != 3 {
		t.Errorf("Expected 3, got %d", actual)
	}
}
