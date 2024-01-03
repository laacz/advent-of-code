package main

import (
	"fmt"

	"github.com/laacz/aoc-2016/util"
)

func FreqMap(lines []string) []map[byte]int {
	freqmap := make([]map[byte]int, len(lines[0]))

	for _, line := range lines {
		for i, c := range []byte(line) {
			if freqmap[i] == nil {
				freqmap[i] = make(map[byte]int)
			}
			freqmap[i][c]++
		}
	}

	return freqmap
}

func partOne(input string) (ret string) {
	lines := util.GetLines(input)

	freqmap := FreqMap(lines)
	for i := 0; i < len(lines[0]); i++ {
		max := 0
		var maxc byte
		for c, f := range freqmap[i] {
			if f > max {
				max = f
				maxc = c
			}
		}
		ret += string(maxc)
	}

	return ret
}

func partTwo(input string) (ret string) {
	lines := util.GetLines(input)

	freqmap := FreqMap(lines)
	for i := 0; i < len(lines[0]); i++ {
		min := 999
		var minc byte
		for c, f := range freqmap[i] {
			if f < min {
				min = f
				minc = c
			}
		}
		ret += string(minc)
	}

	return ret
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}
