package main

import (
	"fmt"
	"strings"

	"github.com/laacz/aoc-2016/util"
)

func parse(input string) (string, int) {
	return util.GetLines(input)[0], util.Atoi(util.GetLines(input)[1])
}

func nextRow(prevrow string) (string, int) {
	ret := ""

	for r := 0; r < len(prevrow); r++ {
		m := ""
		if r < 1 {
			m += "."
		} else {
			m += string(prevrow[r-1])
		}

		m += string(prevrow[r])

		if r+1 >= len(prevrow) {
			m += "."
		} else {
			m += string(prevrow[r+1])
		}

		trapma := map[string]bool{
			"^^.": true,
			".^^": true,
			"^..": true,
			"..^": true,
		}

		if trapma[m] {
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
