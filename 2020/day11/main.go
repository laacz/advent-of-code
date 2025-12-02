package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := part1(string(data))
	fmt.Println("Part 1:", result)

	result2 := part2(string(data))
	fmt.Println("Part 2:", result2)
}

func part1(input string) int {
	grid := parseGrid(input)

	for {
		newGrid, changed := simulateRound(grid)
		if !changed {
			break
		}
		grid = newGrid
	}

	return countOccupied(grid)
}

func parseGrid(input string) [][]rune {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func simulateRound(grid [][]rune) ([][]rune, bool) {
	rows := len(grid)
	cols := len(grid[0])
	newGrid := make([][]rune, rows)
	changed := false

	for i := range rows {
		newGrid[i] = make([]rune, cols)
		copy(newGrid[i], grid[i])
	}

	for i := range rows {
		for j := range cols {
			if grid[i][j] == '.' {
				continue
			}

			occupied := countAdjacentOccupied(grid, i, j)

			if grid[i][j] == 'L' && occupied == 0 {
				newGrid[i][j] = '#'
				changed = true
			} else if grid[i][j] == '#' && occupied >= 4 {
				newGrid[i][j] = 'L'
				changed = true
			}
		}
	}

	return newGrid, changed
}

func countAdjacentOccupied(grid [][]rune, row, col int) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]

		if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
			if grid[newRow][newCol] == '#' {
				count++
			}
		}
	}

	return count
}

func countOccupied(grid [][]rune) int {
	count := 0
	for i := range len(grid) {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '#' {
				count++
			}
		}
	}
	return count
}

func part2(input string) int {
	grid := parseGrid(input)

	for {
		newGrid, changed := simulateRound2(grid)
		if !changed {
			break
		}
		grid = newGrid
	}

	return countOccupied(grid)
}

func simulateRound2(grid [][]rune) ([][]rune, bool) {
	rows := len(grid)
	cols := len(grid[0])
	newGrid := make([][]rune, rows)
	changed := false

	for i := range rows {
		newGrid[i] = make([]rune, cols)
		copy(newGrid[i], grid[i])
	}

	for i := range rows {
		for j := range cols {
			if grid[i][j] == '.' {
				continue
			}

			occupied := countVisibleOccupied(grid, i, j)

			if grid[i][j] == 'L' && occupied == 0 {
				newGrid[i][j] = '#'
				changed = true
			} else if grid[i][j] == '#' && occupied >= 5 {
				newGrid[i][j] = 'L'
				changed = true
			}
		}
	}

	return newGrid, changed
}

func countVisibleOccupied(grid [][]rune, row, col int) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]

		for newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
			if grid[newRow][newCol] == '#' {
				count++
				break
			} else if grid[newRow][newCol] == 'L' {
				break
			}
			newRow += dir[0]
			newCol += dir[1]
		}
	}

	return count
}
