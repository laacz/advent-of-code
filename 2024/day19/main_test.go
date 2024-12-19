package main

import (
	"testing"
)

var input = `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

func TestPartOne(t *testing.T) {
	comp := parseInput(input)
	expected := 6
	if actual := part1(comp); actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	comp := parseInput(input)
	expected := 16
	if actual := part2(comp); actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}
