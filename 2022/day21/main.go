package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	hasValue                   bool
	parentName, name, a, b, op string
	value                      int
	reverse                    bool
}

type Monkeys map[string]*Monkey

func NewMonkeys(input string) Monkeys {
	ret := make(Monkeys)
	for _, line := range strings.Split(input, "\n") {
		var name string
		var val int
		name = line[:4]
		m := &Monkey{name: name}
		val, err := strconv.Atoi(line[6:])
		if err != nil {
			m.a = line[6:10]
			m.op = line[11:12]
			m.b = line[13:17]
		} else {
			m.hasValue = true
			m.value = val
		}

		ret[name] = m
	}

	for _, monkey := range ret {
		if monkey.hasValue {
			continue
		}
		ret[monkey.a].parentName = monkey.name
		ret[monkey.b].parentName = monkey.name
	}

	ret["humn"].reverse = true
	for m := ret["humn"]; m.parentName != ""; m = ret[m.parentName] {
		ret[m.parentName].reverse = true
	}

	return ret
}

func (mm *Monkeys) isSolved() bool {
	for _, monkey := range *mm {
		if !monkey.hasValue {
			return false
		}
	}

	return true
}

func (mm *Monkeys) Solve() int {
	for !mm.isSolved() {
		for _, monkey := range *mm {
			if monkey.hasValue {
				continue
			}

			a := (*mm)[monkey.a]
			b := (*mm)[monkey.b]

			if a.hasValue && b.hasValue {
				switch monkey.op {
				case "+":
					monkey.value = a.value + b.value
				case "-":
					monkey.value = a.value - b.value
				case "*":
					monkey.value = a.value * b.value
				case "/":
					monkey.value = a.value / b.value
				}
				monkey.hasValue = true
			}
		}
	}

	return (*mm)["root"].value
}

func (mm *Monkeys) SolveReversed(m *Monkey, v int) {
	m.value = v
	m.reverse = false

	a := (*mm)[m.a]
	b := (*mm)[m.b]

	if a == nil || b == nil {
		return
	}

	var newv int
	var newm *Monkey
	if a.reverse {
		// left side is simple
		newm = a
		switch m.op {
		case "+":
			newv = v - b.value
		case "-":
			newv = v + b.value
		case "*":
			newv = v / b.value
		case "/":
			newv = v * b.value
		}
	} else if b.reverse {
		// right side is not that simple
		newm = b
		switch m.op {
		case "+":
			newv = v - a.value
		case "-":
			newv = a.value - v
		case "*":
			newv = v / a.value
		case "/":
			newv = a.value / v
		}

	}

	mm.SolveReversed(newm, newv)
}

func partOne(input string) int {
	monkeys := NewMonkeys(input)
	return monkeys.Solve()
}

func partTwo(input string) int {
	monkeys := NewMonkeys(input)
	monkeys.Solve()

	monkeys.SolveReversed(
		monkeys[monkeys["root"].a],
		(*monkeys[monkeys["root"].b]).value,
	)

	return monkeys["humn"].value
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
