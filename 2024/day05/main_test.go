package main

import (
	"strings"
	"testing"
)

var input = `
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func TestPartOne(t *testing.T) {
	expect := 143
	actual := part1(strings.TrimSpace(input))

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 123
	actual := part2(strings.TrimSpace(input))

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
