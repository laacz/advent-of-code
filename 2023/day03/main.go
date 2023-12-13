package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Gear struct {
	x, y int
	char rune
}

type PartNumber struct {
	num           string
	x, y          int
	adjacent      bool
	adjacentGears []Gear
}

// NewPartNumber returns a new PartNumber as well as checks if it's adjacent to a symbol
// and stores all dajacent gears as well
func NewPartNumber(num string, x, y int, lines []string) PartNumber {
	pn := PartNumber{num, x, y, false, []Gear{}}

	for yy := y - 1; yy <= y+1; yy++ {
		if yy < 0 || yy >= len(lines) {
			continue
		}
		for xx := x - 1; xx <= x+len(num); xx++ {
			if xx >= 0 && xx < len(lines[0]) && lines[yy][xx] != '.' && lines[yy][xx]-'0' > 9 {
				pn.adjacent = true
				if lines[yy][xx] == '*' {
					pn.adjacentGears = append(pn.adjacentGears, Gear{xx, yy, '*'})
				}
			}
		}
	}

	return pn
}

// parse parses parsable parsablee
func parse(data string) []PartNumber {
	var ret []PartNumber

	var lines = strings.Split(strings.TrimSpace(data), "\n")

	cn := ""
	for y, line := range lines {
		for x, char := range line {
			if char-'0' >= 0 && char-'0' <= 9 {
				cn += string(char)
			} else {
				if cn != "" {
					ret = append(ret, NewPartNumber(cn, x-len(cn), y, lines))
				}
				cn = ""
			}
		}
	}

	// just in case there is a number in the bottom right corner...
	if cn != "" {
		ret = append(ret, NewPartNumber(cn, len(lines[0])-len(cn), len(lines)-1, lines))
	}

	return ret
}

// partOne returns the answer to part one of this day's puzzle.
func partOne(input string) int {
	sum := 0
	pns := parse(input)

	for _, pn := range pns {
		if pn.adjacent {
			i, _ := strconv.Atoi(pn.num)
			sum += i
		}
	}

	return sum
}

// partTwo returns the answer to part one of this day's puzzle.
func partTwo(input string) int {
	sum := 0
	pns := parse(input)

	// DRY my ass
	var lines = strings.Split(strings.TrimSpace(input), "\n")

	// now let's iterate over all the gears
	for y, row := range lines {
		for x, char := range row {
			if char == '*' {
				var nn []int

				// we'll try to find all the part numbers that are adjacent to this gear
				for _, pn := range pns {
					for _, gear := range pn.adjacentGears {
						if gear.x == x && gear.y == y {
							i, _ := strconv.Atoi(pn.num)
							nn = append(nn, i)
						}
					}
				}

				// as per the instructions, gear ratio is the product of all the part numbers, but we're
				// only interested in gears with more than one adjacent part
				if len(nn) > 1 {
					mul := 1
					for _, n := range nn {
						mul *= n
					}
					sum += mul
				}
			}
		}
	}

	return sum
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
