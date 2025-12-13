package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func parse(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n");
}

func seatId(boardingPass string) int {
	row := 0
	col := 0

	for i := 0; i < 7; i++ {
		if boardingPass[i] == 'B' {
			row |= 1 << (6 - i)
		}
	}

	for i := 7; i < 10; i++ {
		if boardingPass[i] == 'R' {
			col |= 1 << (2 - (i - 7))
		}
	}

	return row*8 + col
}

func part1(input []string) int {
	ret := 0

	for _, seat := range input {
		id := seatId(seat)
		if id > ret {
			ret = id
		}
	}

	return ret
}

func part2(input []string) int {
	ids := []int{}
	for _, seat := range input {
		id := seatId(seat)
		ids = append(ids, id)
	}

	sort.Ints(ids)

	for i := 1; i < len(ids); i++ {
		if ids[i] != ids[i-1]+1 {
			return ids[i-1] + 1
		}
	}

	return -1
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Println("Part 1:", part1(parse(string(data))))
	fmt.Println("Part 2:", part2(parse(string(data))))
}
