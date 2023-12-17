package main

import (
	"fmt"
	"os"
	"strings"
)

type Contraption struct {
	grid    [][]rune
	visited map[Coords]bool
	beams   []Beam
}

type Coords [2]int

type Beam struct {
	x, y   int
	dx, dy int
}

func (c Contraption) String() (ret string) {
	for y, row := range c.grid {
		for x, col := range row {
			if c.visited[Coords{y, x}] {
				ret += "#"
			} else {
				ret += string(col)
			}
		}
		ret += "\n"
	}

	ret += fmt.Sprintf("%+v", c.beams)

	return ret
}

func (c *Contraption) Energize() {
	for i := range c.beams {
		beam := &c.beams[i]

		x1 := beam.x + beam.dx
		y1 := beam.y + beam.dy

		if x1 < 0 || x1 >= len(c.grid[0]) || y1 < 0 || y1 >= len(c.grid) {
			continue
		}

		c.visited[Coords{y1, x1}] = true

		beam.x = x1
		beam.y = y1

		if c.grid[y1][x1] == '\\' {
			beam.dx, beam.dy = beam.dy, beam.dx
		} else if c.grid[y1][x1] == '/' {
			beam.dx, beam.dy = -beam.dy, -beam.dx
		} else if c.grid[y1][x1] == '|' && beam.dx != 0 {
			beam.dx, beam.dy = beam.dy, beam.dx
			c.beams = append(c.beams, Beam{x1, y1, beam.dx, -beam.dy})
		} else if c.grid[y1][x1] == '-' && beam.dy != 0 {
			beam.dx, beam.dy = beam.dy, beam.dx
			c.beams = append(c.beams, Beam{x1, y1, -beam.dx, beam.dy})
		}

		if c.grid[y1][x1] == '.' {
			continue
		}
	}
}

func parse(input string) Contraption {
	ret := Contraption{
		visited: make(map[Coords]bool),
	}

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		ret.grid = append(ret.grid, []rune(line))
	}

	return ret
}

func (c *Contraption) Shoot(b Beam) {
	c.beams = append(c.beams, b)

	i, j := 0, 0
	for {
		cnt := len(c.visited)
		c.Energize()
		if cnt == len(c.visited) {
			if j+5 < i {
				break
			}
		} else {
			j = i
		}
		i++
	}

}

// partOne returns the answer to part one of this day's puzzle.
func partOne(input string) int {
	g := parse(input)
	g.Shoot(Beam{-1, 0, 1, 0})
	return len(g.visited)
}

// partTwo returns the answer to part one of this day's puzzle.
func partTwo(input string) (ret int) {
	g := parse(input)

	beams := []Beam{}
	for x := 0; x < len(g.grid[0]); x++ {
		beams = append(beams, Beam{x, -1, 0, 1})
		beams = append(beams, Beam{x, len(g.grid), 0, -1})
	}

	for y := 0; y < len(g.grid); y++ {
		beams = append(beams, Beam{-1, y, 1, 0})
		beams = append(beams, Beam{len(g.grid[0]), y, -1, 0})
	}

	for i, b := range beams {
		q := Contraption{
			grid:    g.grid,
			visited: make(map[Coords]bool),
		}

		q.Shoot(b)
		if len(q.visited) > ret {
			fmt.Println("New record:", len(q.visited), "with", len(q.beams), "beams, starting at", b, "take", i, "of", len(beams))
			ret = len(q.visited)
		}
	}

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
