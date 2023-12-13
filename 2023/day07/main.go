package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards string
	bid   int
}

// This is a type of hand
const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

// Type returns the type of the hand
func (h Hand) Type() int {
	cards := make(map[rune]int)
	counts := make(map[int]string)

	for _, c := range h.cards {
		cards[c]++
	}

	for a, c := range cards {
		counts[c] += string(rune(a))
	}

	if counts[5] != "" {
		return FiveOfAKind
	}

	if counts[4] != "" {
		return FourOfAKind
	}

	if counts[3] != "" && counts[2] != "" {
		return FullHouse
	}

	if counts[3] != "" {
		return ThreeOfAKind
	}

	if counts[2] != "" && len(counts[2]) == 2 {
		return TwoPair
	}

	if counts[2] != "" {
		return OnePair
	}

	if len(counts[1]) == 5 {
		return HighCard
	}

	return 0
}

// StrongestType returns the strongest possible hand given that there is a joker
func (h Hand) StrongestType() int {
	ret := 0
	for c := range strength {
		hand := Hand{strings.Replace(h.cards, "0", string(c), -1), 0}
		if hand.Type() > ret {
			ret = hand.Type()
		}
	}

	return ret
}

// strength is a map of card strengths
var strength = map[rune]int{
	'0': 0,
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'J': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

// parse (you won't believe it!) returns parsed input
func parse(input string) (parsed []Hand) {
	ret := []Hand{}

	for _, line := range strings.Split(input, "\n") {
		hand := Hand{}
		tmp := strings.Split(line, " ")
		hand.cards = tmp[0]
		hand.bid, _ = strconv.Atoi(tmp[1])
		ret = append(ret, hand)
	}

	return ret
}

// partOne returns the answer to part one of this day's puzzle.
func partOne(input string) int {
	sum := 0

	hands := parse(input)
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Type() == hands[j].Type() {
			for k := 0; k < 5; k++ {
				if strength[rune(hands[i].cards[k])] == strength[rune(hands[j].cards[k])] {
					continue
				}

				return strength[rune(hands[i].cards[k])] < strength[rune(hands[j].cards[k])]
			}
		}

		return hands[i].Type() < hands[j].Type()
	})

	rank := 1
	for _, h := range hands {
		sum += h.bid * rank
		rank += 1
	}

	return sum
}

// partTwo returns the answer to part one of this day's puzzle.
func partTwo(input string) int {
	sum := 0

	hands := parse(strings.Replace(input, "J", "0", -1))
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].StrongestType() == hands[j].StrongestType() {
			for k := 0; k < 5; k++ {
				if strength[rune(hands[i].cards[k])] == strength[rune(hands[j].cards[k])] {
					continue
				}

				return strength[rune(hands[i].cards[k])] < strength[rune(hands[j].cards[k])]
			}
		}

		return hands[i].StrongestType() < hands[j].StrongestType()
	})

	rank := 1
	for _, h := range hands {
		sum += h.bid * rank
		rank += 1
	}

	return sum
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
