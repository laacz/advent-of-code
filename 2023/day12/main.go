package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	pattern    string
	groupsizes []int
}

func (r *Record) Unfold() {
	tmp := make([]string, 5)
	groupsizes := []int{}
	for i := range tmp {
		tmp[i] = r.pattern
		groupsizes = append(groupsizes, r.groupsizes...)
	}

	r.pattern = strings.Join(tmp, "?")
	r.groupsizes = groupsizes
}

type Records []Record

func parse(input string) (ret Records) {
	for _, line := range strings.Split(input, "\n") {
		tmp := strings.Split(line, " ")
		tmp2 := strings.Split(tmp[1], ",")
		sizes := make([]int, 0)
		for i, s := range tmp2 {
			sizes = append(sizes, 0)
			sizes[i], _ = strconv.Atoi(s)
		}
		ret = append(ret, Record{
			pattern:    tmp[0],
			groupsizes: sizes,
		})
	}

	return ret
}

var cache = map[string]int{}

func combos(pattern string, groupsizes []int) (res int) {
	key := pattern + "," + fmt.Sprint(groupsizes)

	// memoization
	if val, ok := cache[key]; ok {
		return val
	}

	// if pattern is empty
	if pattern == "" {
		// if there are no groupsizes left
		if len(groupsizes) == 0 {
			return 1
		}
		return 0
	}

	// if there are no groupsizes left
	if len(groupsizes) == 0 {
		// if there are any springs left
		if strings.Contains(pattern, "#") {
			return 0
		}
		return 1
	}

	// if pattern starts with a possible whitespace
	if pattern[0] == '.' || pattern[0] == '?' {
		res += combos(pattern[1:], groupsizes)
	}

	if pattern[0] == '#' || pattern[0] == '?' {
		// if pattern starts with a possible spring
		if groupsizes[0] <= len(pattern) && !strings.Contains(pattern[:groupsizes[0]], ".") &&
			(groupsizes[0] == len(pattern) || pattern[groupsizes[0]] != '#') {
			// if group is at all possible with the remaining pattern
			if groupsizes[0] == len(pattern) {
				// first group's size is exactly the length of the pattern
				res += combos("", groupsizes[1:])
			} else {
				// first group's size is less than the length of the pattern
				res += combos(pattern[groupsizes[0]+1:], groupsizes[1:])
			}
		}
	}

	cache[key] = res

	return res
}

// partOne returns the answer to part one of this day's puzzle.
func partOne(input string) (ret int) {
	records := parse(input)

	for _, r := range records {
		ret += combos(r.pattern, r.groupsizes)
	}

	return ret
}

// partTwo returns the answer to part one of this day's puzzle.
func partTwo(input string) (ret int) {
	records := parse(input)

	for _, r := range records {
		r.Unfold()
		c := combos(r.pattern, r.groupsizes)
		ret += c
	}

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
