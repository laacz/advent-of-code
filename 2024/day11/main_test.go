package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	for _, tt := range []struct {
		input  []int
		blinks int
		expect int
	}{
		{[]int{0, 1, 10, 99, 999}, 1, 7},
		{[]int{125, 17}, 25, 55312},
	} {

		actual := part1(tt.input, tt.blinks)

		if actual != tt.expect {
			t.Errorf("Expected %d, got %d", tt.expect, actual)
		}
	}
}
func TestPartTwo(t *testing.T) {
	for _, tt := range []struct {
		input  []int
		blinks int
		expect int
	}{
		{[]int{0, 1, 10, 99, 999}, 1, 7},
		{[]int{125, 17}, 25, 55312},
	} {

		actual := part2(tt.input, tt.blinks)

		if actual != tt.expect {
			t.Errorf("Expected %d, got %d", tt.expect, actual)
		}
	}
}
