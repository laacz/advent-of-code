package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput(input string) []int {
	ret := make([]int, 0)
	file := true
	var index int
	var app int

	for _, c := range strings.TrimSpace(input) {
		len := c - '0'
		if !file {
			app = -1
		} else {
			app = index
			index++
		}
		for range len {
			ret = append(ret, app)
		}
		file = !file
	}
	return ret
}

func DefragmentBlocks(input []int) []int {
	ret := make([]int, len(input))
	copy(ret, input)

	tail := len(ret) - 1
	for head := range len(ret) {
		for ret[head] != -1 {
			head++
		}
		for ret[tail] == -1 {
			tail--
		}
		if head >= tail {
			break
		}

		ret[head] = ret[tail]
		ret[tail] = -1
		tail--
	}

	return ret
}

func DefragmentFiles(input []int) []int {
	ret := make([]int, len(input))
	copy(ret, input)

	tail := len(ret) - 1
	moved := make(map[int]bool)

	for tail >= 0 {
		for tail >= 0 && ret[tail] == -1 {
			tail--
		}
		if tail < 0 {
			break
		}

		index := ret[tail]
		tailEnd := tail
		for tail > 0 && ret[tail-1] == index {
			tail--
		}

		if tail == 0 {
			break
		}

		length := tailEnd - tail + 1
		moved[index] = true

		head := 0
		for head < tail {
			for head < tail && ret[head] != -1 {
				head++
			}
			if head >= tail {
				break
			}

			start := head
			for head < len(ret) && ret[head] == -1 {
				head++
			}

			if gap := head - start; gap >= length {
				for i := range length {
					ret[start+i] = index
				}
				for i := range tailEnd - tail + 1 {
					ret[tail+i] = -1
				}
				break
			}
		}
		tail--
	}

	return ret
}

func Checksum(input []int) int {
	var ret int
	for i, c := range input {
		if c == -1 {
			continue
		}
		ret += c * i
	}
	return ret
}

func part1(input string) int {
	disk := DefragmentBlocks(parseInput(input))
	return Checksum(disk)
}

func part2(input string) int {
	disk := DefragmentFiles(parseInput(input))
	return Checksum(disk)
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println("part1", part1(string(input)))
	fmt.Println("part2", part2(string(input)))
}
