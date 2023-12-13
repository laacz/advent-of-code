package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	for _, c := range []struct {
		input    string
		expected int
	}{
		{
			input:    `[1,2,3]`,
			expected: 6,
		}, {
			input:    `{"a":2,"b":4}`,
			expected: 6,
		}, {
			input:    `[[[3]]]`,
			expected: 3,
		},
		{
			input:    `{"a":{"b":4},"c":-1}`,
			expected: 3,
		}, {
			input:    `{"a":[-1,1]}`,
			expected: 0,
		}, {
			input:    `[-1,{"a":1}]`,
			expected: 0,
		}, {
			input:    `[]`,
			expected: 0,
		}, {
			input:    `{}`,
			expected: 0,
		}, {
			input:    `[1,{"c":"red","b":2},3]`,
			expected: 6,
		}, {
			input:    `{"d":"1",[2]]`,
			expected: 2,
		},
	} {
		actual := partOne(c.input)
		if actual != c.expected {
			t.Errorf("PartOne(%v) = %v, expected %v", c.input, actual, c.expected)
		}
	}
}

func TestPartTwo(t *testing.T) {
	for _, c := range []struct {
		input    string
		expected int
	}{
		{
			input:    `[1,2,3]`,
			expected: 6,
		}, {
			input:    `{"a":2,"b":4}`,
			expected: 6,
		}, {
			input:    `[[[3]]]`,
			expected: 3,
		}, {
			input:    `{"a":{"b":4},"c":-1}`,
			expected: 3,
		}, {
			input:    `{"a":[-1,1]}`,
			expected: 0,
		}, {
			input:    `[-1,{"a":1}]`,
			expected: 0,
		}, {
			input:    `[]`,
			expected: 0,
		}, {
			input:    `{}`,
			expected: 0,
		}, {
			input:    `[1,{"c":"red","b":2},3]`,
			expected: 4,
		}, {
			input:    `{"d":"1","":[2]}`,
			expected: 2,
		}, {
			input:    `[1,"red",5]`,
			expected: 6,
		}, {
			input:    `{"d":"red","e":[1,2,3,4],"f":5}`,
			expected: 0,
		},
	} {
		actual := partTwo(c.input)
		if actual != c.expected {
			t.Errorf("PartOne(%v) = %v, expected %v", c.input, actual, c.expected)
		}
	}
}
