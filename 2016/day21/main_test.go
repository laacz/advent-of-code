package main

import "testing"

func TestApply(t *testing.T) {
	input := "abcdefgh"
	tests := []struct {
		name string
		i    Instruction
		want string
	}{
		{"swap position", SwapPosition{4, 0}, "ebcdafgh"},
		{"swap letter", SwapLetter{'d', 'b'}, "adcbefgh"},
		{"rotate left 1", RotateLeft{1}, "bcdefgha"},
		{"rotate right 1", RotateRight{1}, "habcdefg"},
		{"rotate left 2", RotateLeft{2}, "cdefghab"},
		{"rotate right 2", RotateRight{2}, "ghabcdef"},
		{"rotate based on position of letter b", RotateBased{"b"}, "ghabcdef"},
		{"rotate based on position of letter d", RotateBased{"d"}, "efghabcd"},
		{"rotate based on position of letter e", RotateBased{"e"}, "cdefghab"},
		{"reverse positions", Reverse{0, 4}, "edcbafgh"},
		{"reverse positions", Reverse{4, 7}, "abcdhgfe"},
		{"move position", Move{1, 4}, "acdebfgh"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Apply(input); got != tt.want {
				t.Errorf("Apply(%s) = %v, want %v", input, got, tt.want)
			}
		})
	}
}

func TestUnapply(t *testing.T) {
	input := "abcdefgh"
	tests := []struct {
		name string
		i    Instruction
		want string
	}{
		{"swap position", SwapPosition{4, 0}, "ebcdafgh"},
		{"swap letter", SwapLetter{'d', 'b'}, "adcbefgh"},
		{"rotate left 1", RotateLeft{1}, "bcdefgha"},
		{"rotate right 1", RotateRight{1}, "habcdefg"},
		{"rotate left 2", RotateLeft{2}, "cdefghab"},
		{"rotate right 2", RotateRight{2}, "ghabcdef"},
		{"rotate based on position of letter b", RotateBased{"b"}, "ghabcdef"},
		{"rotate based on position of letter d", RotateBased{"d"}, "efghabcd"},
		{"rotate based on position of letter e", RotateBased{"e"}, "cdefghab"},
		{"reverse positions", Reverse{0, 4}, "edcbafgh"},
		{"reverse positions", Reverse{4, 7}, "abcdhgfe"},
		{"move position", Move{1, 4}, "acdebfgh"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Unapply(tt.want); got != input {
				t.Errorf("Apply(%s) = %v, want %v", tt.want, got, input)
			}
		})
	}
}

func TestPartOne(t *testing.T) {
	tests := []struct {
		name      string
		password  string
		scrambled string
	}{
		{`abcde
		decab
			swap position 4 with position 0
			swap letter d with letter b
			reverse positions 0 through 4
			rotate left 1 step
			move position 1 to position 4
			move position 3 to position 0
			rotate based on position of letter b
			rotate based on position of letter d
			`,
			"abcde",
			"decab",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partOne(tt.name); got != tt.scrambled {
				t.Errorf("partOne() = %v, want %v", got, tt.scrambled)
			}
		})
	}
}
func TestPartTwo(t *testing.T) {
	tests := []struct {
		name      string
		password  string
		scrambled string
	}{
		{`abcde
		decab
			swap position 4 with position 0
			swap letter d with letter b
			reverse positions 0 through 4
			rotate left 1 step
			move position 1 to position 4
			move position 3 to position 0
			rotate based on position of letter b
			rotate based on position of letter d
			`,
			"abcde",
			"decab",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partTwo(tt.name); got != tt.password {
				t.Errorf("partTwo() = %v, want %v", got, tt.password)
			}
		})
	}
}
