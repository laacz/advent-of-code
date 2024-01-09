package main

import "testing"

func TestPartOne(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"ihgpwlah", "DDRRRD"},
		{"kglvqrro", "DDUDRLRRUDRD"},
		{"ulqzkmiv", "DRURDRUDDLLDLUURRDULRLDUUDDDRR"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partOne(tt.name); got != tt.want {
				t.Errorf("partOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestPartTwo(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"ihgpwlah", 370},
		{"kglvqrro", 492},
		{"ulqzkmiv", 830},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partTwo(tt.name); got != tt.want {
				t.Errorf("partTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
