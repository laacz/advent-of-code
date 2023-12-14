package main

import (
	"fmt"
	"os"
	"strings"
)

type Field struct {
	rocks []string
}

func (f *Field) Slide(dx, dy int) {
	for {
		moved := false
		for y, row := range f.rocks {
			for x, rock := range row {

				if rock == 'O' {
					y1 := y + dy
					x1 := x + dx

					if y1 < 0 || y1 >= len(f.rocks) ||
						x1 < 0 || x1 >= len(row) {
						continue
					}

					if f.rocks[y1][x1] != '.' {
						continue
					}

					f.rocks[y1] = f.rocks[y1][:x1] + "O" + f.rocks[y1][x1+1:]
					f.rocks[y] = f.rocks[y][:x] + "." + f.rocks[y][x+1:]

					moved = true
				}
			}
		}

		if !moved {
			break
		}
	}
}

func (f Field) Cycle() (ret int) {
	dirs := [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

	for _, dir := range dirs {
		f.Slide(dir[0], dir[1])
	}

	return ret
}

func (f Field) Load() (ret int) {
	for y, row := range f.rocks {
		for _, rock := range row {
			if rock == 'O' {
				ret += len(f.rocks) - y
			}
		}
	}

	return ret
}

func (f Field) String() (ret string) {
	for _, row := range f.rocks {
		ret += string(row) + "\n"
	}

	return ret
}

func parse(input string) (ret Field) {
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}

		ret.rocks = append(ret.rocks, line)
	}

	return ret
}

// partOne returns the answer to part one of this day's puzzle.
func partOne(input string) int {
	f := parse(input)

	f.Slide(0, -1)
	return f.Load()
}

// detectCylcle detects a cycle in a list of integers, given that it can start not at the beginning of the list and can be of artbitrary length.
func detectCycle(l []int) (offset, length int) {
	if len(l) < 2 {
		return 0, 0
	}
	for i := 0; i < len(l)/2; i++ {
		b := l[i:]
		if len(b)%2 != 0 {
			continue
		}

		if eq(b[:len(b)/2], b[len(b)/2:]) {
			return i, len(b) / 2
		}
	}

	return 0, 0
}

// eq compares two slices of integers and returns true if they are equal.
func eq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, n := range a {
		if n != b[i] {
			return false
		}
	}

	return true
}

// partTwo returns the answer to part one of this day's puzzle.
func partTwo(input string) int {
	f := parse(input)

	var loads []int
	var offset, length, breakat int

	loads = append(loads, f.Load())
	for i := 0; i < 1000000000; i++ {
		f.Cycle()
		loads = append(loads, f.Load())

		offset, length = detectCycle(loads)
		if length > 1 && breakat == 0 {
			breakat = i + (1000000000-offset)%length
		}

		if i == breakat+1 && breakat != 0 {
			break
		}
	}

	return f.Load()
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
