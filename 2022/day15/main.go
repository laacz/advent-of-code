package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X, Y int
	D    int
}

var beacons []Point
var sensors []Point
var minx, maxx, maxd int

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func manhattanDistance(a, b Point) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func parseInput(input string) {
	beacons = []Point{}
	sensors = []Point{}
	for _, line := range strings.Split(input, "\n") {
		var bx, by, sx, sy int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		d := abs(sx-bx) + abs(sy-by)
		minx = min(minx, sx)
		maxx = max(maxx, sx)
		maxd = max(maxd, d)
		beacons = append(beacons, Point{X: bx, Y: by})
		sensors = append(sensors, Point{X: sx, Y: sy, D: d})
	}
}

func partOne(y int) int {
	ret := 0
	for x := minx - maxd; x < maxx+maxd; x++ {
		for _, b := range beacons {
			if b.X == x && b.Y == y {
				ret -= 1
				break
			}
		}
		for _, s := range sensors {
			if manhattanDistance(s, Point{X: x, Y: y}) <= s.D {
				ret += 1
				break
			}
		}
	}

	return ret
}

type Line struct {
	x1, y1, x2, y2 int
	s              float64
	yi             float64
}

func (l Line) slope() float64 {
	if l.s == 0 {
		l.s = float64(l.y2-l.y1) / float64(l.x2-l.x1)
	}
	return l.s
}

func (l Line) yIntercept() float64 {
	if l.yi == 0 {
		l.yi = float64(l.y1) - l.slope()*float64(l.x1)
	}
	return l.yi
}

func partTwo(mi, ma int) int {
	ret := 0

	// Generate lines of manhattan distance + 1
	lines := []Line{}
	for _, pt := range sensors {
		lines = append(lines, Line{x1: pt.X - pt.D - 1, y1: pt.Y, x2: pt.X, y2: pt.Y - pt.D - 1})
		lines = append(lines, Line{x1: pt.X + pt.D + 1, y1: pt.Y, x2: pt.X, y2: pt.Y - pt.D - 1})
		lines = append(lines, Line{x1: pt.X - pt.D - 1, y1: pt.Y, x2: pt.X, y2: pt.Y + pt.D + 1})
		lines = append(lines, Line{x1: pt.X + pt.D + 1, y1: pt.Y, x2: pt.X, y2: pt.Y + pt.D + 1})
	}

	// Now find all possible itnersections (ignore if lines do not cross)
	points := []Point{}
	for _, line1 := range lines {
		for _, line2 := range lines {
			if line1.slope() == line2.slope() {
				continue
			}
			x := (line2.yIntercept() - line1.yIntercept()) / (line1.slope() - line2.slope())
			y := line1.slope()*x + line1.yIntercept()
			points = append(points, Point{X: int(x), Y: int(y)})
		}
	}

	// Now find the first point which is not in the range of any sensor
outer:
	for _, pt := range points {
		if pt.X < mi || pt.X > ma || pt.Y < mi || pt.Y > ma {
			continue
		}
		for _, s := range sensors {
			if manhattanDistance(s, pt) <= s.D {
				continue outer
			}
		}

		return pt.X*4000000 + pt.Y
	}

	return ret
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	parseInput(string(data))

	fmt.Printf("Part one: %d\n", partOne(2000000))
	fmt.Printf("Part two: %d\n", partTwo(0, 4000000))
}
