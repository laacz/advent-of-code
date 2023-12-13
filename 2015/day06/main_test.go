package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	for _, test := range []struct {
		input    string
		expected []Instruction
	}{
		{
			`turn on 0,0 through 999,999`,
			[]Instruction{
				{"on", Coords{0, 0}, Coords{999, 999}},
			},
		},
		{
			`toggle 0,0 through 999,0`,
			[]Instruction{
				{"toggle", Coords{0, 0}, Coords{999, 0}},
			},
		},

		{
			`turn off 499,499 through 500,500`,
			[]Instruction{
				{"off", Coords{499, 499}, Coords{500, 500}},
			},
		},
	} {
		actual := parse(test.input)
		if len(actual) != len(test.expected) {
			t.Errorf("parse(%s) = %v, expected %v", test.input, actual, test.expected)
		}
		for i, v := range actual {
			if v != test.expected[i] {
				t.Errorf("parse(%s) = %v, expected %v", test.input, actual, test.expected)
			}
		}
	}
}

func TestPartOne(t *testing.T) {
	for _, test := range []struct {
		input    string
		expected int
	}{
		{
			`turn on 0,0 through 999,999
			toggle 0,0 through 999,0
			turn off 499,499 through 500,500`,
			998996,
		},
	} {
		actual := partOne(test.input)
		if actual != test.expected {
			t.Errorf("PartOne(%s) = %v, expected %v", test.input, actual, test.expected)
		}
	}
}

func TestPartTwo(t *testing.T) {
	for _, test := range []struct {
		input    string
		expected int
	}{
		{
			`turn on 0,0 through 0,0
			toggle 0,0 through 999,999`,
			2000001,
		},
	} {
		actual := partTwo(test.input)
		if actual != test.expected {
			t.Errorf("PartTwo(%s) = %v, expected %v", test.input, actual, test.expected)
		}
	}
}
