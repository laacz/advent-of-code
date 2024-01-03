package main

import (
	"fmt"
	"strings"

	"github.com/laacz/aoc-2016/util"
)

type Step struct {
	lr byte
	n  int
}

func (s Step) String() string {
	return fmt.Sprintf("%c%d", s.lr, s.n)
}

func parse(input string) (ret []Step) {
	for _, line := range strings.Split(input, ", ") {
		ret = append(ret, Step{
			lr: line[0],
			n:  util.Atoi(line[1:]),
		})
	}
	return ret
}

func partOne(steps []Step) (ret int) {
	x, y := 0, 0
	dirs := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	dir := 0

	for _, step := range steps {
		switch step.lr {
		case 'R':
			dir = (dir + 1) % 4
		case 'L':
			dir = (dir + 3) % 4
		}

		x += dirs[dir][0] * step.n
		y += dirs[dir][1] * step.n
	}

	ret = util.Abs(x) + util.Abs(y)

	return ret
}

func partTwo(steps []Step) (ret int) {
	x, y := 0, 0
	dirs := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	dir := 0

	var visited = make(map[[2]int]bool)
outer:
	for _, step := range steps {
		switch step.lr {
		case 'R':
			dir = (dir + 1) % 4
		case 'L':
			dir = (dir + 3) % 4
		}

		for i := 0; i < step.n; i++ {
			x += dirs[dir][0]
			y += dirs[dir][1]

			if visited[[2]int{x, y}] {
				break outer
			}

			visited[[2]int{x, y}] = true
		}
	}

	ret = util.Abs(x) + util.Abs(y)

	return ret
}

func main() {
	input := parse(util.ReadFile("input.txt"))
	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}
