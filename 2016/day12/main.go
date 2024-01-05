package main

import (
	"fmt"
	"strings"

	"github.com/laacz/aoc-2016/util"
)

type Program struct {
	Registers    map[string]int
	Instructions []Instruction
}

func (p Program) GetRegOrVal(s string) int {
	if s[0] >= 'a' && s[0] <= 'd' {
		return p.Registers[s]
	}
	return util.Atoi(s)
}

func (p *Program) Execute() {
	i := 0
	for i < len(p.Instructions) {
		instr := p.Instructions[i]

		switch instr.Type {
		case "cpy":
			val := p.GetRegOrVal(instr.Args[0])
			p.Registers[instr.Args[1]] = val
		case "inc":
			p.Registers[instr.Args[0]]++
		case "dec":
			p.Registers[instr.Args[0]]--
		case "jnz":
			val := p.GetRegOrVal(instr.Args[0])
			if val != 0 {
				i += util.Atoi(instr.Args[1])
				continue
			}
		}

		i += 1
	}
}

type Instruction struct {
	Type string
	Args []string
}

func parse(input string) (ret Program) {
	for _, line := range util.GetLines(input) {
		var instr Instruction
		var parts = strings.Split(line, " ")

		instr.Type = parts[0]
		instr.Args = parts[1:]

		ret.Instructions = append(ret.Instructions, instr)
	}
	ret.Registers = map[string]int{"a": 0, "b": 0, "c": 0, "d": 0}

	return
}

func partOne(input string) (ret int) {
	p := parse(input)
	p.Execute()

	return p.Registers["a"]
}

func partTwo(input string) (ret int) {
	p := parse(input)
	p.Registers["c"] = 1
	p.Execute()

	return p.Registers["a"]
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}
