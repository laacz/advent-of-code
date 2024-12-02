package main

import (
	"testing"
)

var inputs1 [][]int = [][]int{
	{7, 6, 4, 2, 1},
	{1, 2, 7, 8, 9},
	{9, 7, 6, 2, 1},
	{1, 3, 2, 4, 5},
	{8, 6, 4, 4, 1},
	{1, 3, 6, 7, 9},
}

func TestPartOne(t *testing.T) {
	expect := 2
	actual := part1(inputs1)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 4
	actual := part2(inputs1)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
