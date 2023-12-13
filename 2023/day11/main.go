package main

import (
	"fmt"
	"os"
	"strings"
)

type Map struct {
	cols  []int
	rows  []int
	stars map[[2]int]int
}

// String returns a string representation of the Map
func (m Map) String() string {
	ret := ""

	for y, ny := range m.rows {
		for x, nx := range m.cols {
			if m.stars[[2]int{y, x}] > 0 {
				ret += "#"
			} else if nx > 1 && ny > 1 {
				ret += "┼"
			} else if nx > 1 {
				ret += "│"
			} else if ny > 1 {
				ret += "─"
			} else {
				ret += "·"
			}
		}
		ret += "\n"
	}

	ret += fmt.Sprintf("cols: %v\n", m.cols)
	ret += fmt.Sprintf("rows: %v\n", m.rows)
	ret += fmt.Sprintf("stars: %v\n", m.stars)

	return ret
}

// Expand doubles empty rows and columns
func (m *Map) Expand(amount int) {
	cols := make(map[int]bool)
	rows := make(map[int]bool)

	for y := range m.rows {
		for x := range m.cols {
			if m.stars[[2]int{y, x}] > 0 {
				cols[x] = true
				rows[y] = true
			}
		}
	}

	for y := 0; y < len(m.rows); y++ {
		if !rows[y] {
			m.rows[y] = amount
		}
	}

	for x := 0; x < len(m.cols); x++ {
		if !cols[x] {
			m.cols[x] = amount
		}
	}

}

// parse (you won't believe it!) returns parsed input
func parse(input string) Map {
	ret := Map{
		stars: make(map[[2]int]int),
	}

	idx := 0
	for y, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		ret.rows = append(ret.rows, 1)
		for x, c := range line {
			if c == '#' {
				idx += 1
				ret.stars[[2]int{y, x}] = idx
			}
		}
	}

	for range strings.Split(input, "\n")[0] {
		ret.cols = append(ret.cols, 1)
	}

	return ret
}

func (m *Map) SumDistances() int {
	sum := 0

	visited := make(map[[2]int]bool)

	for coords1, s1 := range m.stars {
		for coords2, s2 := range m.stars {
			if s1 == s2 || visited[[2]int{s2, s1}] || visited[[2]int{s1, s2}] {
				continue
			}

			dx := 0
			dy := 0
			y := min(coords1[0], coords2[0])
			x := min(coords1[1], coords2[1])

			for y < max(coords1[0], coords2[0]) {
				y += 1
				dy += m.rows[y]
			}

			for x < max(coords1[1], coords2[1]) {
				x += 1
				dx += m.cols[x]
			}

			visited[[2]int{s1, s2}] = true

			sum += dx + dy
		}
	}

	return sum
}

// partOne returns the answer to part one of this day's puzzle.
func partOne(input string) int {
	m := parse(input)
	m.Expand(2)
	fmt.Println(m)

	return m.SumDistances()
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// partTwo returns the answer to part one of this day's puzzle.
func partTwo(input string) int {
	m := parse(input)
	m.Expand(1000000)
	fmt.Println(m)

	return m.SumDistances()
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
