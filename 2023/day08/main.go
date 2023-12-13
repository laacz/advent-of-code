package main

import (
	"fmt"
	"os"
	"strings"
)

// parse (you won't believe it!) returns parsed input
func parse(input string) (string, map[string][]string) {
	i := ""
	m := make(map[string][]string)

	for k, line := range strings.Split(input, "\n") {
		if k == 0 {
			i = line
		} else if strings.Contains(line, " = ") {
			m[line[:3]] = []string{
				line[7:10],
				line[12:15],
			}
		}
	}

	return i, m
}

// partOne returns the answer to part one of this day's puzzle.
func partOne(input string) int {
	sum := 0

	i, m := parse(input)

	location := m["AAA"]
	var nextloc string
	for {
		dir := i[sum%len(i)]
		if dir == 'R' {
			nextloc = location[1]
		} else {
			nextloc = location[0]
		}
		sum++

		if nextloc == "ZZZ" {
			break
		}

		location = m[nextloc]
	}

	return sum
}

// gcd returns the greatest common divisor of two integers
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// partTwo returns the answer to part one of this day's puzzle.
func partTwo(input string) int {
	ret := 1

	i, m := parse(input)

	locations := []string{}

	for k := range m {
		if k[2] == 'A' {
			locations = append(locations, k)
		}
	}

	steps := []int{}

	for _, loc := range locations {
		step := 0
		location := m[loc]
		var nextloc string
		for {
			dir := i[step%len(i)]
			if dir == 'R' {
				nextloc = location[1]
			} else {
				nextloc = location[0]
			}
			step++

			if nextloc[2] == 'Z' {
				break
			}

			location = m[nextloc]
		}

		steps = append(steps, step)

	}
	fmt.Println(locations, steps)

	// least common multiple is the answer
	for _, step := range steps {
		ret = ret * step / gcd(ret, step)
	}

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
