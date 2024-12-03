package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	input := "5 1 9 5\n7 5 3\n2 4 6 8"
	expect := 18
	if actual := part1(input); actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	input := "5 9 2 8\n9 4 7 3\n3 8 6 5"
	expect := 9
	if actual := part2(input); actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}

}
