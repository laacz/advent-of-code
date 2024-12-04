package main

import (
	"fmt"
	"os"
	"strings"
)

type Coord [2]int

func (c Coord) Add(c2 Coord) Coord {
	return Coord{c[0] + c2[0], c[1] + c2[1]}
}

func parseInput(input string) (map[Coord]byte, int) {
	var l int
	ret := make(map[Coord]byte)

	for y, line := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, c := range line {
			ret[Coord{x, y}] = byte(c)
		}
		l = len(line)
	}

	return ret, l
}

func part1(input string) int {
	var ret int

	m, l := parseInput(input)

	for y := range l {
		for x := range l {
			if (m[Coord{x, y}] == 'X' && m[Coord{x + 1, y}] == 'M' && m[Coord{x + 2, y}] == 'A' && m[Coord{x + 3, y}] == 'S') {
				ret++
			}
			if (m[Coord{x, y}] == 'X' && m[Coord{x - 1, y}] == 'M' && m[Coord{x - 2, y}] == 'A' && m[Coord{x - 3, y}] == 'S') {
				ret++
			}
			if (m[Coord{x, y}] == 'X' && m[Coord{x, y + 1}] == 'M' && m[Coord{x, y + 2}] == 'A' && m[Coord{x, y + 3}] == 'S') {
				ret++
			}
			if (m[Coord{x, y}] == 'X' && m[Coord{x, y - 1}] == 'M' && m[Coord{x, y - 2}] == 'A' && m[Coord{x, y - 3}] == 'S') {
				ret++
			}
			if (m[Coord{x, y}] == 'X' && m[Coord{x + 1, y + 1}] == 'M' && m[Coord{x + 2, y + 2}] == 'A' && m[Coord{x + 3, y + 3}] == 'S') {
				ret++
			}
			if (m[Coord{x, y}] == 'X' && m[Coord{x - 1, y - 1}] == 'M' && m[Coord{x - 2, y - 2}] == 'A' && m[Coord{x - 3, y - 3}] == 'S') {
				ret++
			}
			if (m[Coord{x, y}] == 'X' && m[Coord{x + 1, y - 1}] == 'M' && m[Coord{x + 2, y - 2}] == 'A' && m[Coord{x + 3, y - 3}] == 'S') {
				ret++
			}
			if (m[Coord{x, y}] == 'X' && m[Coord{x - 1, y + 1}] == 'M' && m[Coord{x - 2, y + 2}] == 'A' && m[Coord{x - 3, y + 3}] == 'S') {
				ret++
			}
		}
	}

	return ret
}

func part2(input string) int {
	var ret int

	m, l := parseInput(input)

	for y := range l {
		for x := range l {
			if m[Coord{x, y}] == 'A' &&
				m[Coord{x - 1, y - 1}] == 'M' &&
				m[Coord{x - 1, y + 1}] == 'M' &&
				m[Coord{x + 1, y - 1}] == 'S' &&
				m[Coord{x + 1, y + 1}] == 'S' {
				ret++
			}
			if m[Coord{x, y}] == 'A' &&
				m[Coord{x - 1, y - 1}] == 'M' &&
				m[Coord{x - 1, y + 1}] == 'S' &&
				m[Coord{x + 1, y - 1}] == 'M' &&
				m[Coord{x + 1, y + 1}] == 'S' {
				ret++
			}
			if m[Coord{x, y}] == 'A' &&
				m[Coord{x - 1, y - 1}] == 'S' &&
				m[Coord{x - 1, y + 1}] == 'S' &&
				m[Coord{x + 1, y - 1}] == 'M' &&
				m[Coord{x + 1, y + 1}] == 'M' {
				ret++
			}
			if m[Coord{x, y}] == 'A' &&
				m[Coord{x - 1, y - 1}] == 'S' &&
				m[Coord{x - 1, y + 1}] == 'M' &&
				m[Coord{x + 1, y - 1}] == 'S' &&
				m[Coord{x + 1, y + 1}] == 'M' {
				ret++
			}
		}
	}

	return ret
}

func main() {

	input, _ := os.ReadFile("input.txt")

	fmt.Println("part1", part1(string(input)))
	fmt.Println("part2", part2(string(input)))
}
