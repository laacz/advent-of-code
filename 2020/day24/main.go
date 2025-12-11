package main

import (
	"fmt"
	"maps"
	"os"
	"strings"
)

// Red Blob Games hex grid reference FTW!
// https://www.redblobgames.com/grids/hexagons/#neighbors
func parse(input string) [][]string {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")

	ret := [][]string{}

	for _, line := range strings.Split(parts[0], "\n") {
		steps := []string{}
		for i := 0; i < len(line); {
			if line[i] == 's' || line[i] == 'n' {
				steps = append(steps, line[i:i+2])
				i += 2
			} else {
				steps = append(steps, line[i:i+1])
				i++
			}
		}
		ret = append(ret, steps)
	}

	return ret
}

type Coords [3]int

type Floor struct {
	tiles map[Coords]bool
	pos   Coords
}

func NewFloor() *Floor {
	return &Floor{
		tiles: make(map[Coords]bool),
		pos:   Coords{0, 0, 0},
	}
}

const BLACK = true
const WHITE = false

var DIRS = map[string]Coords{
	"e":  {1, -1, 0},
	"w":  {-1, 1, 0},
	"ne": {1, 0, -1},
	"nw": {0, 1, -1},
	"se": {0, -1, 1},
	"sw": {-1, 0, 1},
}

func (f *Floor) step(direction string) {
	dir := DIRS[direction]
	f.pos[0] += dir[0]
	f.pos[1] += dir[1]
	f.pos[2] += dir[2]
}

func (f *Floor) blackNeighboursCount(coord Coords) int {
	count := 0

	for _, dir := range DIRS {
		neighbour := Coords{
			coord[0] + dir[0],
			coord[1] + dir[1],
			coord[2] + dir[2],
		}
		if f.tiles[neighbour] == BLACK {
			count++
		}
	}

	return count
}

func (f *Floor) DailyFlip() {
	// check all tiles, even the white ones next to black ones
	// took half an hour to hunt this down...
	tilesToCheck := make(map[Coords]bool)

	for coord, isBlack := range f.tiles {
		if isBlack {
			tilesToCheck[coord] = true
			for _, dir := range DIRS {
				neighbour := Coords{
					coord[0] + dir[0],
					coord[1] + dir[1],
					coord[2] + dir[2],
				}
				tilesToCheck[neighbour] = true
			}
		}
	}

	flippedTiles := make(map[Coords]bool)

	for coord := range tilesToCheck {
		cnt := f.blackNeighboursCount(coord)
		isBlack := f.tiles[coord]

		if isBlack && (cnt == 0 || cnt > 2) {
			flippedTiles[coord] = WHITE
		} else if !isBlack && cnt == 2 {
			flippedTiles[coord] = BLACK
		}
	}

	maps.Copy(f.tiles, flippedTiles)
}

func (f *Floor) CountBlackTiles() int {
	count := 0
	for _, tile := range f.tiles {
		if tile == BLACK {
			count++
		}
	}
	return count
}

func part1(instructions [][]string) int {
	floor := NewFloor()

	for _, tile := range instructions {
		floor.pos = Coords{0, 0, 0}
		for _, step := range tile {
			floor.step(step)
		}

		floor.tiles[floor.pos] = !floor.tiles[floor.pos]
	}

	return floor.CountBlackTiles()
}

func part2(instructions [][]string) int {
	floor := NewFloor()

	for _, tile := range instructions {
		floor.pos = Coords{0, 0, 0}
		for _, step := range tile {
			floor.step(step)
		}

		floor.tiles[floor.pos] = !floor.tiles[floor.pos]
	}

	fmt.Printf("Day 0: %d black tiles\n", floor.CountBlackTiles())

	for day := 1; day <= 100; day++ {
		floor.DailyFlip()
		fmt.Printf("Day %d: %d black tiles\n", day, floor.CountBlackTiles())
	}

	return floor.CountBlackTiles()
}

func main() {
	data, _ := os.ReadFile("input.txt")
	instructions := parse(string(data))

	fmt.Println("Part 1:", part1(instructions))
	fmt.Println("Part 2:", part2(instructions))
}
