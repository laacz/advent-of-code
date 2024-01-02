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
			}
		}
	}

	return hp, d, a
}

type Boss struct {
	hp, d, a int
	name     string
	hard     bool
}

type Effects struct {
	shield, poison, recharge int
}

type State struct {
	mana, hp, boss_hp int
	spells            map[string]int
	mana_spent        int
}

type Spell struct {
	cost                      int
	damage, heal, armor, mana int
	duration                  int
}

var spells = map[string]Spell{
	"Magic Missile": {53, 4, 0, 0, 0, 1},
	"Drain":         {73, 2, 2, 0, 0, 1},
	"Shield":        {113, 0, 0, 7, 0, 6},
	"Poison":        {173, 3, 0, 0, 0, 6},
	"Recharge":      {229, 0, 0, 0, 101, 5},
}

func applySpells(state *State, is_boss bool) {
	for name := range state.spells {
		if state.spells[name] > 0 {
			state.boss_hp -= spells[name].damage
			state.hp += spells[name].heal
			state.mana += spells[name].mana
			if is_boss {
				// I do not have armor, but while the spell is active, I do
				state.hp += spells[name].armor
			}
			state.spells[name]--
		}
	}
}

func clash(boss Boss) (ret int) {
	states := []State{{
		500, 50, boss.hp,
		map[string]int{},
		0,
	}}

	ret = 999999
	for len(states) > 0 {
		state := states[0]
		states = states[1:]

		// hard mode means losing 1 hp per turn
		if boss.hard {
			state.hp -= 1
		}

		if state.hp <= 0 {
			continue
		}

		// player's turn
		applySpells(&state, false)

		if state.boss_hp <= 0 {
			if state.mana_spent < ret {
				ret = state.mana_spent
			}
			continue
		}

		for name, spell := range spells {
			if spell.cost <= state.mana && // we can affor the spell
				state.mana_spent+spell.cost < ret && // it's not more expensive than already cheapest solution
				state.spells[name] == 0 { // it's not active

				new_state := State{
					state.mana, state.hp, state.boss_hp,
					map[string]int{},
					state.mana_spent,
				}

				for k, v := range state.spells {
					new_state.spells[k] = v
				}

				new_state.mana -= spell.cost
				new_state.mana_spent += spell.cost
				new_state.spells[name] = spell.duration

				// boss's turn
				applySpells(&new_state, true)

				if new_state.boss_hp <= 0 {
					if new_state.mana_spent < ret {
						ret = new_state.mana_spent
					}
					continue
				}

				new_state.hp -= boss.d
				if new_state.hp > 0 {
					states = append(states, new_state)
				}
			}
		}
	}
	return ret
}

func partOne(lines string) (ret int) {
	hp, d, a := parse(lines)
	boss := Boss{hp, d, a, "boss", false}
	ret = clash(boss)

	return ret
}

func partTwo(lines string) (ret int) {
	hp, d, a := parse(lines)
	boss := Boss{hp, d, a, "boss", true}
	ret = clash(boss)

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
