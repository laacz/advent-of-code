package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `939
7,13,x,x,59,x,31,19`

	expected := 295
	result := part1(input)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	tests := []struct{
		expected int
		input string
	}{
		{1068781,`7,13,x,x,59,x,31,19`},
		{754018, `67,7,59,61`},
		{779210, `67,x,7,59,61`},
		{1261476, `67,7,x,59,61`},
		{1202161486, `1789,37,47,1889`},
	}

	for _, test := range tests {
		result := part2("42\n"+test.input)

		if result != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, result)
		}
	}
}
