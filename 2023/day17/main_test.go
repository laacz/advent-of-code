package main

import (
	"testing"
)

var input = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

// var input = `19991
// 19111
// 11191
// 29191`

func TestPartOne(t *testing.T) {
	expected := 102
	actual := partOne(input)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 94
	actual := partTwo(input)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
