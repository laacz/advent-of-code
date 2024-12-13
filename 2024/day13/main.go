package main

import (
	"fmt"
	"os"
	"strings"
)

type Coord [2]int

type Machine struct {
	A     Coord
	B     Coord
	Prize Coord
}

func parseInput(input string) []Machine {
	var ret []Machine

	for _, lines := range strings.Split(input, "\n\n") {
		machine := Machine{}
		for i, line := range strings.Split(lines, "\n") {
			switch i {
			case 0:
				fmt.Sscanf(line, "Button A: X+%d, Y+%d", &machine.A[0], &machine.A[1])
			case 1:
				fmt.Sscanf(line, "Button B: X+%d, Y+%d", &machine.B[0], &machine.B[1])
			case 2:
				fmt.Sscanf(line, "Prize: X=%d, Y=%d", &machine.Prize[0], &machine.Prize[1])
				ret = append(ret, machine)
			}
		}
	}

	return ret
}

func solve(ax, bx, ay, by, px, py int) (int, int) {
	det := ax*by - ay*bx
	if det == 0 {
		return -1, -1
	}

	x := (px*by - py*bx) / det
	y := (py*ax - px*ay) / det

	if (px*by-py*bx)%det != 0 || (py*ax-px*ay)%det != 0 {
		return -1, -1
	}

	return x, y
}

func part1(input []Machine) int {
	var ret int
	for _, m := range input {
		x, y := solve(m.A[0], m.B[0], m.A[1], m.B[1], m.Prize[0], m.Prize[1])

		if x < 0 || y < 0 {
			continue
		}

		if x > 100 || y > 100 {
			continue
		}

		ret += x*3 + y
	}

	return ret
}
func part2(input []Machine) int {
	var ret int

	for k := range input {
		input[k].Prize[0] += 10000000000000
		input[k].Prize[1] += 10000000000000
	}

	for _, m := range input {
		x, y := solve(m.A[0], m.B[0], m.A[1], m.B[1], m.Prize[0], m.Prize[1])

		if x < 0 || y < 0 {
			continue
		}

		ret += x*3 + y
	}

	return ret
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("part1", part1(parseInput(string(input))))
	fmt.Println("part2", part2(parseInput(string(input))))
}
