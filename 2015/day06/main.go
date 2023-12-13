package main

import (
	"fmt"
	"os"
	"strings"
)

type Coords [2]int

type Instruction struct {
	operation string
	from      Coords
	to        Coords
}

func parse(lines string) []Instruction {
	ret := make([]Instruction, 0)

	for i, line := range strings.Split(lines, "\n") {
		if line == "" {
			continue
		}

		line = strings.TrimSpace(line)

		var instruction Instruction
		var fromX, fromY, toX, toY int
		var operation string

		if strings.HasPrefix(line, "turn on") {
			operation = "on"
			fmt.Sscanf(line, "turn on %d,%d through %d,%d", &fromX, &fromY, &toX, &toY)
		} else if strings.HasPrefix(line, "turn off") {
			operation = "off"
			fmt.Sscanf(line, "turn off %d,%d through %d,%d", &fromX, &fromY, &toX, &toY)
		} else if strings.HasPrefix(line, "toggle") {
			operation = "toggle"
			fmt.Sscanf(line, "toggle %d,%d through %d,%d", &fromX, &fromY, &toX, &toY)
		} else {
			panic(fmt.Sprintf("unknown operation on line %d: %s", i, line))
		}

		instruction.operation = operation
		instruction.from = Coords{fromX, fromY}
		instruction.to = Coords{toX, toY}

		ret = append(ret, instruction)
	}

	return ret
}

type Grid map[Coords]int

func (g Grid) Process(i Instruction) {
	for x := i.from[0]; x <= i.to[0]; x++ {
		for y := i.from[1]; y <= i.to[1]; y++ {
			switch i.operation {
			case "on":
				g[Coords{x, y}] = 1
			case "off":
				g[Coords{x, y}] = 0
			case "toggle":
				if g[Coords{x, y}] == 0 {
					g[Coords{x, y}] = 1
				} else {
					g[Coords{x, y}] = 0
				}
			}
		}
	}
}

func (g Grid) Process2(i Instruction) {
	for x := i.from[0]; x <= i.to[0]; x++ {
		for y := i.from[1]; y <= i.to[1]; y++ {
			switch i.operation {
			case "on":
				g[Coords{x, y}] += 1
			case "off":
				if g[Coords{x, y}] > 0 {
					g[Coords{x, y}] -= 1
				}
			case "toggle":
				g[Coords{x, y}] += 2
			}
		}
	}
}

func (g Grid) Count() int {
	ret := 0
	for _, v := range g {
		if v > 0 {
			ret++
		}
	}
	return ret
}

func (g Grid) Brightness() int {
	ret := 0
	for _, v := range g {
		ret += v
	}
	return ret
}

func partOne(lines string) int {
	var grid Grid = make(map[Coords]int)

	instructions := parse(lines)
	for _, i := range instructions {
		grid.Process(i)
	}

	return grid.Count()
}

func partTwo(lines string) int {
	var grid Grid = make(map[Coords]int)

	instructions := parse(lines)
	for _, i := range instructions {
		grid.Process2(i)
	}

	return grid.Brightness()
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
