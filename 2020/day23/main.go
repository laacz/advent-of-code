package main

import (
	"fmt"
	"strconv"
)

func parse(input string) []int {
	var cups []int
	for _, c := range input {
		n, _ := strconv.Atoi(string(c))
		cups = append(cups, n)
	}
	return cups
}

func play(cups []int, moves int) []int {
	n := len(cups)
	next := make([]int, n+1)

	for i := 0; i < n-1; i++ {
		next[cups[i]] = cups[i+1]
	}
	next[cups[n-1]] = cups[0]

	current := cups[0]

	for m := 0; m < moves; m++ {
		p1 := next[current]
		p2 := next[p1]
		p3 := next[p2]

		next[current] = next[p3]

		dest := current - 1
		if dest < 1 {
			dest = n
		}
		for dest == p1 || dest == p2 || dest == p3 {
			dest--
			if dest < 1 {
				dest = n
			}
		}

		next[p3] = next[dest]
		next[dest] = p1

		current = next[current]
	}

	return next
}

func part1(cups []int) string {
	next := play(cups, 100)

	result := ""
	cur := next[1]
	for cur != 1 {
		result += strconv.Itoa(cur)
		cur = next[cur]
	}
	return result
}

func part2(cups []int) int {
	extended := make([]int, 1000000)
	copy(extended, cups)
	for i := len(cups); i < 1000000; i++ {
		extended[i] = i + 1
	}

	next := play(extended, 10000000)

	c1 := next[1]
	c2 := next[c1]
	return c1 * c2
}

func main() {
	cups := parse("872495136")

	fmt.Println("Part 1:", part1(cups))
	fmt.Println("Part 2:", part2(cups))
}
