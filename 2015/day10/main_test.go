package main

import (
	"testing"
)

var cases = []struct {
	input    string
	expected string
}{
	{
		input:    "1",
		expected: "11",
	}, {
		input:    "11",
		expected: "21",
	}, {
		input:    "211",
		expected: "1221",
	},
	{
		input:    "21",
		expected: "1211",
	}, {
		input:    "1211",
		expected: "111221",
	}, {
		input:    "111221",
		expected: "312211",
	},
}

func TestPartOne(t *testing.T) {
	for _, c := range cases {
		actual := lookAndSay(c.input)
		if actual != c.expected {
			t.Errorf("PartOne(%v) = %v, expected %v", c.input, actual, c.expected)
		}
	}
}

// func TestPartTwo(t *testing.T) {
// 	expected := 982
// 	actual := partTwo(input)
// 	if actual != expected {
// 		t.Errorf("PartTwo(%s) = %v, expected %v", input, actual, expected)
// 	}
// }
