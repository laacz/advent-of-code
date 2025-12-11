package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"0,3,6", 436},
		{"1,3,2", 1},
		{"2,1,3", 10},
		{"1,2,3", 27},
		{"2,3,1", 78},
		{"3,2,1", 438},
		{"3,1,2", 1836},
	}

	for _, tt := range tests {
		result := part1(parse(tt.input))
		if result != tt.expected {
			t.Errorf("part1(%s) = %d, expected %d", tt.input, result, tt.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"0,3,6", 175594},
	}

	for _, tt := range tests {
		result := part2(parse(tt.input))
		if result != tt.expected {
			t.Errorf("part2(%s) = %d, expected %d", tt.input, result, tt.expected)
		}
	}
}

func BenchmarkPart2(b *testing.B) {
	input := parse("0,3,6")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(input)
	}
}
