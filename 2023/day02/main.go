package main

import (
	"fmt"
	"os"
	"strings"
)

type Games map[int]Game
type Game []Round

// Possible returns true if the game is possible with the given number of red, green, and blue cubes.
func (g Game) Possible(red, green, blue int) bool {
	for _, round := range g {
		if !round.Possible(red, green, blue) {
			return false
		}
	}
	return true
}

// maxint returns the larger of two integers
func maxint(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Fewest returns the least number of cubes needed to play the game.
func (g Game) Fewest() (int, int, int) {
	var red, green, blue int

	for _, round := range g {
		for color, count := range round {
			switch color {
			case "red":
				red = maxint(red, count)
			case "green":
				green = maxint(green, count)
			case "blue":
				blue = maxint(blue, count)
			}
		}
	}

	return red, green, blue
}

type Round map[string]int

// Possible returns true if the round is possible with the given number of red, green, and blue cubes.
func (r Round) Possible(red, green, blue int) bool {
	return r["red"] <= red && r["green"] <= green && r["blue"] <= blue
}

// parse does just that
func parse(data string) Games {
	lines := strings.Split(data, "\n")
	games := make(Games)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		var game int
		var round Round

		l, _ := fmt.Sscanf(line, "Game %d:", &game)

		if _, ok := games[game]; !ok {
			games[game] = make([]Round, 0)
		}

		for _, part := range strings.Split(line[l+7:], ";") {
			round = make(Round)

			for _, item := range strings.Split(part, ",") {
				var count int
				var color string

				fmt.Sscanf(item, "%d %s", &count, &color)
				if color != "" {
					round[color] = count
				}
			}

			games[game] = append(games[game], round)
		}
	}

	return games
}

// partOne returns the answer to part one of this day's puzzle.
func partOne(games Games) int {
	sum := 0

	for gn, game := range games {
		if game.Possible(12, 13, 14) {
			sum += gn
		}
	}

	return sum
}

// partTwo returns the answer to part two of this day's puzzle.
func partTwo(games Games) int {
	sum := 0

	for _, game := range games {
		r, g, b := game.Fewest()

		sum += r * g * b
	}

	return sum
}

func main() {
	data, _ := os.ReadFile("input.txt")
	games := parse(string(data))

	fmt.Printf("Part one: %d\n", partOne(games))
	fmt.Printf("Part two: %d\n", partTwo(games))
}
