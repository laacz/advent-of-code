package main

import (
	"testing"
)

var tests = []struct {
	input     string
	expected1 int
	expected2 int
}{
	{
		input: `19, 13, 30 @ -2,  1, -2
		18, 19, 22 @ -1, -1, -2
		20, 25, 34 @ -2, -2, -4
		12, 31, 28 @ -1, -2, -1
		20, 19, 15 @  1, -5, -3`,
		expected1: 2,
		expected2: 47,
	},
}

func TestPartOne(t *testing.T) {
	for _, tt := range tests {
		actual := partOne(tt.input, 7, 27)
		if actual != tt.expected1 {
			t.Errorf("Expected %d, got %d", tt.expected1, actual)
		}
	}
}

func TestPartTwo(t *testing.T) {
	for _, tt := range tests {
		actual := partTwo(tt.input)
		if actual != tt.expected2 {
			t.Errorf("Expected %d, got %d", tt.expected2, actual)
		}
	}
}
