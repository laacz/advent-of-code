package main

import (
	"testing"
)

func TestRound(t *testing.T) {
	player := Player{8, 5, 5, "player"}
	boss := Player{12, 7, 2, "boss"}
	round(&player, &boss)
}

// func TestPartTwo(t *testing.T) {
// 	expect := 6
// 	if got := partTwo(`e => H
// 	e => O
// 	H => HO
// 	H => OH
// 	O => HH

// 	HOHOHO`); got != expect {
// 		t.Errorf("PartTwo() = %v, want %v", got, expect)
// 	}
// }
