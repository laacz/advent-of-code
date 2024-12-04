package main

import (
	"fmt"
	"os"
	"strings"
)

type Coord [2]int

type Map map[Coord]byte

var directions = map[string]Coord{
	"N":  {0, -1},
	"S":  {0, 1},
	"E":  {1, 0},
	"W":  {-1, 0},
	"NE": {1, -1},
	"SE": {1, 1},
	"SW": {-1, 1},
	"NW": {-1, -1},
}

func (m *Map) HasWord(coord, dir Coord, word string) int {
	for i, c := range word {
		pos := Coord{coord[0] + dir[0]*i, coord[1] + dir[1]*i}
		if (*m)[pos] != byte(c) {
			return 0
		}
	}

	return 1
}

func parseInput(input string) (Map, int) {
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

func (m *Map) GetWord(pos, dir Coord, len int) string {
	var ret string

	for i := 0; i < len; i++ {
		ret += string((*m)[Coord{pos[0] + dir[0]*i, pos[1] + dir[1]*i}])
	}

	return ret
}

func part1(input string) int {
	var ret int

	m, l := parseInput(input)

	for y := range l {
		for x := range l {
			ret += m.HasWord(Coord{x, y}, directions["N"], "XMAS")
			ret += m.HasWord(Coord{x, y}, directions["E"], "XMAS")
			ret += m.HasWord(Coord{x, y}, directions["S"], "XMAS")
			ret += m.HasWord(Coord{x, y}, directions["W"], "XMAS")
			ret += m.HasWord(Coord{x, y}, directions["NE"], "XMAS")
			ret += m.HasWord(Coord{x, y}, directions["SE"], "XMAS")
			ret += m.HasWord(Coord{x, y}, directions["SW"], "XMAS")
			ret += m.HasWord(Coord{x, y}, directions["NW"], "XMAS")
		}
	}

	return ret
}

func part2(input string) int {
	var ret int

	m, l := parseInput(input)

	for y := range l {
		for x := range l {
			w1 := m.GetWord(Coord{x - 1, y - 1}, directions["SE"], 3)
			w2 := m.GetWord(Coord{x + 1, y - 1}, directions["SW"], 3)
			if (w1 == "MAS" || w1 == "SAM") && (w2 == "MAS" || w2 == "SAM") {
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
