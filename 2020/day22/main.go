package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(input string) ([]int, []int) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")

	var p1, p2 []int
	for _, line := range strings.Split(parts[0], "\n")[1:] {
		n, _ := strconv.Atoi(line)
		p1 = append(p1, n)
	}
	for _, line := range strings.Split(parts[1], "\n")[1:] {
		n, _ := strconv.Atoi(line)
		p2 = append(p2, n)
	}

	return p1, p2
}

func score(deck []int) int {
	total := 0
	for i, card := range deck {
		total += card * (len(deck) - i)
	}
	return total
}

func part1(p1, p2 []int) int {
	d1 := make([]int, len(p1))
	d2 := make([]int, len(p2))
	copy(d1, p1)
	copy(d2, p2)

	for len(d1) > 0 && len(d2) > 0 {
		c1, c2 := d1[0], d2[0]
		d1, d2 = d1[1:], d2[1:]

		if c1 > c2 {
			d1 = append(d1, c1, c2)
		} else {
			d2 = append(d2, c2, c1)
		}
	}

	if len(d1) > 0 {
		return score(d1)
	}
	return score(d2)
}

func deckKey(d1, d2 []int) string {
	var sb strings.Builder
	for _, c := range d1 {
		sb.WriteString(strconv.Itoa(c))
		sb.WriteByte(',')
	}
	sb.WriteByte('|')
	for _, c := range d2 {
		sb.WriteString(strconv.Itoa(c))
		sb.WriteByte(',')
	}
	return sb.String()
}

func recursiveCombat(d1, d2 []int) (int, []int) {
	seen := make(map[string]bool)

	for len(d1) > 0 && len(d2) > 0 {
		key := deckKey(d1, d2)
		if seen[key] {
			return 1, d1
		}
		seen[key] = true

		c1, c2 := d1[0], d2[0]
		d1, d2 = d1[1:], d2[1:]

		var winner int
		if len(d1) >= c1 && len(d2) >= c2 {
			sub1 := make([]int, c1)
			sub2 := make([]int, c2)
			copy(sub1, d1[:c1])
			copy(sub2, d2[:c2])
			winner, _ = recursiveCombat(sub1, sub2)
		} else if c1 > c2 {
			winner = 1
		} else {
			winner = 2
		}

		if winner == 1 {
			d1 = append(d1, c1, c2)
		} else {
			d2 = append(d2, c2, c1)
		}
	}

	if len(d1) > 0 {
		return 1, d1
	}
	return 2, d2
}

func part2(p1, p2 []int) int {
	d1 := make([]int, len(p1))
	d2 := make([]int, len(p2))
	copy(d1, p1)
	copy(d2, p2)

	_, winningDeck := recursiveCombat(d1, d2)
	return score(winningDeck)
}

func main() {
	data, _ := os.ReadFile("input.txt")
	p1, p2 := parse(string(data))

	fmt.Println("Part 1:", part1(p1, p2))
	fmt.Println("Part 2:", part2(p1, p2))
}
