package main

import (
	"testing"
)

var tests = []struct {
	input string
	part1 int
	part2 int
}{
	{
		`Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`,
		4635635210, 64,
	},
}

func TestPartOne(t *testing.T) {
	for i, test := range []struct {
		computer        Computer
		expectRegisters map[rune]int
		expectOutput    string
	}{
		{
			Computer{
				Registers: map[rune]int{'C': 9},
				Program:   []int{2, 6},
			},
			map[rune]int{'B': 1},
			"",
		},
		{
			Computer{
				Registers: map[rune]int{'A': 10},
				Program:   []int{5, 0, 5, 1, 5, 4},
			},
			map[rune]int{},
			"0,1,2",
		},
		{
			Computer{
				Registers: map[rune]int{'A': 2024},
				Program:   []int{0, 1, 5, 4, 3, 0},
			},
			map[rune]int{'A': 0},
			"4,2,5,6,7,7,7,7,3,1,0",
		},
		{
			Computer{
				Registers: map[rune]int{'B': 29},
				Program:   []int{1, 7},
			},
			map[rune]int{'B': 26},
			"",
		},
		{
			Computer{
				Registers: map[rune]int{'B': 2024, 'C': 43690},
				Program:   []int{4, 0},
			},
			map[rune]int{'B': 44354},
			"",
		},
		{
			Computer{
				Registers: map[rune]int{'A': 729},
				Program:   []int{0, 1, 5, 4, 3, 0},
			},
			map[rune]int{},
			"4,6,3,5,6,3,5,2,1,0",
		},
	} {
		test.computer.Run()
		for reg, val := range test.expectRegisters {
			if test.computer.Registers[reg] != val {
				t.Errorf("Test #%d: expected %d, got %d", i, val, test.computer.Registers[reg])
			}
		}
		if len(test.expectOutput) > 0 {
			if test.computer.GetOutput() != test.expectOutput {
				t.Errorf("Test #%d: expected %s, got %s", i, test.expectOutput, test.computer.GetOutput())
			}
		}
	}

}

func TestPartTwo(t *testing.T) {
	program := `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`
	expected := 117440
	actual := part2(parseInput(program))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

// There are two types of operands; each instruction specifies the type of its operand. The value of a literal operand is the operand itself. For example, the value of the literal operand 7 is the number 7. The value of a combo operand can be found as follows:

// The adv instruction (opcode 0) performs division. The numerator is the value in the A register. The denominator is found by raising 2 to the power of the instruction's combo operand. (So, an operand of 2 would divide A by 4 (2^2); an operand of 5 would divide A by 2^B.) The result of the division operation is truncated to an integer and then written to the A register.
func TestAdvOperator(t *testing.T) {
	for _, test := range []struct {
		registers map[rune]int
		operand   int
		expect    int
	}{
		{map[rune]int{'A': 10, 'B': 1}, 2, 2},
		{map[rune]int{'A': 10, 'B': 2}, 5, 2},
	} {
		c := Computer{Registers: test.registers}
		c.Adv(test.operand)
		if c.Registers['A'] != test.expect {
			t.Errorf("Expected %d, got %d", test.expect, c.Registers['A'])
		}
	}
}

// The bxl instruction (opcode 1) calculates the bitwise XOR of register B and the instruction's literal operand, then stores the result in register B.
func TestBxlOperator(t *testing.T) {
	for _, test := range []struct {
		registers map[rune]int
		operand   int
		expect    int
	}{
		{map[rune]int{'A': 10, 'B': 1}, 2, 3},
		{map[rune]int{'A': 10, 'B': 2}, 5, 7},
	} {
		c := Computer{Registers: test.registers}
		c.Bxl(test.operand)
		if c.Registers['B'] != test.expect {
			t.Errorf("Expected %d, got %d", test.expect, c.Registers['B'])
		}
	}
}

// The bst instruction (opcode 2) calculates the value of its combo operand modulo 8 (thereby keeping only its lowest 3 bits), then writes that value to the B register.
func TestBstOperator(t *testing.T) {
	for _, test := range []struct {
		registers map[rune]int
		operand   int
		expect    int
	}{
		{map[rune]int{}, 10, 2},
		{map[rune]int{'A': 10}, 4, 2},
		{map[rune]int{'C': 9}, 6, 1},
	} {
		c := Computer{Registers: test.registers}
		c.Bst(test.operand)
		if c.Registers['B'] != test.expect {
			t.Errorf("Expected %d, got %d", test.expect, c.Registers['B'])
		}
	}
}

// The jnz instruction (opcode 3) does nothing if the A register is 0. However, if the A register is not zero, it jumps by setting the instruction pointer to the value of its literal operand; if this instruction jumps, the instruction pointer is not increased by 2 after this instruction.
func TestJnzOperator(t *testing.T) {
	for _, test := range []struct {
		registers map[rune]int
		operand   int
		expect    int
	}{
		{map[rune]int{'A': 0}, 10, 0},
		{map[rune]int{'A': 1}, 10, 10},
		{map[rune]int{'A': 2}, 10, 10},
	} {
		c := Computer{Registers: test.registers}
		c.Jnz(test.operand)
		if c.InstructionPointer != test.expect {
			t.Errorf("Expected %d, got %d", test.expect, c.InstructionPointer)
		}
	}
}

// The bxc instruction (opcode 4) calculates the bitwise XOR of register B and register C, then stores the result in register B. (For legacy reasons, this instruction reads an operand but ignores it.)
func TestBxcOperator(t *testing.T) {
	for _, test := range []struct {
		registers map[rune]int
		operand   int
		expect    int
	}{
		{map[rune]int{'B': 0, 'C': 1}, 10, 1},
		{map[rune]int{'B': 1, 'C': 2}, 10, 3},
	} {
		c := Computer{Registers: test.registers}
		c.Bxc(test.operand)
		if c.Registers['B'] != test.expect {
			t.Errorf("Expected %d, got %d", test.expect, c.Registers['B'])
		}
	}
}

// The out instruction (opcode 5) calculates the value of its combo operand modulo 8, then outputs that value. (If a program outputs multiple values, they are separated by commas.)
func TestOutOperator(t *testing.T) {
	for _, test := range []struct {
		registers map[rune]int
		operand   int
		expect    int
	}{
		{map[rune]int{}, 10, 2},
		{map[rune]int{'A': 10}, 4, 2},
	} {
		c := Computer{Registers: test.registers}
		c.Out(test.operand)
		if c.Output[0] != test.expect {
			t.Errorf("Expected %d, got %d", test.expect, c.Output[0])
		}
	}
}

// The bdv instruction (opcode 6) works exactly like the adv instruction except that the result is stored in the B register. (The numerator is still read from the A register.)
func TestBdvOperator(t *testing.T) {
	for _, test := range []struct {
		registers map[rune]int
		operand   int
		expect    int
	}{
		{map[rune]int{'A': 10}, 2, 2},
		{map[rune]int{'A': 10, 'B': 2}, 5, 2},
	} {
		c := Computer{Registers: test.registers}
		c.Bdv(test.operand)
		if c.Registers['B'] != test.expect {
			t.Errorf("Expected %d, got %d", test.expect, c.Registers['B'])
		}
	}
}

// The cdv instruction (opcode 7) works exactly like the adv instruction except that the result is stored in the C register. (The numerator is still read from the A register.)
func TestCdvOperator(t *testing.T) {
	for _, test := range []struct {
		registers map[rune]int
		operand   int
		expect    int
	}{
		{map[rune]int{'A': 10}, 2, 2},
		{map[rune]int{'A': 10, 'B': 2}, 5, 2},
	} {
		c := Computer{Registers: test.registers}
		c.Cdv(test.operand)
		if c.Registers['C'] != test.expect {
			t.Errorf("Expected %d, got %d", test.expect, c.Registers['C'])
		}
	}
}
