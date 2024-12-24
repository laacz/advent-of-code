package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Gate struct {
	inputs    [2]string
	output    string
	operation string
}

type Device struct {
	io      map[string]int
	program []Gate
}

func (d Device) String(prefix rune) string {
	var ret string

	keys := make([]string, 0, len(d.io))
	for k := range d.io {
		if k[0] != byte(prefix) {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i := range len(keys) {
		ret += fmt.Sprintf("%d", d.io[keys[len(keys)-i-1]])
	}

	return ret
}

func (d *Device) Simulate() bool {
	waiting := false
	for _, gate := range d.program {
		var a, b int
		var ok bool

		if a, ok = d.io[gate.inputs[0]]; !ok {
			waiting = true
			continue
		}

		if b, ok = d.io[gate.inputs[1]]; !ok {
			waiting = true
			continue
		}

		switch gate.operation {
		case "AND":
			d.io[gate.output] = a & b
		case "XOR":
			d.io[gate.output] = a ^ b
		case "OR":
			d.io[gate.output] = a | b
		default:
			panic("unknown operation")
		}
	}

	return waiting
}

func parseInput(input string) Device {
	ret := Device{
		io: make(map[string]int),
	}

	parts := strings.Split(strings.TrimSpace(input), "\n\n")

	for _, values := range strings.Split(parts[0], "\n") {
		parts := strings.Split(values, ": ")
		num, _ := strconv.Atoi(parts[1])
		ret.io[parts[0]] = num
	}

	for _, values := range strings.Split(parts[1], "\n") {
		parts := strings.Split(values, " ")
		if parts[0] > parts[2] {
			parts[0], parts[2] = parts[2], parts[0]
		}
		ret.program = append(ret.program, Gate{
			inputs:    [2]string{parts[0], parts[2]},
			output:    parts[4],
			operation: parts[1],
		})
	}

	return ret
}

func (d *Device) IsGateUsed(gate, operation string) bool {
	for _, g := range d.program {
		if g.operation == operation &&
			(g.inputs[0] == gate || g.inputs[1] == gate) {
			return true
		}
	}
	return false
}

func (d *Device) FindAndFix() []string {
	var answer []string

	for _, gate := range d.program {
		if gate.operation == "AND" &&
			gate.inputs[0][0] == 'x' &&
			gate.inputs[1][0] == 'y' &&
			gate.output[0] == 'z' &&
			gate.output != "z00" {
			answer = append(answer, gate.output)
			continue
		}

		if gate.operation == "XOR" &&
			gate.inputs[0][0] != 'x' &&
			gate.inputs[1][0] != 'y' &&
			gate.output[0] != 'z' {
			answer = append(answer, gate.output)
			continue
		}

		if gate.output[0] == 'z' &&
			gate.operation != "XOR" &&
			gate.output != "z45" {
			answer = append(answer, gate.output)
			continue
		}

		if gate.operation == "AND" &&
			gate.inputs[0][0] == 'x' &&
			gate.inputs[1][0] == 'y' &&
			gate.inputs[0] != "x00" {
			if !d.IsGateUsed(gate.output, "OR") {
				answer = append(answer, gate.output)
				continue
			}
		}

		if gate.operation == "XOR" &&
			gate.inputs[0][0] == 'x' &&
			gate.inputs[1][0] == 'y' &&
			gate.inputs[0] != "x00" {
			if !d.IsGateUsed(gate.output, "AND") {
				answer = append(answer, gate.output)
			}
		}
	}

	sort.Strings(answer)
	return answer
}

func part1(d Device) int {
	for d.Simulate() {
	}

	state := d.String('z')
	ret, _ := strconv.ParseInt(state, 2, 64)
	return int(ret)
}

func part2(d Device) string {
	return strings.Join(d.FindAndFix(), ",")
}

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Println("part1", part1(parseInput(string(input))))
	fmt.Println("part2", part2(parseInput(string(input))))
}
