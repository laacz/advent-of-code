package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := part1(string(data))
	fmt.Println("Part 1:", result)

	result2 := part2(string(data))
	fmt.Println("Part 2:", result2)
}

func parse(input string) (int, []int) {
	lines := strings.Split(input, "\n")
	time, _ := strconv.Atoi(lines[0])
	ids := strings.Split(lines[1], ",")
	var intervals []int
	for _, id := range ids {
		id, _ := strconv.Atoi(id)
		intervals = append(intervals, id)
	}

	return time, intervals
}

func part1(input string) int {
	time, intervals := parse(input)

	min := 99999999
	id := 0
	for _, interval := range intervals {
		if interval == 0 {
			continue
		}
		diff := interval - time%interval
		if diff < min {
			min = diff
			id = interval
		}
	}

	return min * id
}


func part2(input string) int {
	_, ids := parse(input)

	t := 0
	step := 1

	for offset, id := range ids {
		if id == 0 {
			continue
		}
		for (t+offset)%id != 0 {
			t += step
		}
		step *= id
	}

	return t
}
