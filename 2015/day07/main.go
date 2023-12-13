package main

import (
	"fmt"
	"os"
	"strings"
)

type Registers map[string]uint16

func (registers Registers) getVal(s string) uint16 {
	var ret uint16
	if _, err := fmt.Sscanf(s, "%d", &ret); err == nil {
		return ret
	}

	return registers[s]
}
func getInt(s string) uint16 {
	var ret uint16
	fmt.Sscanf(s, "%d", &ret)
	return ret
}

func isInt(s string) bool {
	_, err := fmt.Sscanf(s, "%d", new(int))
	return err == nil
}

func partOne(lines string) int {
	registers := make(Registers)
	round := 0
	for {
		round += 1
		for _, line := range strings.Split(lines, "\n") {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			parts := strings.Split(line, " -> ")
			target := parts[1]

			msg := ""
			op := strings.Split(parts[0], " ")
			if len(op) == 1 {
				if _, ok := registers[op[0]]; !isInt(op[0]) && !ok {
					continue
				}
				val := getInt(op[0])
				if _, err := fmt.Sscanf(op[0], "%d", &val); err != nil {
					val = registers[op[0]]
				}
				registers[target] = val

				msg = fmt.Sprintf("%s (%d) -> %s (%d)", op[0], val, target, registers[target])
			} else if len(op) == 2 {
				if _, ok := registers[op[1]]; !isInt(op[0]) && !ok {
					continue
				}
				val := registers.getVal(op[1])
				registers[target] = ^val

				msg = fmt.Sprintf("%s %s (%d) -> %s (%d)", op[0], op[1], val, target, registers[target])
			} else if len(op) == 3 {
				if _, ok := registers[op[0]]; !isInt(op[0]) && !ok {
					continue
				}
				if _, ok := registers[op[2]]; !isInt(op[2]) && !ok {
					continue
				}
				val1 := registers.getVal(op[0])
				val2 := registers.getVal(op[2])
				switch op[1] {
				case "AND":
					registers[target] = val1 & val2
				case "OR":
					registers[target] = val1 | val2
				case "LSHIFT":
					registers[target] = val1 << val2
				case "RSHIFT":
					registers[target] = val1 >> val2
				}

				msg = fmt.Sprintf("%s (%d) %s %s (%d) -> %s (%d)", op[0], val1, op[1], op[2], val2, target, registers[target])
			}

			if false {
				fmt.Printf("%d \033[1;33m%s\033[0m\n", round, msg)
			}

			if target == "a" {
				// fmt.Println(registers)
				return int(registers["a"])
			}
		}
	}
}

func partTwo(lines string, p1 int) int {
	registers := make(Registers)
	round := 0

	for {
		round += 1
		for _, line := range strings.Split(lines, "\n") {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			parts := strings.Split(line, " -> ")
			target := parts[1]

			msg := ""
			op := strings.Split(parts[0], " ")
			if len(op) == 1 {
				if _, ok := registers[op[0]]; !isInt(op[0]) && !ok {
					continue
				}
				val := getInt(op[0])
				if _, err := fmt.Sscanf(op[0], "%d", &val); err != nil {
					val = registers[op[0]]
				}
				if target == "b" {
					val = uint16(p1)
				}
				registers[target] = val

				msg = fmt.Sprintf("%s (%d) -> %s (%d)", op[0], val, target, registers[target])
			} else if len(op) == 2 {
				if _, ok := registers[op[1]]; !isInt(op[0]) && !ok {
					continue
				}
				val := registers.getVal(op[1])
				registers[target] = ^val

				msg = fmt.Sprintf("%s %s (%d) -> %s (%d)", op[0], op[1], val, target, registers[target])
			} else if len(op) == 3 {
				if _, ok := registers[op[0]]; !isInt(op[0]) && !ok {
					continue
				}
				if _, ok := registers[op[2]]; !isInt(op[2]) && !ok {
					continue
				}
				val1 := registers.getVal(op[0])
				val2 := registers.getVal(op[2])
				switch op[1] {
				case "AND":
					registers[target] = val1 & val2
				case "OR":
					registers[target] = val1 | val2
				case "LSHIFT":
					registers[target] = val1 << val2
				case "RSHIFT":
					registers[target] = val1 >> val2
				}

				msg = fmt.Sprintf("%s (%d) %s %s (%d) -> %s (%d)", op[0], val1, op[1], op[2], val2, target, registers[target])
			}

			if false {
				fmt.Printf("%d \033[1;33m%s\033[0m\n", round, msg)
			}

			if target == "a" {
				// fmt.Println(registers)
				return int(registers["a"])
			}
		}
	}
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data), partOne(string(data))))
}
