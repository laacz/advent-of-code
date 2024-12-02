package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func validateList(l []int) bool {
	increasing := l[1] > l[0]
	for i := 0; i < len(l)-1; i++ {
		delta := abs(l[i+1] - l[i])

		if increasing != (l[i+1] > l[i]) {
			return false
		}

		if delta < 1 || delta > 3 {
			return false
		}
	}

	return true
}

func part1(list [][]int) int {
	var ret int

	for _, l := range list {
		if validateList(l) {
			ret++
		}
	}

	return ret
}

func part2(list [][]int) int {
	var ret int

	for _, l := range list {
		for i := range l {
			newList := append(slices.Clone(l[:i]), l[i+1:]...)
			if validateList(newList) {
				ret++
				break
			}
		}
	}

	return ret
}

func main() {
	var lists [][]int

	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var list []int
		for _, num := range strings.Fields(line) {
			n, _ := strconv.Atoi(num)
			list = append(list, n)
		}
		lists = append(lists, list)
	}

	fmt.Println("part1", part1(lists))
	fmt.Println("part2", part2(lists))
}
