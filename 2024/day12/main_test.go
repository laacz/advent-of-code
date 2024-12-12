package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	for _, tt := range []struct {
		input  string
		expect int
	}{
		{"AA\nAA", 8 * 4},
		{"AA\nBB", 2*6 + 2*6},
		{"A", 4},
		{"AAA\nBBB\nCCC", 3 * 3 * 8},
		{"AAAA\nBBCD\nBBCC\nEEEC", 140},
		{"OOOOO\nOXOXO\nOOOOO\nOXOXO\nOOOOO", 772},
		{"RRRRIICCFF\nRRRRIICCCF\nVVRRRCCFFF\nVVRCCCJFFF\nVVVVCJJCFE\nVVIVCCJJEE\nVVIIICJJEE\nMIIIIIJJEE\nMIIISIJEEE\nMMMISSJEEE", 1930},
	} {

		actual := part1(parseInput(tt.input))

		if actual != tt.expect {
			t.Errorf("Expected %d, got %d", tt.expect, actual)
		}
	}
}
func TestPartTwo(t *testing.T) {
	for _, tt := range []struct {
		input  string
		expect int
	}{
		{"AA\nAA", 4 * 4},
		{"AB\nBA", 4 * 4},
		{"AA\nBB", 2*4 + 2*4},
		{"A", 1 * 4},
		{"AAA\nBBB\nCCC", 3 * 3 * 4},
		{"AAAA\nBBCD\nBBCC\nEEEC", 80},
		{"OOOOO\nOXOXO\nOOOOO\nOXOXO\nOOOOO", 436},
		{"AAAAAA\nAAABBA\nAAABBA\nABBAAA\nABBAAA\nAAAAAA", 368},
		{"RRRRIICCFF\nRRRRIICCCF\nVVRRRCCFFF\nVVRCCCJFFF\nVVVVCJJCFE\nVVIVCCJJEE\nVVIIICJJEE\nMIIIIIJJEE\nMIIISIJEEE\nMMMISSJEEE", 1206},
	} {

		actual := part2(parseInput(tt.input))

		if actual != tt.expect {
			t.Errorf("Expected %d, got %d", tt.expect, actual)
		}
	}
}
