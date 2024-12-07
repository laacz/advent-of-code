package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	Result   int
	Operands []int
}

func (e Equation) SolvableBinary() bool {
	var n = len(e.Operands) - 1

	for i := 0; i < (1 << n); i++ {
		result := e.Operands[0]
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				result += e.Operands[j+1]
			} else {
				result *= e.Operands[j+1]
			}
		}

		if result == e.Result {
			return true
		}
	}

	return false
}

func (e Equation) SolvableTernary() bool {
	n := len(e.Operands) - 1

	for i := 0; i < int(math.Pow(3, float64(n))); i++ {
		temp := i
		result := e.Operands[0]

		for j := 0; j < n; j++ {
			op := temp % 3
			temp /= 3

			switch op {
			case 0:
				result += e.Operands[j+1]
			case 1:
				result *= e.Operands[j+1]
			case 2:
				result, _ = strconv.Atoi(fmt.Sprintf("%d%d", result, e.Operands[j+1]))
			}
		}

		if result == e.Result {
			return true
		}
	}
	return false
}

func parseInput(input string) []Equation {
	var ret []Equation
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		var eq Equation
		parts := strings.Split(line, ": ")
		eq.Result, _ = strconv.Atoi(parts[0])

		for _, num := range strings.Fields(parts[1]) {
			n, _ := strconv.Atoi(num)
			eq.Operands = append(eq.Operands, n)
		}

		ret = append(ret, eq)
	}
	return ret
}

func part1(input string) int {
	var ret int

	equations := parseInput(input)
	for _, eq := range equations {
		if eq.SolvableBinary() {
			ret += eq.Result
		}
	}

	return ret
}

func part2(input string) int {
	var ret int
	equations := parseInput(input)
	for _, eq := range equations {
		if eq.SolvableTernary() {
			ret += eq.Result
		}
	}

	return ret
}
func main() {

	input, _ := os.ReadFile("input.txt")

	fmt.Println("part1", part1(string(input)))
	fmt.Println("part2", part2(string(input)))
}
