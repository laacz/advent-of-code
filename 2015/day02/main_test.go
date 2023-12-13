package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	var input = `2x3x4
	1x1x10`
	expected := 58 + 43

	actual := partOne(input)
	if actual != expected {
		t.Errorf("PartOne(%s) = %d, expected %d", input, actual, expected)
	}
}
func TestPartwo(t *testing.T) {
	var input = `2x3x4
	1x1x10`
	expected := 34 + 14

	actual := partTwo(input)
	if actual != expected {
		t.Errorf("PartTwo(%s) = %d, expected %d", input, actual, expected)
	}
}
