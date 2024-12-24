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

func (d Device) String() string {
	var ret string

	keys := make([]string, 0, len(d.io))
	for k := range d.io {
		if k[0] != 'z' {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		ret += fmt.Sprintf("%d", d.io[k])
	}

	// reverse
	runes := []rune(ret)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
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
		ret.program = append(ret.program, Gate{
			inputs:    [2]string{parts[0], parts[2]},
			output:    parts[4],
			operation: parts[1],
		})
	}

	return ret
}

func part1(d Device) int {
	for d.Simulate() {
	}

	state := d.String()
	ret, _ := strconv.ParseInt(state, 2, 64)
	return int(ret)
}

func part2(d Device) int {
	var ret int
	return ret
}

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Println("part1", part1(parseInput(string(input))))
	fmt.Println("part2", part2(parseInput(string(input))))
}
