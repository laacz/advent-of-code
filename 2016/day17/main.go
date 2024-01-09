package main

import (
	"fmt"

	"github.com/laacz/aoc-2016/util"
)

func parse(input string) (ret string) {
	return util.GetLines(input)[0]
}

type State struct {
	x, y int
	path string
}

func partOne(input string) (ret string) {
	passcode := parse(input)

	q := []State{{0, 0, ""}}
	for len(q) > 0 {
		s := q[0]
		q = q[1:]

		hash := util.Md5(passcode + s.path)
		for d, dir := range map[int][2]int{
			0: {0, -1},
			1: {0, 1},
			2: {-1, 0},
			3: {1, 0},
		} {
			if hash[d] >= 'b' && hash[d] <= 'f' {
				x, y := s.x+dir[0], s.y+dir[1]
				fp := s.path + string("UDLR"[d])

				if x == 3 && y == 3 {
					ret = fp
					return
				}

				if x >= 0 && x < 4 && y >= 0 && y < 4 {
					q = append(q, State{x, y, fp})
				}
			}
		}
	}

	return
}

func partTwo(input string) (ret int) {
	passcode := parse(input)

	q := []State{{0, 0, ""}}

	seen := map[string]bool{}

	for len(q) > 0 {
		s := q[0]
		q = q[1:]

		hash := util.Md5(passcode + s.path)
		for d, dir := range map[int][2]int{
			0: {0, -1},
			1: {0, 1},
			2: {-1, 0},
			3: {1, 0},
		} {
			if hash[d] >= 'b' && hash[d] <= 'f' {
				x, y := s.x+dir[0], s.y+dir[1]
				fp := s.path + string("UDLR"[d])

				if x == 3 && y == 3 {
					if !seen[fp] {
						ret = util.Max(ret, len(fp))
					}
					continue
				}

				if x >= 0 && x < 4 && y >= 0 && y < 4 {
					q = append(q, State{x, y, fp})
				}
			}
		}
	}

	return
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}
