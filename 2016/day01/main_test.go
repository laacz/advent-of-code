package main

import "testing"

func TestPartOne(t *testing.T) {
	var cases = []struct {
		input  string
		output int
	}{
		{
			input:  `R2, L3`,
			output: 5,
		},
		{
			input:  `R2, R2, R2`,
			output: 2,
		},
		{
			input:  `R5, L5, R5, R3`,
			output: 12,
		},
	}

	for _, c := range cases {
		input := parse(c.input)
		output := partOne(input)
		if output != c.output {
			t.Errorf("Expected %d, got %d", c.output, output)
		}
	}
}

func TestPartTwo(t *testing.T) {
	if partTwo(parse(`R8, R4, R4, R8`)) != 4 {
		t.Error("Expected 4")
	}
}
