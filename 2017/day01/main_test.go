package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	for input, expect := range map[string]int{
		"1122":     3,
		"1111":     4,
		"1234":     0,
		"91212129": 9,
	} {
		actual := part1(input)

		if actual != expect {
			t.Errorf("Expected %d, got %d", expect, actual)
		}
	}
}

func TestPartTwo(t *testing.T) {
	for input, expect := range map[string]int{
		"1212":     6,
		"1221":     0,
		"123425":   4,
		"123123":   12,
		"12131415": 4,
	} {
		actual := part2(input)

		if actual != expect {
			t.Errorf("Expected %d, got %d", expect, actual)
		}
	}
}
