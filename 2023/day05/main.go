package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	dstRange, srcRange, rangeLen int
}

type Map struct {
	from, to string
	ranges   []Range
}

type Almanac struct {
	maps map[string]Map
}

func (a Almanac) Locate(seed int) int {
	state := "seed"
	number := seed

	for {
		m := a.maps[state]

		for _, r := range m.ranges {
			if number >= r.srcRange && number < r.srcRange+r.rangeLen {
				number = r.dstRange + (number - r.srcRange)
				state = m.to
				break
			}
		}

		if state != m.to {
			state = m.to
		}

		if state == "location" {
			break
		}
	}

	return number
}

// parse (you won't believe it!) returns parsed input
func parse(input string) ([]int, Almanac) {
	seeds := []int{}
	almanac := Almanac{
		maps: make(map[string]Map),
	}

	lines := strings.Split(input, "\n")

	var m Map

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "seeds: ") {
			nums := strings.Fields(line[6:])
			for _, numStr := range nums {
				num, err := strconv.Atoi(numStr)
				if err == nil {
					seeds = append(seeds, num)
				}
			}
		} else if strings.HasSuffix(line, " map:") {
			arr := strings.Split(line[:len(line)-5], "-")
			m = Map{
				from: arr[0],
				to:   arr[2],
			}
			almanac.maps[m.from] = m
		} else if line != "" {
			var dstRange, srcRange, rangeLen int
			fmt.Sscanf(line, "%d %d %d", &dstRange, &srcRange, &rangeLen)

			m := almanac.maps[almanac.maps[m.from].from]
			m.ranges = append(m.ranges, Range{
				dstRange: dstRange,
				srcRange: srcRange,
				rangeLen: rangeLen,
			})
			almanac.maps[m.from] = m
		}
	}

	return seeds, almanac
}

// partOne returns the answer to part one of this day's puzzle.
func partOne(input string) int {
	sum := -1

	s, a := parse(input)

	for _, seed := range s {
		l := a.Locate(seed)
		if sum == -1 || l < sum {
			sum = l
		}
	}

	return sum
}

// partTwo returns the answer to part one of this day's puzzle.
func partTwo(input string) int {
	sum := -1

	s, a := parse(input)

	for i := 0; i < len(s); i += 2 {
		start, length := s[i], s[i+1]

		for seed := start; seed < start+length; seed++ {

			l := a.Locate(seed)

			if sum == -1 || l < sum {
				sum = l
			}
		}
	}

	return sum
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
