package main

import (
	"fmt"
	"os"
	"strings"
)

type Aunt struct {
	props map[string]int
}

func parse(line string) map[int]Aunt {
	ret := make(map[int]Aunt)

	for _, line := range strings.Split(line, "\n") {
		aunt := Aunt{
			props: make(map[string]int),
		}

		var id int
		parts := strings.Split(line, ":")
		fmt.Sscanf(parts[0], "Sue %d", &id)

		for _, part := range strings.Split(strings.Join(parts[1:], ""), ",") {
			var prop string
			var val int
			fmt.Sscanf(strings.TrimSpace(part), "%s %d", &prop, &val)
			aunt.props[prop] = val
		}

		ret[id] = aunt
	}

	return ret
}

var MFCSAM = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func partOne(lines string) (ret int) {
	aunts := parse(lines)

	for id, aunt := range aunts {
		match := true
		for prop, val := range aunt.props {
			if MFCSAM[prop] != val {
				// fmt.Println("aunt", id, "doesn't match", prop, val, "tickertape had", MFCSAM[prop])
				match = false
				break
			}
		}

		if match {
			return id
		}
	}

	return ret
}

func partTwo(lines string) (ret int) {
	aunts := parse(lines)

	for id, aunt := range aunts {
		match := true
		for prop, val := range aunt.props {
			switch prop {
			case "cats", "trees":
				if val <= MFCSAM[prop] {
					match = false
				}
			case "pomeranians", "goldfish":
				if val >= MFCSAM[prop] {
					match = false
				}
			default:
				if val != MFCSAM[prop] {
					match = false
				}
			}
		}

		if match {
			return id
		}
	}

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
