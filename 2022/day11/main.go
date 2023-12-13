package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var monkeys []Monkey

type Monkey struct {
	Name        int
	Items       []int
	Operation   string
	DivisableBy int
	TargetTrue  int
	TargetFalse int
	Inspections int
}

var lcm int

func (m *Monkey) Turn() {
	for _, item := range m.Items {
		m.Inspections += 1
		item = m.Operate(item)
		item = item / 3
		if item%m.DivisableBy == 0 {
			monkeys[m.TargetTrue].Items = append(monkeys[m.TargetTrue].Items, item)
		} else {
			monkeys[m.TargetFalse].Items = append(monkeys[m.TargetFalse].Items, item)
		}
	}
	m.Items = []int{}
}

func (m *Monkey) Turn2() {
	for _, item := range m.Items {
		m.Inspections += 1

		item = m.Operate(item)
		isDivisableBy := item%m.DivisableBy == 0
		if item > lcm {
			item = lcm + item%lcm
		}
		if isDivisableBy {
			monkeys[m.TargetTrue].Items = append(monkeys[m.TargetTrue].Items, item)
		} else {
			monkeys[m.TargetFalse].Items = append(monkeys[m.TargetFalse].Items, item)
		}
	}
	m.Items = []int{}
}

func (m *Monkey) Operate(item int) int {
	var a int
	var b int

	parts := strings.Split(m.Operation, " ")

	if parts[0] == "old" {
		a = item
	} else {
		a, _ = strconv.Atoi(parts[0])
	}

	if parts[2] == "old" {
		b = item
	} else {
		b, _ = strconv.Atoi(parts[2])
	}

	switch parts[1] {
	case "+":
		item = a + b
	case "*":
		item = a * b
	default:
		panic("Unknown operation " + parts[2])
	}

	return item
}

func parseInput(input string) []Monkey {
	var ret []Monkey
	lcm = 1
	for _, str := range strings.Split(input, "\n\n") {
		if strings.TrimSpace(str) == "" {
			continue
		}
		m := Monkey{}
		for _, line := range strings.Split(str, "\n") {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "Starting items:") {
				for _, item := range strings.Split(line[15:], ", ") {
					i, _ := strconv.Atoi(strings.TrimSpace(item))
					m.Items = append(m.Items, i)
				}
			}

			if strings.HasPrefix(line, "Monkey ") {
				m.Name, _ = strconv.Atoi(line[7:8])
			}

			if strings.HasPrefix(line, "Operation:") {
				m.Operation = line[17:]
			}

			if strings.HasPrefix(line, "Test: divisible by") {
				m.DivisableBy, _ = strconv.Atoi(line[19:])
				lcm *= m.DivisableBy
			}

			if strings.HasPrefix(line, "If true: throw to monkey") {
				m.TargetTrue, _ = strconv.Atoi(line[25:])
			}

			if strings.HasPrefix(line, "If false: throw to monkey") {
				m.TargetFalse, _ = strconv.Atoi(line[26:])
			}

		}
		ret = append(ret, m)

	}
	return ret
}

func partOne(input string) int {
	monkeys = parseInput(input)
	i := 0
	for i < 20 {
		for i := range monkeys {
			monkeys[i].Turn()
		}
		i += 1
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Inspections > monkeys[j].Inspections
	})

	return monkeys[0].Inspections * monkeys[1].Inspections
}

func partTwo(input string) int {
	monkeys = parseInput(input)

	i := 0
	for i < 10000 {
		for i := range monkeys {
			monkeys[i].Turn2()
		}
		i += 1
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Inspections > monkeys[j].Inspections
	})

	return monkeys[0].Inspections * monkeys[1].Inspections
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
