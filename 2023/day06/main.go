package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// parse (you won't believe it!) returns parsed input
func parse(input string) [][]int {
	ret := [][]int{}

	times := []string{}
	distances := []string{}
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "Time:") {
			times = strings.Fields(line[10:])
		} else if strings.HasPrefix(line, "Distance:") {
			distances = strings.Fields(line[10:])
		}
	}

	for i := 0; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])
		ret = append(ret, []int{time, distance})
	}

	return ret
}

// partOne returns the answer to part one of this day's puzzle.
func partOne(input string) int {
	sum := 1

	for _, r := range parse(input) {
		t := r[0]
		d := r[1]

		wins := 0

		for i := 0; i <= t; i++ {
			d1 := (t - i) * i
			if d1 > d {
				wins += 1
			}
		}

		sum *= wins

	}

	return sum
}

// partTwo returns the answer to part one of this day's puzzle.
func partTwo(input string) int {
	sum := 1

	lines := strings.Split(input, "\n")

	for _, r := range parse(
		lines[0][0:11] + strings.Replace(lines[0][11:], " ", "", -1) +
			"\n" +
			lines[1][0:11] + strings.Replace(lines[1][11:], " ", "", -1),
	) {
		t := r[0]
		d := r[1]

		wins := 0

		for i := 0; i <= t; i++ {
			d1 := (t - i) * i
			if d1 > d {
				wins += 1
			}
		}

		sum *= wins

	}
	return sum
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
