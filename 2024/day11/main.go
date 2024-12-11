package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(input []int, blinks int) int {
	stones := make([]int, len(input))
	copy(stones, input)

	for range blinks {
		newstones := []int{}
		for _, stone := range stones {
			a := strconv.Itoa(stone)
			if stone == 0 {
				newstones = append(newstones, 1)
			} else if len(a)%2 == 0 {
				n1, _ := strconv.Atoi(a[:len(strconv.Itoa(stone))/2])
				n2, _ := strconv.Atoi(strings.TrimLeft(a[len(strconv.Itoa(stone))/2:], "0"))
				newstones = append(newstones, n1, n2)
			} else {
				newstones = append(newstones, stone*2024)
			}
		}
		stones = newstones
	}

	return len(stones)
}

func part2(input []int, blinks int) int {
	var ret int

	stones := map[int]int{}
	for _, stone := range input {
		stones[stone]++
	}

	for range blinks {
		newstones := map[int]int{}
		for stone, cnt := range stones {
			a := strconv.Itoa(stone)
			if stone == 0 {
				newstones[1] += cnt
			} else if len(a)%2 == 0 {
				n1, _ := strconv.Atoi(a[:len(strconv.Itoa(stone))/2])
				n2, _ := strconv.Atoi(strings.TrimLeft(a[len(strconv.Itoa(stone))/2:], "0"))
				newstones[n1] += cnt
				newstones[n2] += cnt
			} else {
				newstones[stone*2024] += cnt
			}
		}
		stones = newstones
	}

	for _, cnt := range stones {
		ret += cnt
	}
	return ret
}

func parseInput(input string) []int {
	var ret []int
	for _, s := range strings.Fields(strings.TrimSpace(input)) {
		num, _ := strconv.Atoi(s)
		ret = append(ret, num)
	}

	return ret
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("part1", part1(parseInput(string(input)), 25))
	fmt.Println("part2", part2(parseInput(string(input)), 75))
}
