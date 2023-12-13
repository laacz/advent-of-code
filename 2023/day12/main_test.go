package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	var inputs = []struct {
		input  string
		expect int
	}{
		{"???.### 1,1,3", 1},
		{".??..??...?##. 1,1,3", 4},
		{"?#?#?#?#?#?#?#? 1,3,1,6", 1},
		{"????.#...#... 4,1,1", 1},
		{"????.######..#####. 1,6,5", 4},
		{"?###???????? 3,2,1", 10},
	}

	for _, input := range inputs {
		actual := partOne(input.input)

		if actual != input.expect {
			t.Errorf("Expected %d, got %d", input.expect, actual)
		}
	}
}

func TestPartTwo(t *testing.T) {
	var inputs = []struct {
		input  string
		expect int
	}{
		{"???.### 1,1,3", 1},
		{".??..??...?##. 1,1,3", 16384},
		{"?#?#?#?#?#?#?#? 1,3,1,6", 1},
		{"????.#...#... 4,1,1", 16},
		{"????.######..#####. 1,6,5", 2500},
		{"?###???????? 3,2,1", 506250},
	}

	for _, input := range inputs {
		actual := partTwo(input.input)

		if actual != input.expect {
			t.Errorf("Expected %d, got %d", input.expect, actual)
		}
	}
}
