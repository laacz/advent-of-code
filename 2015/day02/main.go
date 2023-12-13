package main

import (
	"fmt"
	"os"
	"strings"
)

func partOne(lines string) int {
	sum := 0

	for _, c := range strings.Split(lines, "\n") {
		if len(c) == 0 {
			continue
		}

		var l, w, h int
		fmt.Sscanf(c, "%dx%dx%d", &l, &w, &h)

		area := 2*l*w + 2*w*h + 2*h*l
		area += min(l*w, min(w*h, h*l))

		sum += area
	}

	return sum
}

func partTwo(lines string) int {
	sum := 0

	for _, c := range strings.Split(lines, "\n") {
		if len(c) == 0 {
			continue
		}

		var l, w, h int
		fmt.Sscanf(c, "%dx%dx%d", &l, &w, &h)

		ribbon := min(2*l+2*w, min(2*w+2*h, 2*h+2*l))
		ribbon += l * w * h

		sum += ribbon
	}

	return sum
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
