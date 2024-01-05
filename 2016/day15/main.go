package main

import (
	"fmt"

	"github.com/laacz/aoc-2016/util"
)

type Disc struct {
	num       int
	positions int
	start     int
}

type Discs []Disc

func parse(input string) (ret Discs) {
	// Disc #1 has 17 positions; at time=0, it is at position 1.
	// Disc #2 has 7 positions; at time=0, it is at position 0.
	// Disc #3 has 19 positions; at time=0, it is at position 2.
	// Disc #4 has 5 positions; at time=0, it is at position 0.
	// Disc #5 has 3 positions; at time=0, it is at position 0.
	// Disc #6 has 13 positions; at time=0, it is at position 5.
	for _, line := range util.GetLines(input) {
		var d Disc
		fmt.Sscanf(line, "Disc #%d has %d positions; at time=0, it is at position %d.", &d.num, &d.positions, &d.start)
		ret = append(ret, d)
	}

	return
}

func partOne(input string) (ret int) {
	d := parse(input)
	ret = 1
outer:
	for {
		for i, disc := range d {
			pos := (disc.start + ret + i + 1) % disc.positions
			if pos%disc.positions != 0 {
				ret++
				continue outer
			}
		}
		return
	}
}

func partTwo(input string) (ret int) {
	input += "\nDisc #7 has 11 positions; at time=0, it is at position 0."
	d := parse(input)
	ret = 1
outer:
	for {
		for i, disc := range d {
			pos := (disc.start + ret + i + 1) % disc.positions
			if pos%disc.positions != 0 {
				ret++
				continue outer
			}
		}
		return
	}
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}
