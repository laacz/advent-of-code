package main

import (
	"fmt"
	"os"
	"strings"
)

type Range struct {
	start, end int
}

func (r *Range) Includes(ra Range) bool {
	return (r.start >= ra.start && r.end <= ra.end) ||
		(ra.start >= r.start && ra.end <= r.end)
}

func (r *Range) Overlaps(ra Range) bool {
	return (r.start >= ra.start && r.start <= ra.end) ||
		(r.end >= ra.start && r.end <= ra.end) ||
		(ra.start >= r.start && ra.start <= r.end) ||
		(ra.end >= r.start && ra.end <= r.end)
}

func rangePairFromString(r string) (Range, Range) {
	ranges := strings.Split(r, ",")
	var r1 Range
	var r2 Range

	fmt.Sscanf(ranges[0], "%d-%d", &r1.start, &r1.end)
	fmt.Sscanf(ranges[1], "%d-%d", &r2.start, &r2.end)

	return r1, r2
}

func partOne(input string) int {
	score := 0

	for _, rs := range strings.Split(input, "\n") {
		r1, r2 := rangePairFromString(rs)
		if r1.Includes(r2) || r2.Includes(r1) {
			score++
		}
	}

	return score
}

func partTwo(input string) int {
	score := 0

	for _, rs := range strings.Split(input, "\n") {
		r1, r2 := rangePairFromString(rs)
		if r1.Overlaps(r2) {
			score++
		}
	}

	return score
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
