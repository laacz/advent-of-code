package main

import (
	"testing"
)

var inputs1 [][]int = [][]int{
	{4, 3, 2, 1, 3, 3},
	{4, 3, 5, 3, 9, 3},
}

func TestPartOne(t *testing.T) {
	expect := 11
	actual := part1(inputs1[0], inputs1[1])

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 31
	actual := part2(inputs1[0], inputs1[1])

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
