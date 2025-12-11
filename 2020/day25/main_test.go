package main

import "testing"

var cardPub, doorPub = parse(`5764801
17807724`)

func TestPart1(t *testing.T) {
	expected := 14897079
	actual := part1(cardPub, doorPub)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
