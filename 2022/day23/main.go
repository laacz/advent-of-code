package main

import (
	"fmt"
	"os"
	"strings"
)

var Directions = [][]Coords{
	{
		// North
		{-1, -1},
		{0, -1},
		{1, -1},
	},
	{
		// South
		{-1, 1},
		{0, 1},
		{1, 1},
	},
	{
		// West
		{-1, -1},
		{-1, 0},
		{-1, 1},
	},
	{
		// East
		{1, -1},
		{1, 0},
		{1, 1},
	},
}

type Coords struct {
	X, Y int
}

type Elf struct {
	Num         int
	ProposedPos *Coords
	CanMove     bool
}

var DirIndex int

type Elves map[Coords]*Elf

func (ee *Elves) Consider() {
	for pos, elf := range *ee {
		// consider all directions
		elf.CanMove = false
	directions:
		for _, dir := range Directions {
			for _, d := range dir {
				if _, ok := (*ee)[Coords{pos.X + d.X, pos.Y + d.Y}]; ok {
					elf.CanMove = true
					break directions
				}
			}
		}
	}

	for pos, elf := range *ee {
		elf.ProposedPos = nil
		// consider each direction
		if elf.CanMove {
			// iterate over directions
			elf.CanMove = false
		adjacents:
			for i := 0; i < 4; i++ {
				dirIndex := (DirIndex + i) % 4
				dir := Directions[dirIndex]
				// consider each position in the direction
				for _, d := range dir {
					if _, ok := (*ee)[Coords{pos.X + d.X, pos.Y + d.Y}]; ok {
						// if there's an elf in the way, don't move
						elf.CanMove = false
						continue adjacents
					}
				}
				elf.ProposedPos = &Coords{pos.X + dir[1].X, pos.Y + dir[1].Y}
				elf.CanMove = true
				break
			}
		}
	}
	DirIndex = (DirIndex + 1) % 4
}

func Move(ee *Elves) int {
	moves := 0

	for _, elf := range *ee {
		if elf.CanMove {
			// see if any other elf is going to the same spot
			for _, other := range *ee {
				if other.ProposedPos != nil && other != elf && *other.ProposedPos == *elf.ProposedPos {
					// if so, don't move
					elf.CanMove = false
					other.CanMove = false
					break
				}
			}
		}
	}

Loop:
	for {
		for pos, elf := range *ee {
			// if elf still can move, move
			if elf.CanMove {
				elf.CanMove = false
				moves++
				(*ee)[*elf.ProposedPos] = elf
				delete(*ee, pos)
				continue Loop
			}
		}
		break
	}
	return moves
}

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

func (ee *Elves) CountEmptyTiles() int {
	maxx, maxy, minx, miny := 0, 0, 10, 10
	for c := range *ee {
		maxx = max(maxx, c.X)
		maxy = max(maxy, c.Y)
		minx = min(minx, c.X)
		miny = min(miny, c.Y)
	}

	return (maxx-minx+1)*(maxy-miny+1) - len(*ee)
}

func NewMap(input string) Elves {
	ee := make(Elves)
	i := 0
	for y, line := range strings.Split(input, "\n") {
		for x, c := range line {
			if c == '#' {
				i++
				ee[Coords{x, y}] = &Elf{Num: i}
			}
		}
	}
	return ee
}

func partOne(input string) int {
	m := NewMap(input)
	for i := 0; i < 10; i++ {
		m.Consider()
		Move(&m)
	}
	return m.CountEmptyTiles()
}

func partTwo(input string) int {
	m := NewMap(input)
	DirIndex = 0
	for i := 0; true; i++ {
		m.Consider()
		moves := Move(&m)
		if moves == 0 {
			return i + 1
		}
	}
	return 0
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
