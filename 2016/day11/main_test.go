package main

import "testing"

func TestPartOne(t *testing.T) {
	expect := 11
	actual := partOne(State{
		floor: 1,
		combos: Combos{
			"H": {2, 1}, // H
			"L": {3, 1}, // Li
		}})

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	// Part 2 uses the same algorithm, just with more elements
	// Testing with 3 pairs starting on floor 1
	expect := 27 // 3 pairs all on floor 1 takes 27 steps
	actual := partTwo(State{
		floor: 1,
		combos: Combos{
			"A": {1, 1},
			"B": {1, 1},
			"C": {1, 1},
		},
	})

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
