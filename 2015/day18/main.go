package main

import (
	"fmt"
	"os"
	"strings"
)

type Field [][]bool

func (f Field) String() (ret string) {
	for _, row := range f {
		for _, col := range row {
			if col {
				ret += "#"
			} else {
				ret += "."
			}
		}
		ret += "\n"
	}

	return ret
}

var dirs = [][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, 1}, {1, 0}, {1, -1},
}

func (f *Field) Step(second bool) {
	field := make(Field, len(*f))
	if second {
		(*f)[0][0] = true
		(*f)[0][len(*f)-1] = true
		(*f)[len(*f)-1][0] = true
		(*f)[len(*f)-1][len(*f)-1] = true
	}

	for y, row := range *f {
		field[y] = make([]bool, len(row))

		for x, col := range row {
			if second && (x == 0 || x == len(row)-1) && (y == 0 || y == len(*f)-1) {
				field[y][x] = true
			}
			var on int
			for _, dir := range dirs {
				if y+dir[0] < 0 || y+dir[0] >= len(*f) {
					continue
				}
				if x+dir[1] < 0 || x+dir[1] >= len(row) {
					continue
				}
				if (*f)[y+dir[0]][x+dir[1]] {
					on++
				}
			}

			if col && (on == 2 || on == 3) {
				field[y][x] = true
			}

			if !col && on == 3 {
				field[y][x] = true
			}
		}
	}

	*f = field
}

func (f Field) CountOn() (ret int) {
	for _, row := range f {
		for _, col := range row {
			if col {
				ret++
			}
		}
	}
	return ret
}

func parse(lines string) (ret Field) {
	for _, line := range strings.Split(lines, "\n") {
		row := []bool{}
		for _, char := range line {
			if char == '#' {
				row = append(row, true)
			} else {
				row = append(row, false)
			}
		}
		ret = append(ret, row)
	}

	return ret
}

func partOne(lines string, iterations int) (ret int) {
	field := parse(lines)

	for i := 0; i < iterations; i++ {
		field.Step(false)
	}

	return field.CountOn()
}

func partTwo(lines string, iterations int) (ret int) {
	field := parse(lines)

	for i := 0; i < iterations; i++ {
		field.Step(true)
	}

	return field.CountOn()
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %d\n", partOne(string(data), 100))
	fmt.Printf("Part two: %d\n", partTwo(string(data), 100))
}
