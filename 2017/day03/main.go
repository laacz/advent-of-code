package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func SpiralCoord(n int) Coord {
	if n == 1 {
		return [2]int{0, 0}
	}
	sideLen := 1
	for sideLen*sideLen < n {
		sideLen += 2
	}

	start := (sideLen-2)*(sideLen-2) + 1
	side := (n - start) / (sideLen - 1)
	sideStart := start + side*(sideLen-1)
	sideMid := sideStart + (sideLen-1)/2 - 1
	pos := sideMid - n

	switch side {
	case 0:
		return Coord{sideLen / 2, pos}
	case 1:
		return Coord{pos, -sideLen / 2}
	case 2:
		return Coord{-sideLen / 2, -pos}
	case 3:
		return Coord{-pos, sideLen / 2}
	}

	panic("unreachable")
}

func part1(input int) int {
	pos := SpiralCoord(input)

	return int(math.Abs(float64(pos[0])) + math.Abs(float64(pos[1])))
}

type Coord [2]int

func part2(input int) int {
	ret := 1

	maze := map[Coord]int{
		{0, 0}: 1,
	}
	for i := 2; ; i++ {
		pos := SpiralCoord(i)

		ret = 0
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				if x == 0 && y == 0 {
					continue
				}
				ret += maze[Coord{pos[0] + x, pos[1] + y}]
			}
		}

		if ret > input {
			break
		}

		maze[pos] = ret
	}

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")

	input, _ := strconv.Atoi(strings.TrimSpace(string(data)))

	fmt.Println("part1", part1(input))
	fmt.Println("part2", part2(input))
}
