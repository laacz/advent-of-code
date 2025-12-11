package main

import (
	"fmt"
	"os"
	"strings"
)

type Expression string

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Part 1:", part1(string(data)))
	fmt.Println("Part 2:", part2(string(data)))
}

func (e Expression) Evaluate() int {
	pos := 0
	return e.evalExpr(&pos)
}

func (e Expression) evalExpr(pos *int) int {
	result := e.evalValue(pos)

	for *pos < len(e) {
		e.skipSpaces(pos)
		if *pos >= len(e) {
			break
		}

		ch := e[*pos]
		if ch == ')' {
			break
		}

		if ch == '+' || ch == '*' {
			*pos++
			e.skipSpaces(pos)
			right := e.evalValue(pos)
			if ch == '+' {
				result += right
			} else {
				result *= right
			}
		} else {
			break
		}
	}

	return result
}

func (e Expression) evalValue(pos *int) int {
	e.skipSpaces(pos)
	if *pos >= len(e) {
		return 0
	}

	ch := e[*pos]
	if ch == '(' {
		*pos++
		result := e.evalExpr(pos)
		if *pos < len(e) && e[*pos] == ')' {
			*pos++
		}
		return result
	}

	num := 0
	for *pos < len(e) && e[*pos] >= '0' && e[*pos] <= '9' {
		num = num*10 + int(e[*pos]-'0')
		*pos++
	}
	return num
}

func (e Expression) skipSpaces(pos *int) {
	for *pos < len(e) && e[*pos] == ' ' {
		*pos++
	}
}

func (e Expression) Evaluate2() int {
	pos := 0
	return e.evalExpr2(&pos)
}

func (e Expression) evalExpr2(pos *int) int {
	result := e.evalTerm(pos)

	for *pos < len(e) {
		e.skipSpaces(pos)
		if *pos >= len(e) || e[*pos] == ')' {
			break
		}

		if e[*pos] == '*' {
			*pos++
			e.skipSpaces(pos)
			right := e.evalTerm(pos)
			result *= right
		} else {
			break
		}
	}

	return result
}

func (e Expression) evalTerm(pos *int) int {
	result := e.evalValue2(pos)

	for *pos < len(e) {
		e.skipSpaces(pos)
		if *pos >= len(e) || e[*pos] == ')' || e[*pos] == '*' {
			break
		}

		if e[*pos] == '+' {
			*pos++
			e.skipSpaces(pos)
			right := e.evalValue2(pos)
			result += right
		} else {
			break
		}
	}

	return result
}

func (e Expression) evalValue2(pos *int) int {
	e.skipSpaces(pos)
	if *pos >= len(e) {
		return 0
	}

	ch := e[*pos]
	if ch == '(' {
		*pos++
		result := e.evalExpr2(pos)
		if *pos < len(e) && e[*pos] == ')' {
			*pos++
		}
		return result
	}

	num := 0
	for *pos < len(e) && e[*pos] >= '0' && e[*pos] <= '9' {
		num = num*10 + int(e[*pos]-'0')
		*pos++
	}
	return num
}

func part1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	sum := 0
	for _, line := range lines {
		sum += Expression(line).Evaluate()
	}
	return sum
}

func part2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	sum := 0
	for _, line := range lines {
		sum += Expression(line).Evaluate2()
	}
	return sum
}
