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
		input: `1,0,1~1,2,1
				0,0,2~2,0,2
				0,2,3~2,2,3
				0,0,4~0,2,4
				2,0,5~2,2,5
				0,1,6~2,1,6
				1,1,8~1,1,9`,
		expected1: 5,
		expected2: 7,
	},
	{
		input: `0,0,1~0,1,1
				1,1,1~1,1,1
				0,0,2~0,0,2
				0,1,2~1,1,2`,
		expected1: 3,
		expected2: 1,
	},
	{
		input: `0,0,1~1,0,1
				0,1,1~0,1,2
				0,0,5~0,0,5
				0,0,4~0,1,4`,
		expected1: 2,
		expected2: 3,
	},
	{
		input: `0,0,1~0,0,1
				1,1,1~1,1,1
				0,0,2~0,1,2
				0,1,3~1,1,3`,
		expected1: 2,
		expected2: 3,
	},
	{
		input: `5,2,265~5,5,265
		5,5,345~5,8,345
		6,3,96~6,3,98
		7,9,67~7,9,68
		3,2,274~3,3,274
		4,6,287~6,6,287
		1,3,227~1,5,227
		5,1,86~5,4,86
		9,5,184~9,6,184
		5,6,165~5,6,168`,
		expected1: 7,
		expected2: 4,
	},
}

func TestPartOne(t *testing.T) {
	for _, tt := range tests {
		actual := partOne(tt.input)
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
