package main

import (
	"fmt"
	"os"
	"strings"
)

type Map [][]rune

func (m Map) String() string {
	ret := ""
	for _, row := range m {
		for _, col := range row {
			// inverse ansi
			if col == 'S' {
				ret += "\033[7m"
			}
			ret += string(col)
			ret += "\033[0m"
		}
		ret += "\n"
	}
	return ret
}

func (m Map) Start() (x, y int) {
	for y, row := range m {
		for x, col := range row {
			if col == 'S' {
				return x, y
			}
		}
	}
	return -1, -1
}

func (m *Map) Step() (ret int) {
	newmap := make(Map, len(*m))
	for y := range *m {
		newmap[y] = make([]rune, len((*m)[y]))
		for x := range (*m)[y] {
			if (*m)[y][x] == 'O' || (*m)[y][x] == 'S' {
				newmap[y][x] = '.'
			} else {
				newmap[y][x] = (*m)[y][x]
			}
		}
	}

	for y, row := range *m {
		for x, col := range row {
			if col == 'S' || col == 'O' {
				for _, dir := range []struct{ x, y int }{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
					if y+dir.y < 0 || y+dir.y >= len(*m) || x+dir.x < 0 || x+dir.x >= len((*m)[y]) {
						continue
					}
					if (*m)[y+dir.y][x+dir.x] != '#' {
						newmap[y+dir.y][x+dir.x] = 'O'
					}
				}
			}
		}
	}

	*m = newmap

	for y := range *m {
		for x := range (*m)[y] {
			if (*m)[y][x] == 'O' {
				ret += 1
			}
		}
	}

	return ret
}

// parse parses
func parse(input string) (ret Map) {
	for _, line := range strings.Split(input, "\n") {
		row := []rune{}
		for _, char := range line {
			row = append(row, char)
		}
		ret = append(ret, row)
	}

	return ret
}

// partOne returns the answer to part one of this day's puzzle.
func partOne(input string, steps int) (ret int) {
	m := parse(input)
	for i := 0; i < steps; i++ {
		ret = m.Step()
	}

	// fmt.Println(m)

	return ret
}

func (m Map) Solve(steps int) (ret int) {

	n := (steps*2+1)/len(m)/2*2 + 3

	mm := make(Map, n*len(m[0]))
	// fmt.Println("--=[ map size: ", len(mm), "x", len(mm), "]=---")

	for y := 0; y < n*len(m); y++ {
		mm[y] = make([]rune, n*len(m[0]))
		for x := 0; x < n*len(m[0]); x++ {
			xx := x % len(m[0])
			yy := y % len(m)

			mm[y][x] = m[yy][xx]
			// check if we're on the center square
			if (x/len(m[0]) != n/2 || y/len(m) != n/2) && m[yy][xx] == 'S' {
				// fmt.Println("remove S from", x/len(m[0]), y/len(m))
				mm[y][x] = '.'
			} else if m[yy][xx] == 'S' {
				// fmt.Println("found S at", x/len(m[0]), y/len(m))
			}

		}
	}

	for i := 0; i < steps; i++ {
		ret = mm.Step()
	}
	// fmt.Println("steps:", steps, "ret:", ret, " map size", len(mm), len(mm[0]))
	return ret
}

// partTwo returns the answer to part one of this day's puzzle.
func partTwo(input string, steps int) (ret int) {
	m := parse(input)

	n0 := m.Solve(65)
	n1 := m.Solve(196)
	n2 := m.Solve(327)

	// coeffs
	c := n0
	b := (4*n1 - 3*n0 - n2) / 2
	a := n1 - n0 - b

	x := (steps - len(m)/2) / len(m)

	ret = a*x*x + b*x + c

	// fmt.Println("a:", a, "b:", b, "c:", c, "steps:", steps, "ret:", ret)

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data), 64))
	fmt.Printf("Part two: %d\n", partTwo(string(data), 26501365))
}
