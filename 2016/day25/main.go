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

func isRegister(s string) bool {
	return s[0] >= 'a' && s[0] <= 'd'
}

func (p Program) GetRegOrVal(s string) int {
	if isRegister(s) {
		return p.Registers[s]
	}
	return util.Atoi(s)
}

// hint from teh internets is that we can replace this pattern:
// > cpy b c / inc a / dec c / jnz c -2 / dec d / jnz d -5
// with
// > a += b * d, c = 0, d = 0
func (p *Program) optimize(i int) bool {
	if i+5 >= len(p.Instructions) {
		return false
	}

	i0, i1, i2, i3, i4, i5 := p.Instructions[i], p.Instructions[i+1], p.Instructions[i+2], p.Instructions[i+3], p.Instructions[i+4], p.Instructions[i+5]

	if i0.Type != "cpy" || i1.Type != "inc" || i2.Type != "dec" || i3.Type != "jnz" || i4.Type != "dec" || i5.Type != "jnz" {
		return false
	}
	if i3.Args[1] != "-2" || i5.Args[1] != "-5" {
		return false
	}
	if i0.Args[1] != i2.Args[0] || i0.Args[1] != i3.Args[0] { // c
		return false
	}
	if i4.Args[0] != i5.Args[0] { // d
		return false
	}
	if !isRegister(i1.Args[0]) || !isRegister(i0.Args[1]) || !isRegister(i4.Args[0]) {
		return false
	}

	p.Registers[i1.Args[0]] += p.GetRegOrVal(i0.Args[0]) * p.Registers[i4.Args[0]]
	p.Registers[i0.Args[1]] = 0
	p.Registers[i4.Args[0]] = 0
	return true
}

// Execute runs the program and returns true if it produces valid clock signal
// (alternating 0, 1, 0, 1...) for enough outputs to be confident
func (p *Program) Execute() bool {
	i := 0
	expectedOutput := 0
	outputCount := 0
	maxOutputs := 100 // Check enough outputs to be confident

	for i < len(p.Instructions) {
		if p.optimize(i) {
			i += 6
			continue
		}

		instr := p.Instructions[i]

		switch instr.Type {
		case "cpy":
			if !isRegister(instr.Args[1]) {
				break // invalid instruction, skip
			}
			val := p.GetRegOrVal(instr.Args[0])
			p.Registers[instr.Args[1]] = val
		case "inc":
			if !isRegister(instr.Args[0]) {
				break // invalid instruction, skip
			}
			p.Registers[instr.Args[0]]++
		case "dec":
			if !isRegister(instr.Args[0]) {
				break // invalid instruction, skip
			}
			p.Registers[instr.Args[0]]--
		case "jnz":
			val := p.GetRegOrVal(instr.Args[0])
			if val != 0 {
				i += p.GetRegOrVal(instr.Args[1])
				continue
			}
		case "tgl":
			offset := p.GetRegOrVal(instr.Args[0])
			idx := i + offset
			if idx < 0 || idx >= len(p.Instructions) {
				break
			}
			targetInstr := p.Instructions[idx]
			switch targetInstr.Type {
			case "inc":
				targetInstr.Type = "dec"
			case "dec", "tgl":
				targetInstr.Type = "inc"
			case "jnz":
				targetInstr.Type = "cpy"
			case "cpy":
				targetInstr.Type = "jnz"
			}
			p.Instructions[idx] = targetInstr
		case "out":
			val := p.GetRegOrVal(instr.Args[0])
			if val != expectedOutput {
				return false // Wrong output, not a valid clock signal
			}
			expectedOutput = 1 - expectedOutput // Toggle between 0 and 1
			outputCount++
			if outputCount >= maxOutputs {
				return true // Enough correct outputs, assume it's valid
			}
		}

		i += 1
	}
	return false
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
	for a := 1; ; a++ {
		p := parse(input)
		p.Registers["a"] = a
		if p.Execute() {
			return a
		}
	}
}

func partTwo(input string) (ret int) {
	// No part two for day 25 - just need all 50 stars
	return 0
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println("Part one:", partOne(input))
}
