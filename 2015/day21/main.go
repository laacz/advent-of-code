package main

import (
	"fmt"
	"os"
	"strings"
)

func parse(input string) (hp, d, a int) {
	lines := strings.Split(input, "\n\n")
	for _, line := range lines {
		parts := strings.Split(line, "\n")
		for _, part := range parts {
			if strings.HasPrefix(part, "Hit Points") {
				fmt.Sscanf(part, "Hit Points: %d", &hp)
			} else if strings.HasPrefix(part, "Damage") {
				fmt.Sscanf(part, "Damage: %d", &d)
			} else if strings.HasPrefix(part, "Armor") {
				fmt.Sscanf(part, "Armor: %d", &a)
			}
		}
	}

	return hp, d, a
}

type Player struct {
	hp, d, a int
	name     string
}

func round(player, boss *Player) bool {
	for boss.hp > 0 && player.hp > 0 {
		d := player.d - boss.a
		if d < 1 {
			d = 1
		}
		boss.hp -= d
		if boss.hp <= 0 {
			return true
		}

		d = boss.d - player.a
		if d < 1 {
			d = 1
		}
		player.hp -= d
		if player.hp <= 0 {
			return false
		}
	}
	return false
	// toWin := boss.hp / (player.d - boss.a)
	// return player.hp-(boss.a-player.a)*(toWin-1) >= 0
}

type Item struct {
	name string
	cost int
	d, a int
}

var weapons = []Item{
	{"Dagger", 8, 4, 0},
	{"Shortsword", 10, 5, 0},
	{"Warhammer", 25, 6, 0},
	{"Longsword", 40, 7, 0},
	{"Greataxe", 74, 8, 0},
}

var armor = []Item{
	{"No armor", 0, 0, 0},
	{"Leather", 13, 0, 1},
	{"Chainmail", 31, 0, 2},
	{"Splintmail", 53, 0, 3},
	{"Bandedmail", 75, 0, 4},
	{"Platemail", 102, 0, 5},
}

var rings = []Item{
	{"No ring 1", 0, 0, 0},
	{"no ring 2", 0, 0, 0},
	{"Damage +1", 25, 1, 0},
	{"Damage +2", 50, 2, 0},
	{"Damage +3", 100, 3, 0},
	{"Defense +1", 20, 0, 1},
	{"Defense +2", 40, 0, 2},
	{"Defense +3", 80, 0, 3},
}

func partOne(lines string) (ret int) {
	hp, d, a := parse(lines)
	boss := Player{hp, d, a, "boss"}

	ret = 9999
	for _, w := range weapons {
		for _, a := range armor {
			for _, r1 := range rings {
				for _, r2 := range rings {
					if r1.name == r2.name {
						continue
					}

					b := boss
					player := Player{100, w.d + a.d + r1.d + r2.d, w.a + a.a + r1.a + r2.a, "player"}

					if round(&player, &b) {
						if ret > w.cost+a.cost+r1.cost+r2.cost {
							ret = w.cost + a.cost + r1.cost + r2.cost
						}
					}
				}
			}
		}
	}

	return ret
}

func partTwo(lines string) (ret int) {
	hp, d, a := parse(lines)
	boss := Player{hp, d, a, "boss"}

	ret = 0
	for _, w := range weapons {
		for _, a := range armor {
			for _, r1 := range rings {
				for _, r2 := range rings {
					if r1.name == r2.name {
						continue
					}

					b := boss
					player := Player{100, w.d + a.d + r1.d + r2.d, w.a + a.a + r1.a + r2.a, "player"}

					if !round(&player, &b) {
						if ret < w.cost+a.cost+r1.cost+r2.cost {
							ret = w.cost + a.cost + r1.cost + r2.cost
						}
					}
				}
			}
		}
	}

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
