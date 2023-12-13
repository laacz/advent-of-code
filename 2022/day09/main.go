package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}

type Coords struct {
	X int
	Y int
}

func partOne(input string) int {
	directions := map[string]Coords{"U": {1, 0}, "D": {-1, 0}, "L": {0, -1}, "R": {0, 1}}

	visited := map[Coords]bool{}

	head := Coords{0, 0}
	tail := Coords{0, 0}

	for _, line := range strings.Split(input, "\n") {
		steps, _ := strconv.Atoi(line[2:])
		dir := line[:1]
		delta := directions[dir]

		for i := 0; i < steps; i++ {
			head.Y += delta.Y
			head.X += delta.X

			if abs(head.X-tail.X) > 1 || abs(head.Y-tail.Y) > 1 {
				if (dir == "U" || dir == "D") && head.X != tail.X {
					tail.X = head.X
				} else if (dir == "L" || dir == "R") && head.Y != tail.Y {
					tail.Y = head.Y
				}

				tail.Y = head.Y - delta.Y
				tail.X = head.X - delta.X
			}

			visited[tail] = true
		}

	}

	return len(visited)
}

func partTwo(input string) int {
	directions := map[string]Coords{"R": {1, 0}, "L": {-1, 0}, "D": {0, -1}, "U": {0, 1}}

	visited := map[Coords]bool{}
	knots := make([]Coords, 10)

	for _, line := range strings.Split(input, "\n") {
		steps, _ := strconv.Atoi(line[2:])
		dir := line[:1]

		for i := 0; i < steps; i++ {
			dx := directions[dir].X
			dy := directions[dir].Y
			knots[0].Y += dy
			knots[0].X += dx

			for j := 0; j < len(knots)-1; j++ {
				head := knots[j]
				tail := knots[j+1]

				dx = head.X - tail.X
				dy = head.Y - tail.Y

				if abs(dx) > 1 || abs(dy) > 1 {
					tail.X += sign(dx)
					tail.Y += sign(dy)
				}

				knots[j] = head
				knots[j+1] = tail

			}
			visited[knots[9]] = true
		}

	}

	return len(visited)
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
