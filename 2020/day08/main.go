package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Op  string
	Arg int
}

type Program []Instruction

func (p Program) Run(errOnLoop bool) (int, error) {
	acc := 0
	pos := 0
	visited := make(map[int]bool)

	for pos < len(p) {
		inst := p[pos]
		if visited[pos] {
			if errOnLoop {
				return 0, fmt.Errorf("loop detected")
			}
			return acc, nil
		}
		visited[pos] = true

		switch inst.Op {
		case "acc":
			acc += inst.Arg
		case "jmp":
			pos += inst.Arg % len(p)
			continue
		case "nop":
		}

		pos++
	}

	return acc, nil
}

func parse(input string) Program {
	ret := Program{}

	for line := range strings.Lines(strings.TrimSpace(input)) {
		parts := strings.Fields(line)
		num, _ := strconv.Atoi(parts[1])

		ret = append(ret, Instruction{
			Op:  parts[0],
			Arg: num,
		})
	}

	return ret
}

func part1(input Program) int {
	ret, _ := input.Run(false)
	return ret
}

func part2(input Program) int {
	for pos, inst := range input {
		input2 := make(Program, len(input))
		copy(input2, input)

		switch inst.Op {
		case "nop":
			input2[pos].Op = "jmp"
		case "jmp":
			input2[pos].Op = "nop"
		default:
			continue
		}

		acc, err := input2.Run(true)
		if err != nil {
			continue
		}

		return acc
	}

	return 0
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Println("Part 1:", part1(parse(string(data))))
	fmt.Println("Part 2:", part2(parse(string(data))))
}
