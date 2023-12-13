package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	var inputs = map[string]int{
		"(())":    0,
		"()()":    0,
		"(((":     3,
		"(()(()(": 3,
		"))(((((": 3,
		"())":     -1,
		"))(":     -1,
		")))":     -3,
		")())())": -3,
	}

	for input, expected := range inputs {
		actual := partOne(input)
		if actual != expected {
			t.Errorf("PartOne(%s) = %d, expected %d", input, actual, expected)
		}
	}
}
func TestPartwo(t *testing.T) {
	var inputs = map[string]int{
		")":     1,
		"()())": 5,
	}

	for input, expected := range inputs {
		actual := partTwo(input)
		if actual != expected {
			t.Errorf("PartTwo(%s) = %d, expected %d", input, actual, expected)
		}
	}
}
