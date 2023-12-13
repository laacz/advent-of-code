package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	for _, test := range []struct {
		input    string
		expected int
	}{
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
	} {
		actual := partOne(test.input)
		if actual != test.expected {
			t.Errorf("PartOne(%s) = %d, expected %d", test.input, actual, test.expected)
		}
	}
}

func TestPartwo(t *testing.T) {
	for _, test := range []struct {
		input    string
		expected int
	}{
		{">v", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
	} {
		actual := partTwo(test.input)
		if actual != test.expected {
			t.Errorf("PartTwo(%s) = %d, expected %d", test.input, actual, test.expected)
		}
	}
}
