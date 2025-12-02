package main

import (
	"fmt"
	"os"
	"strconv"
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
	instructions := parseInstructions(input)

	// Starting position and direction
	x, y := 0, 0
	direction := 0 // 0=East, 90=South, 180=West, 270=North

	for _, instr := range instructions {
		action := instr.action
		value := instr.value

		switch action {
		case 'N':
			y += value
		case 'S':
			y -= value
		case 'E':
			x += value
		case 'W':
			x -= value
		case 'L':
			direction = (direction - value + 360) % 360
		case 'R':
			direction = (direction + value) % 360
		case 'F':
			switch direction {
			case 0: // East
				x += value
			case 90: // South
				y -= value
			case 180: // West
				x -= value
			case 270: // North
				y += value
			}
		}
	}

	return abs(x) + abs(y)
}

type instruction struct {
	action rune
	value  int
}

func parseInstructions(input string) []instruction {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	instructions := make([]instruction, len(lines))

	for i, line := range lines {
		line = strings.TrimSpace(line)
		action := rune(line[0])
		value, _ := strconv.Atoi(line[1:])
		instructions[i] = instruction{action: action, value: value}
	}

	return instructions
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func part2(input string) int {
	instructions := parseInstructions(input)

	// Starting position of ship
	shipX, shipY := 0, 0
	// Waypoint position relative to ship
	waypointX, waypointY := 10, 1

	for _, instr := range instructions {
		action := instr.action
		value := instr.value

		switch action {
		case 'N':
			waypointY += value
		case 'S':
			waypointY -= value
		case 'E':
			waypointX += value
		case 'W':
			waypointX -= value
		case 'L':
			// Rotate waypoint counter-clockwise
			for i := 0; i < value/90; i++ {
				waypointX, waypointY = -waypointY, waypointX
			}
		case 'R':
			// Rotate waypoint clockwise
			for i := 0; i < value/90; i++ {
				waypointX, waypointY = waypointY, -waypointX
			}
		case 'F':
			// Move ship towards waypoint
			shipX += waypointX * value
			shipY += waypointY * value
		}
	}

	return abs(shipX) + abs(shipY)
}
