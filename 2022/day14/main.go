package main

import (
	"fmt"
	"os"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Point struct {
	X int
	Y int
}

func parseInput(input string) map[Point]rune {
	ret := make(map[Point]rune)

	for _, line := range strings.Split(input, "\n") {
		var x, y int
		for _, path := range strings.Split(line, " -> ") {
			var x1, y1 int
			fmt.Sscanf(path, "%d,%d", &x1, &y1)
			if x != 0 {
				for a := min(x, x1); a <= max(x, x1); a++ {
					for b := min(y, y1); b <= max(y, y1); b++ {
						ret[Point{a, b}] = 'x'
					}
				}
			}
			x = x1
			y = y1
		}
	}

	return ret
}

func isInfiniteFall(m *map[Point]rune) bool {
	x := 500
	y := 0
loop:
	for {
		if y > 200 {
			return true
		}

		for _, dx := range []int{0, -1, 1} {
			_, ok := (*m)[Point{x + dx, y + 1}]
			if !ok {
				x += dx
				y += 1
				continue loop
			}
		}

		break
	}

	(*m)[Point{x, y}] = 'o'
	return false
}

func partOne(input string) int {
	ret := 0
	m := parseInput(input)

	for !isInfiniteFall(&m) {
		ret += 1
	}

	return ret
}

func isFallless(m *map[Point]rune, floor int) bool {
	x := 500
	y := 0
loop:
	for {
		if y == floor-1 {
			break
		}
		for _, dx := range []int{0, -1, 1} {
			_, ok := (*m)[Point{x + dx, y + 1}]
			if !ok {
				x += dx
				y += 1
				continue loop
			}
		}
		if y == 0 {
			return true
		}

		break
	}
	(*m)[Point{x, y}] = 'o'
	return false
}

func partTwo(input string) int {
	ret := 0
	m := parseInput(input)

	floor := 0
	for p := range m {
		floor = max(floor, p.Y)
	}

	floor += 2
	for !isFallless(&m, floor) {
		ret += 1
	}

	return ret + 1
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
