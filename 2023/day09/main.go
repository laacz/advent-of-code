package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type ValueHistory struct {
	values []int
	tails  []int
	heads  []int
}

// sumInts returns the sum of all ints in the slice
func notZeroes(ints []int) bool {
	for _, v := range ints {
		if v != 0 {
			return true
		}
	}
	return false
}

// ExtrapolateTail extrapolates next value of the history
func (h *ValueHistory) ExtrapolateTail() int {
	ret := 0

	values := h.values

	for notZeroes(values) {
		var v []int
		h.tails = append(h.tails, values[len(values)-1])
		for i := 1; i < len(values); i++ {
			v = append(v, values[i]-values[i-1])
		}
		values = v
	}

	slices.Reverse(h.tails)

	for _, v := range h.tails {
		ret = ret + v
	}

	return ret
}

// ExtrapolateHead extrapolates previous value of the history
func (h *ValueHistory) ExtrapolateHead() int {
	ret := 0

	values := h.values

	for notZeroes(values) {
		var v []int
		h.heads = append(h.heads, values[0])
		for i := 1; i < len(values); i++ {
			v = append(v, values[i]-values[i-1])
		}
		values = v
	}

	slices.Reverse(h.heads)

	for _, v := range h.heads {
		ret = v - ret
	}

	return ret
}

// parse (you won't believe it!) returns parsed input
func parse(input string) []ValueHistory {
	var ret []ValueHistory

	for _, line := range strings.Split(input, "\n") {
		h := ValueHistory{}
		for _, w := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(w)
			h.values = append(h.values, num)
		}

		ret = append(ret, h)
	}

	return ret
}

// partOne returns the answer to part one of this day's puzzle.
func partOne(input string) int {
	sum := 0

	histories := parse(input)

	for _, h := range histories {
		r := h.ExtrapolateTail()
		sum += r
	}

	return sum
}

// partTwo returns the answer to part one of this day's puzzle.
func partTwo(input string) int {
	sum := 0

	histories := parse(input)

	for _, h := range histories {
		r := h.ExtrapolateHead()
		sum += r
	}

	return sum
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
