package main

import (
	"testing"
)

var input = `029A
980A
179A
456A
379A`

func TestPartOne(t *testing.T) {
	expected := 126384
	if actual := part1(parseInput(input)); actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}
