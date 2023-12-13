package main

import (
	"testing"
)

var INPUT = `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`

// var INPUT = `Sensor at x=8, y=7: closest beacon is at x=2, y=10`

func TestPartOne(t *testing.T) {
	parseInput(INPUT)
	got := partOne(10)
	expect := 26
	if got != expect {
		t.Errorf("partOne(...) = %d, want %d", got, expect)
	}
}
func TestPartTwo(t *testing.T) {
	parseInput(INPUT)
	got := partTwo(0, 20)
	expect := 56000011
	if got != expect {
		t.Errorf("partTwo(...) got %d want %d", got, expect)
	}
}
