package main

import (
	"fmt"
	"os"
	"strings"
)

type Coord [2]int

type Robot struct {
	pos Coord
	v   Coord
}

func parseInput(input string) Robots {
	var ret Robots

	for _, lines := range strings.Split(strings.TrimSpace(input), "\n") {
		robot := Robot{}
		fmt.Sscanf(lines, "p=%d,%d v=%d,%d", &robot.pos[0], &robot.pos[1], &robot.v[0], &robot.v[1])
		ret = append(ret, robot)
	}

	return ret
}

type Robots []Robot

func (r Robots) String(size Coord) string {
	builder := strings.Builder{}
	builder.Grow(size[0]*size[1] + size[1])

	positions := make(map[Coord]int, len(r))
	for _, robot := range r {
		positions[robot.pos]++
	}

	row := make([]byte, size[0])

	for y := 0; y < size[1]; y++ {
		for x := 0; x < size[0]; x++ {
			row[x] = '.'
		}

		for x := 0; x < size[0]; x++ {
			if count := positions[Coord{x, y}]; count > 0 {
				row[x] = byte('0' + count)
			}
		}

		builder.Write(row)
		builder.WriteByte('\n')
	}

	return builder.String()
}

func part1(input Robots, size Coord) int {
	robots := make(Robots, len(input))
	copy(robots, input)

	for k, r := range robots {
		r.pos[0] += 100 * r.v[0]
		r.pos[1] += 100 * r.v[1]

		r.pos[0] %= size[0]
		r.pos[1] %= size[1]

		if r.pos[0] < 0 {
			r.pos[0] = size[0] + r.pos[0]
		}

		if r.pos[1] < 0 {
			r.pos[1] = size[1] + r.pos[1]
		}

		robots[k] = r
	}

	mid := Coord{size[0] / 2, size[1] / 2}
	counts := make([]int, 4)
	for _, r := range robots {
		var quadrant int
		if r.pos[0] < mid[0] && r.pos[1] < mid[1] {
			quadrant = 0
		} else if r.pos[0] > mid[0] && r.pos[1] < mid[1] {
			quadrant = 1
		} else if r.pos[0] < mid[0] && r.pos[1] > mid[1] {
			quadrant = 2
		} else if r.pos[0] > mid[0] && r.pos[1] > mid[1] {
			quadrant = 3
		} else {
			continue
		}
		counts[quadrant]++
	}

	return counts[0] * counts[1] * counts[2] * counts[3]
}
func part2(input Robots, size Coord) int {
	var ret int

	robots := make(Robots, len(input))
	copy(robots, input)

	for {
		ret++
		if ret%1000 == 0 {
			fmt.Println(ret)
		}
		for k, r := range robots {
			r.pos[0] += r.v[0]
			r.pos[1] += r.v[1]

			r.pos[0] %= size[0]
			r.pos[1] %= size[1]

			if r.pos[0] < 0 {
				r.pos[0] = size[0] + r.pos[0]
			}

			if r.pos[1] < 0 {
				r.pos[1] = size[1] + r.pos[1]
			}

			robots[k] = r
		}

		if strings.Contains(robots.String(size), "1111111111") {
			fmt.Println(robots.String(size))
			break
		}
	}

	return ret
}

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Println("part1", part1(parseInput(string(input)), Coord{101, 103}))
	fmt.Println("part2", part2(parseInput(string(input)), Coord{101, 103}))
}
