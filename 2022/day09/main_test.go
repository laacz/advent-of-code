package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
	got := partOne(input)
	expect := 13
	if got != expect {
		t.Errorf("partOne(%q) = %d, want %d", input, got, expect)
	}
}
func TestPartTwo(t *testing.T) {
	tests := []struct {
		Input  string
		Expect int
	}{
		{
			`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`,
			1,
		},
		{`R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`,
			36,
		},
	}
	for _, test := range tests {
		got := partTwo(test.Input)
		if got != test.Expect {
			t.Errorf("partOne(%q) = %d, want %d", test.Input, got, test.Expect)
		}
	}
}
