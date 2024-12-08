package main

import (
	"fmt"
	"os"
	"strings"
)

type Coord [2]int

type Map struct {
	Size      int
	Antennas  map[Coord]byte
	Antinodes map[Coord]byte
}

func (m *Map) String() string {
	var ret string

	for y := 0; y < m.Size; y++ {
		for x := 0; x < m.Size; x++ {
			if m.Antinodes[Coord{x, y}] != 0 {
				ret += "\033[47;30m"
			}
			if c, ok := m.Antennas[Coord{x, y}]; ok {
				ret += fmt.Sprintf("%c", c)
			} else {
				ret += "·"
			}
			ret += "\033[0m"
		}
		ret += "\n"
	}

	return ret
}

func parseInput(input string) *Map {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	ret := Map{
		Size:      len(lines),
		Antennas:  make(map[Coord]byte),
		Antinodes: make(map[Coord]byte),
	}

	for y, line := range lines {
		for x, c := range line {
			if c != '.' {
				ret.Antennas[Coord{x, y}] = byte(c)
			}
		}
	}
	return &ret
}

func part1(input string) int {
	m := parseInput(input)

	for xy1, a := range m.Antennas {
		for xy2, b := range m.Antennas {
			if a != b || xy1 == xy2 {
				continue
			}

			d := Coord{xy2[0] - xy1[0], xy2[1] - xy1[1]}
			if xy2[0]+d[0] < 0 || xy2[0]+d[0] >= m.Size || xy2[1]+d[1] < 0 || xy2[1]+d[1] >= m.Size {
				continue
			}
			m.Antinodes[Coord{xy2[0] + d[0], xy2[1] + d[1]}] += 1
			if xy1[0]-d[0] < 0 || xy1[0]-d[0] >= m.Size || xy1[1]-d[1] < 0 || xy1[1]-d[1] >= m.Size {
				continue
			}
			m.Antinodes[Coord{xy1[0] - d[0], xy1[1] - d[1]}] += 1
		}
	}

	fmt.Println(m)

	return len(m.Antinodes)
}

func part2(input string) int {
	m := parseInput(input)

	for xy1, a := range m.Antennas {
		for xy2, b := range m.Antennas {
			if a != b || xy1 == xy2 {
				continue
			}

			d := Coord{xy2[0] - xy1[0], xy2[1] - xy1[1]}

			for i := range m.Size {
				xy := Coord{xy2[0] + d[0]*i, xy2[1] + d[1]*i}
				if xy[0] < 0 || xy[0] >= m.Size || xy[1] < 0 || xy[1] >= m.Size {
					break
				}
				m.Antinodes[xy] += 1
			}

			for i := range m.Size {
				xy := Coord{xy1[0] - d[0]*i, xy1[1] - d[1]*i}
				if xy[0] < 0 || xy[0] >= m.Size || xy[1] < 0 || xy[1] >= m.Size {
					break
				}
				m.Antinodes[xy] += 1
			}
		}
	}

	fmt.Println(m)

	return len(m.Antinodes)

}
func main() {

	input, _ := os.ReadFile("input.txt")

	fmt.Println("part1", part1(string(input)))
	fmt.Println("part2", part2(string(input)))
}
