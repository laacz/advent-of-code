package main

import (
	"fmt"
	"os"
	"strings"
)

type Instruction struct {
	Op   string
	Arg1 string
	Arg2 string
}

type Program []Instruction

type Registers map[string]int

func getInt(s string) int {
	if s == "" {
		return 0
	}

	if s[0] == '+' {
		s = s[1:]
	}

	ret := 0
	fmt.Sscanf(s, "%d", &ret)
	return ret
}

func (p Program) Run(regs Registers) Registers {
	pos := 0

	for {
		if pos >= len(p) {
			break
		}
		switch p[pos].Op {
		case "hlf":
			regs[p[pos].Arg1] /= 2
		case "tpl":
			regs[p[pos].Arg1] *= 3
		case "inc":
			regs[p[pos].Arg1] += 1
		case "jmp":
			pos += getInt(p[pos].Arg1)
			continue
		case "jie":
			if regs[p[pos].Arg1]%2 == 0 {
				pos += getInt(p[pos].Arg2)
				continue
			}
		case "jio":
			if regs[p[pos].Arg1] == 1 {
				pos += getInt(p[pos].Arg2)
				continue
			}

		}

		pos += 1
	}

	return regs
}

func parse(input string) (ret Program) {
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		op := line[:3]
		arg1 := ""
		arg2 := ""

		args := strings.Split(line[4:], ",")
		if len(args) > 0 {
			arg1 = strings.TrimSpace(args[0])
		}

		if len(args) > 1 {
			arg2 = strings.TrimSpace(args[1])
		}

		ret = append(ret, Instruction{op, arg1, arg2})
	}

	return ret
}

func partOne(lines string) (ret int) {
	i := parse(lines)
	q := i.Run(map[string]int{"a": 0, "b": 0})

	return q["b"]
}

func partTwo(lines string) (ret int) {
	i := parse(lines)
	q := i.Run(map[string]int{"a": 1, "b": 0})

	return q["b"]
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
