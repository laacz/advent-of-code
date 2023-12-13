package main

import (
	"testing"
)

var input string = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestType(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"AAAAA", FiveOfAKind},
		{"AAAAB", FourOfAKind},
		{"AAABA", FourOfAKind},
		{"AAABB", FullHouse},
		{"AABBB", FullHouse},
		{"AAA98", ThreeOfAKind},
		{"AABBC", TwoPair},
		{"AABCD", OnePair},
		{"ABCDE", HighCard},
	}

	for _, c := range cases {
		h := Hand{c.in, 0}
		got := h.Type()
		if got != c.want {
			t.Errorf("Type(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestPartOne(t *testing.T) {
	expect := 6440
	actual := partOne(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 5905
	actual := partTwo(input)

	if actual != expect {
		t.Errorf("Expected %d, got %d", expect, actual)
	}
}
