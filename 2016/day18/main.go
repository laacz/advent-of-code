package main

import (
	"fmt"
	"strings"

	"github.com/laacz/aoc-2016/util"
)

func parse(input string) (string, int) {
	return util.GetLines(input)[0], util.Atoi(util.GetLines(input)[1])
}

func nextRow(pr string) (string, int) {
	prevrow := []byte("." + pr + ".")
	ret := ""

	for r := 0; r < len(prevrow)-2; r++ {
		w := uint32(prevrow[r])<<16 + uint32(prevrow[r+1])<<8 + uint32(prevrow[r+2])

		// ..^ 3026526
		// .^^ 3038814
		// ^.. 6172206
		// ^^. 6184494
		if w == 3026526 || w == 3038814 || w == 6172206 || w == 6184494 {
			ret += "^"
		} else {
			ret += "."
		}
	}

	return ret, strings.Count(ret, ".")
}

func partOne(input string) (ret int) {
	prevrow, cols := parse(input)

	for c := 1; c <= cols; c++ {
		next, cnt := nextRow(prevrow)
		ret += cnt
		prevrow = next
	}

	return
}

func partTwo(input string) (ret int) {
	prevrow, _ := parse(input)
	cols := 400000

	for c := 1; c <= cols; c++ {
		next, cnt := nextRow(prevrow)
		ret += cnt
		prevrow = next
	}

	return
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}
