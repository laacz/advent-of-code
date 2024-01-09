package main

import (
	"fmt"
	"sort"

	"github.com/laacz/aoc-2016/util"
)

type Range struct {
	From, To int
}

type Ranges []Range

func parse(input string) (ret Ranges) {
	for _, line := range util.GetLines(input) {
		var r Range
		fmt.Sscanf(line, "%d-%d", &r.From, &r.To)
		ret = append(ret, r)
	}

	return
}

func partOne(input string) (ret int) {
	ranges := parse(input)
	ret = int(0xFFFFFFFF)

	// sort ranges
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].From < ranges[j].From
	})

	// now unmerge overlapping ranges
	for i := 0; i < len(ranges)-1; i++ {
		if ranges[i].To >= ranges[i+1].From-1 {
			ranges[i].To = ranges[i+1].To
			ranges = append(ranges[:i+1], ranges[i+2:]...)
			i--
		}
	}

	ret = ranges[0].To + 1

	return
}

func partTwo(input string) (ret int) {
	ranges := parse(input)

	// sort ranges
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].From < ranges[j].From
	})

	// fuck it, there are overlaping ranges and I love them
	max := 0
	for _, r := range ranges {
		if r.From > max {
			ret += r.From - max
		}
		max = util.Max(max, r.To+1)
	}

	ret += 0xffffffff - max + 1

	return
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}
