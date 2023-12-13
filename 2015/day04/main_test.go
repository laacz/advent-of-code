package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	for _, test := range []struct {
		input    string
		expected int
	}{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
	} {
		actual := partOne(test.input)
		if actual != test.expected {
			t.Errorf("PartOne(%s) = %d, expected %d", test.input, actual, test.expected)
		}
	}
}
