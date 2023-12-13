package main

import (
	"testing"
)

var input = `y -> x
12 -> b
65000 -> y
y AND b -> z
NOT z -> a`

func TestPartOne(t *testing.T) {
	expected := 65527
	actual := partOne(input)
	if actual != expected {
		t.Errorf("PartTwo(%s) = %v, expected %v", input, actual, expected)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 543
	actual := partTwo(input, partOne(input))
	if actual != expected {
		t.Errorf("PartTwo(%s) = %v, expected %v", input, actual, expected)
	}
}
