package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	winning []int
	guessed []int
	copies  int
}

// MatchingCount returns the number of matching numbers between the winning and guessed numbers.
func (c *Card) MatchingCount() int {
	n := 0
	for _, w := range c.winning {
		for _, g := range c.guessed {
			if w == g {
				n++
			}
		}
	}

	return n
}

type Cards struct {
	cards []Card
	pos   int
}

// Play plays the part two game with copies.
func (c *Cards) Play() int {
	ret := 0
	for c.pos < len(c.cards) {
		card := c.cards[c.pos]
		cw := card.MatchingCount()
		for j := 0; j < card.copies; j++ {
			for i := 1; i <= cw; i++ {
				c.cards[c.pos+i].copies += 1
			}
			ret += 1
		}
		c.pos += 1
	}

	return ret
}

// parse returns a Cards struct from the input.
func parse(input string) Cards {
	cards := Cards{}
	for _, line := range strings.Split(input, "\n") {
		c := Card{
			copies: 1,
		}
		c.winning = []int{}
		c.guessed = []int{}

		nums := strings.Split(strings.Split(line, ":")[1], "|")

		for _, w := range strings.Split(nums[0], " ") {
			if n, err := strconv.Atoi(w); err == nil {
				c.winning = append(c.winning, n)
			}
		}
		for _, w := range strings.Split(nums[1], " ") {
			if n, err := strconv.Atoi(w); err == nil {
				c.guessed = append(c.guessed, n)
			}
		}

		cards.cards = append(cards.cards, c)
	}

	return cards
}

// partOne returns the answer to part one of this day's puzzle.
func partOne(input string) int {
	sum := 0

	cards := parse(input)
	for _, c := range cards.cards {
		cnt := c.MatchingCount()

		if cnt-2 > 0 {
			cnt = 2 << (cnt - 2)
		}

		sum += cnt
	}

	return sum
}

// partTwo returns the answer to part one of this day's puzzle.
func partTwo(input string) int {
	cards := parse(input)

	return cards.Play()
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
