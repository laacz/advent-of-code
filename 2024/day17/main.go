package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Computer struct {
	Registers          map[rune]int
	Program            []int
	InstructionPointer int
	Output             []int
}

func (c *Computer) String() string {
	ret := "Computer registers [ "
	for _, reg := range []rune{'A', 'B', 'C'} {
		ret += fmt.Sprintf("%c: %5d ", reg, c.Registers[reg])
	}
	ret += "] "

	ret += "output: [ "
	for _, o := range c.Output {
		ret += fmt.Sprintf("%d ", o)
	}
	ret += "]"

	return ret
}

func parseInput(input string) Computer {
	var ret Computer

	ret.Registers = map[rune]int{
		'A': 0,
		'B': 0,
		'C': 0,
	}

	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	for _, line := range strings.Split(parts[0], "\n") {
		var r rune
		var v int
		fmt.Sscanf(line, "Register %c: %d", &r, &v)
		ret.Registers[r] = v
	}

	for _, line := range strings.Split(parts[1], "\n") {
		parts := strings.Split(line[9:], ",")
		for i := 0; i < len(parts); i += 2 {
			ret.Program = append(ret.Program, int(parts[i][0]-'0'), int(parts[i+1][0]-'0'))
		}
	}

	return ret
}

func (c *Computer) ComboOperand(n int) int {
	switch n {
	case 4:
		return c.Registers['A']
	case 5:
		return c.Registers['B']
	case 6:
		return c.Registers['C']
	}

	return n
}

// returns comma joined output
func (c *Computer) GetOutput() string {
	ret := ""
	for i, o := range c.Output {
		if i > 0 {
			ret += ","
		}
		ret += strconv.Itoa(o)
	}
	ret = strings.TrimRight(ret, ",")
	return ret
}

func (c *Computer) GetProgram() string {
	ret := ""
	for i, o := range c.Program {
		if i > 0 {
			ret += ","
		}
		ret += strconv.Itoa(o)
	}
	ret = strings.TrimRight(ret, ",")
	return ret
}

// Adv The adv instruction (opcode 0) performs division. The numerator is
// the value in the A register. The denominator is found by raising 2 to the
// power of the instruction's combo operand. (So, an operand of 2 would
// divide A by 4 (2^2); an operand of 5 would divide A by 2^B.) The result
// of the division operation is truncated to an integer and then written to
// the A register.
func (c *Computer) Adv(operand int) {
	c.Registers['A'] /= 1 << c.ComboOperand(operand)
}

// Bxl The bxl instruction (opcode 1) calculates the bitwise XOR of register B and
// the instruction's literal operand, then stores the result in register B.
func (c *Computer) Bxl(operand int) {
	c.Registers['B'] ^= operand
}

// Bst The bst instruction (opcode 2) calculates the value of its combo operand modulo
// 8 (thereby keeping only its lowest 3 bits), then writes that value to the B register.
func (c *Computer) Bst(operand int) {
	c.Registers['B'] = c.ComboOperand(operand) % 8
}

// Jnz The jnz instruction (opcode 3) does nothing if the A register is 0. However, if the
// A register is not zero, it jumps by setting the instruction pointer to the value
// of its literal operand; if this instruction jumps, the instruction pointer is not
// increased by 2 after this instruction.
func (c *Computer) Jnz(operand int) {
	if c.Registers['A'] != 0 {
		c.InstructionPointer = operand
	}
}

// Bxc The bxc instruction (opcode 4) calculates the bitwise XOR of register B and
// register C, then stores the result in register B. (For legacy reasons, this
// instruction reads an operand but ignores it.)
func (c *Computer) Bxc(operand int) {
	c.Registers['B'] ^= c.Registers['C']
}

// Out The out instruction (opcode 5) calculates the value of its combo operand modulo 8,
// then outputs that value. (If a program outputs multiple values, they are separated
// by commas.)
func (c *Computer) Out(operand int) {
	c.Output = append(c.Output, c.ComboOperand(operand)%8)
}

// Bdv The bdv instruction (opcode 6) works exactly like the adv instruction except that
// the result is stored in the B register. (The numerator is still read from the A
// register.)
func (c *Computer) Bdv(operand int) {
	c.Registers['B'] = c.Registers['A'] / (1 << c.ComboOperand(operand))
}

// Cdv The cdv instruction (opcode 7) works exactly like the adv instruction except that
// the result is stored in the C register. (The numerator is still read from the A
// register.)
func (c *Computer) Cdv(operand int) {
	c.Registers['C'] = c.Registers['A'] / (1 << c.ComboOperand(operand))
}

func (c *Computer) Step() bool {
	if c.InstructionPointer >= len(c.Program) {
		return false
	}

	opcode := c.Program[c.InstructionPointer]
	operand := c.Program[c.InstructionPointer+1]

	c.InstructionPointer += 2

	switch opcode {
	case 0:
		c.Adv(operand)
	case 1:
		c.Bxl(operand)
	case 2:
		c.Bst(operand)
	case 3:
		c.Jnz(operand)
	case 4:
		c.Bxc(operand)
	case 5:
		c.Out(operand)
	case 6:
		c.Bdv(operand)
	case 7:
		c.Cdv(operand)
	}

	return true
}

func (c *Computer) Run() {
	for c.Step() {
	}
}

func part1(comp Computer) string {
	comp.Run()

	return comp.GetOutput()
}

func part2(comp Computer) int {
	var ret int
	for i := len(comp.Program) - 1; i >= 0; i-- {
		ret <<= 3
		for {
			c := Computer{
				Registers: map[rune]int{'A': ret},
				Program:   comp.Program,
			}
			c.Run()
			if slices.Equal(c.Output, c.Program[i:]) {
				break
			}
			ret++
		}
	}

	return ret
}

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Println("part1", part1(parseInput(string(input))))
	fmt.Println("part2", part2(parseInput(string(input))))
}
