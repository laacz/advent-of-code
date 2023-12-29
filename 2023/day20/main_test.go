package main

import (
	"testing"
)

var cases = []struct {
	input     string
	expected1 int
}{
	{
		`broadcaster -> a, b, c
		%a -> b
		%b -> c
		%c -> inv
		&inv -> a`,
		32000000,
	},
	{
		`broadcaster -> a
		%a -> inv, con
		&inv -> b
		%b -> con
		&con -> output`,
		11687500,
	},
}

func TestPartOne(t *testing.T) {
	for _, c := range cases {
		actual := partOne(c.input)

		if actual != c.expected1 {
			t.Errorf("Expected %d, got %d", c.expected1, actual)
		}
	}
}
