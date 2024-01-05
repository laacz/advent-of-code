package main

import "testing"

func TestFillData(t *testing.T) {
	tests := []struct {
		name   string
		want   string
		length int
	}{
		{`1`, `100`, 3},
		{`0`, `001`, 3},
		{`11111`, `11111000000`, 11},
		{`111100001010`, `1111000010100101011110000`, 25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FillData(tt.name, tt.length); got != tt.want {
				t.Errorf("FillData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecksum(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{`110010110100`, `100`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Checksum(tt.name); got != tt.want {
				t.Errorf("Checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartOne(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"10000\n20", `01100`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partOne(tt.name); got != tt.want {
				t.Errorf("partOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
