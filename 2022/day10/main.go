package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partOne(input string) int {
	result := 0
	cycle := 1
	x := 1
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		op := line[:4]
		arg := 0
		dc := 1
		x1 := x
		if op == "addx" {
			arg, _ = strconv.Atoi(line[5:])
			dc = 2
			x1 = x + arg
		}

		if (cycle+19)/40 != (cycle+19+dc)/40 {
			result += ((cycle+19+dc)/40*40 - 20) * x
		}
		cycle += dc
		x = x1
	}

	return result
}

func cycle(x int, c *int, crt *string) {
	result := "."
	*c = *c % 40
	if *c <= x+1 && *c >= x-1 {
		result = "#"
	}

	if *c == 39 {
		result += "\n"
	}

	*c += 1
	*crt += result
}

func partTwo(input string) string {
	result := ""
	x := 1
	c := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		op := line[:4]
		arg := 0

		if op == "addx" {
			arg, _ = strconv.Atoi(line[5:])
			cycle(x, &c, &result)
			cycle(x, &c, &result)
			x = x + arg
		} else {
			cycle(x, &c, &result)
		}
	}

	return strings.Trim(result, "\n")
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: \n\n%s\n", partTwo(string(data)))
}
