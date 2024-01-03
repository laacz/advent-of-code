package main

import "testing"

func TestPartOne(t *testing.T) {
	got := partOne(`aaaaa-bbb-z-y-x-123[abxyz]
	a-b-c-d-e-f-g-h-987[abcde]
	not-a-real-room-404[oarel]
	totally-real-room-200[decoy]`)
	expected := 1514
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
func TestPartTwo(t *testing.T) {
	room := NewRoom(`qzmt-zixmtkozy-ivhz-343[zimth]`)
	got := room.Decrypt()
	expected := "very encrypted name"
	if got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}
