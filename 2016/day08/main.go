package main

import (
	"fmt"

	"github.com/laacz/aoc-2016/util"
)

const (
	RECT = iota
	ROTATE_ROW
	ROTATE_COL
)

type Instruction struct {
	t    int
	x, y int
}

func (i Instruction) String() (ret string) {
	switch i.t {
	case RECT:
		ret = fmt.Sprintf("rect %dx%d", i.x, i.y)
	case ROTATE_ROW:
		ret = fmt.Sprintf("rotate row y=%d by %d", i.x, i.y)
	case ROTATE_COL:
		ret = fmt.Sprintf("rotate column x=%d by %d", i.x, i.y)
	}
	return ret
}

type Display [][]bool

func (d Display) String() (ret string) {
	for _, row := range d {
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

func (d Display) Rect(x, y int) {
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			d[i][j] = true
		}
	}
}

func (d Display) RotateRow(y, by int) {
	for i := 0; i < by; i++ {
		d[y] = append(d[y][len(d[y])-1:], d[y][:len(d[y])-1]...)
	}
}

func (d Display) RotateCol(x, by int) {
	c := make([]bool, len(d))
	for i := range c {
		c[i] = d[i][x]
	}
	for i := 0; i < len(c); i++ {
		d[(i+by)%len(d)][x] = c[i]
	}
}

func (d Display) PixelsLit() (ret int) {
	for _, row := range d {
		for _, col := range row {
			if col {
				ret++
			}
		}
	}
	return ret
}

func parse(input string) (ret []Instruction) {
	for _, line := range util.GetLines(input) {
		var i Instruction

		if line[0:4] == "rect" {
			i.t = RECT
			fmt.Sscanf(line, "rect %dx%d", &i.x, &i.y)
		} else if line[0:11] == "rotate row " {
			i.t = ROTATE_ROW
			fmt.Sscanf(line, "rotate row y=%d by %d", &i.x, &i.y)
		} else if line[0:13] == "rotate column" {
			i.t = ROTATE_COL
			fmt.Sscanf(line, "rotate column x=%d by %d", &i.x, &i.y)
		}

		ret = append(ret, i)

	}
	return ret
}

func partOne(input string, w, h int) (ret int) {
	instructions := parse(input)
	d := make(Display, h)
	for i := range d {
		d[i] = make([]bool, w)
	}
	for _, i := range instructions {
		switch i.t {
		case RECT:
			d.Rect(i.x, i.y)
		case ROTATE_ROW:
			d.RotateRow(i.x, i.y)
		case ROTATE_COL:
			d.RotateCol(i.x, i.y)
		}
	}
	return d.PixelsLit()
}

func partTwo(input string, w, h int) (ret string) {
	instructions := parse(input)
	d := make(Display, h)
	for i := range d {
		d[i] = make([]bool, w)
	}
	for _, i := range instructions {
		switch i.t {
		case RECT:
			d.Rect(i.x, i.y)
		case ROTATE_ROW:
			d.RotateRow(i.x, i.y)
		case ROTATE_COL:
			d.RotateCol(i.x, i.y)
		}
	}
	return d.String()
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println("Part one:", partOne(input, 50, 6))
	fmt.Println("Part two:\n", partTwo(input, 50, 6))
}
