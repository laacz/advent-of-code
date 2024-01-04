package main

import (
	"fmt"

	"github.com/laacz/aoc-2016/util"
)

type Bot struct {
	id        string
	chips     []int
	low, high *Bot
}

type Bots map[string]*Bot

func (b *Bots) EnsureExists(id string) {
	if _, ok := (*b)[id]; !ok {
		(*b)[id] = &Bot{id: id}
	}
}
func (b Bots) String() string {
	ret := ""
	for _, bot := range b {
		ret += fmt.Sprintf("%s: %v\n", bot.id, bot.chips)
	}
	return ret
}

func (b Bots) Givers() (ret []*Bot) {
	for _, bot := range b {
		if len(bot.chips) == 2 {
			ret = append(ret, bot)
		}
	}
	return ret
}

func parse(input string) *Bots {
	bots := make(Bots)

	for _, line := range util.GetLines(input) {
		if line[0] == 'v' {
			var bot, chip int
			fmt.Sscanf(line, "value %d goes to bot %d", &chip, &bot)
			id := fmt.Sprintf("bot %d", bot)
			bots.EnsureExists(id)
			bots[id].chips = append(bots[id].chips, chip)
		} else {
			var bot, low, high int
			var lowType, highType string
			fmt.Sscanf(line, "bot %d gives low to %s %d and high to %s %d", &bot, &lowType, &low, &highType, &high)
			id := fmt.Sprintf("bot %d", bot)
			lowid := fmt.Sprintf("%s %d", lowType, low)
			highid := fmt.Sprintf("%s %d", highType, high)

			bots.EnsureExists(id)
			bots.EnsureExists(lowid)
			bots.EnsureExists(highid)

			bots[id].low = bots[lowid]
			bots[id].high = bots[highid]
		}
	}

	return &bots
}

func partOne(input string, l, h int) (ret int) {
	bots := parse(input)

	for len(bots.Givers()) > 0 {
		for _, bot := range bots.Givers() {
			var high, low = util.Max(bot.chips[0], bot.chips[1]), util.Min(bot.chips[0], bot.chips[1])

			if high == h && low == l {
				return util.Atoi(bot.id[4:])
			}

			if bot.chips[0] == 17 && bot.chips[1] == 61 {
				fmt.Println(bot.id)
			}
			if bot.low != nil {
				bot.low.chips = append(bot.low.chips, util.Min(bot.chips[0], bot.chips[1]))
			}
			if bot.high != nil {
				bot.high.chips = append(bot.high.chips, util.Max(bot.chips[0], bot.chips[1]))
			}
			bot.chips = nil
		}
	}

	return ret
}

func partTwo(input string) (ret int) {
	bots := parse(input)

	for len(bots.Givers()) > 0 {
		for _, bot := range bots.Givers() {
			if bot.low != nil {
				bot.low.chips = append(bot.low.chips, util.Min(bot.chips[0], bot.chips[1]))
			}
			if bot.high != nil {
				bot.high.chips = append(bot.high.chips, util.Max(bot.chips[0], bot.chips[1]))
			}
			bot.chips = nil
		}
	}

	ret = (*bots)["output 0"].chips[0] * (*bots)["output 1"].chips[0] * (*bots)["output 2"].chips[0]

	return ret
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println("Part one:", partOne(input, 17, 61))
	fmt.Println("Part two:", partTwo(input))
}
