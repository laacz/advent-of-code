package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type Combo struct {
	generator, chip int
}

type State struct {
	combos Combos
	floor  int
}

func (s State) Assembled() bool {
	for _, c := range s.combos {
		if c.generator != 4 || c.chip != 4 {
			return false
		}
	}
	return true
}

func (s *State) Move(dir int, chips, generators []string) bool {
	for _, chip := range chips {
		c := s.combos[chip]
		c.chip += dir
		s.combos[chip] = c
	}

	for _, generator := range generators {
		c := s.combos[generator]
		c.generator += dir
		s.combos[generator] = c
	}

	s.floor += dir

	for i := 1; i <= 4; i++ {
		for _, c := range s.combos {
			if c.generator == i {
				for _, cc := range s.combos {
					if cc.chip == i && cc.generator != i {
						return false
					}
				}
			}
		}
	}

	return true
}

func (s State) copy() State {
	ret := State{
		combos: make(Combos, len(s.combos)),
		floor:  s.floor,
	}
	for name, c := range s.combos {
		ret.combos[name] = c
	}
	return ret
}

func (s State) nextStates() (ret []State) {
	// holds all chips and generators on the current floor
	var chips, generators []string

	for name, c := range s.combos {
		if c.generator == s.floor {
			generators = append(generators, name)
		}
		if c.chip == s.floor {
			chips = append(chips, name)
		}
	}

	// Build move options: singles and pairs
	type move struct {
		chips      []string
		generators []string
		count      int // number of items being moved
	}
	var moves []move

	// Single chips
	for _, c := range chips {
		moves = append(moves, move{[]string{c}, []string{}, 1})
	}
	// Single generators
	for _, g := range generators {
		moves = append(moves, move{[]string{}, []string{g}, 1})
	}
	// Chip pairs
	for i, c := range chips {
		for j := i + 1; j < len(chips); j++ {
			moves = append(moves, move{[]string{c, chips[j]}, []string{}, 2})
		}
	}
	// Generator pairs
	for i, g := range generators {
		for j := i + 1; j < len(generators); j++ {
			moves = append(moves, move{[]string{}, []string{g, generators[j]}, 2})
		}
	}
	// Matching chip+generator pairs
	for name, c := range s.combos {
		if c.generator == s.floor && c.chip == s.floor {
			moves = append(moves, move{[]string{name}, []string{name}, 2})
		}
	}

	// Check if all floors below are empty (optimization: don't go down if nothing below)
	allBelowEmpty := true
	for _, c := range s.combos {
		if c.generator < s.floor || c.chip < s.floor {
			allBelowEmpty = false
			break
		}
	}

	// Generate states for each direction
	for _, dir := range []int{-1, 1} {
		if s.floor+dir < 1 || s.floor+dir > 4 {
			continue
		}

		// Skip going down if all floors below are empty
		if dir == -1 && allBelowEmpty {
			continue
		}

		var validMoves []State
		hasTwo := false
		hasOne := false

		for _, m := range moves {
			newState := s.copy()
			if newState.Move(dir, m.chips, m.generators) {
				validMoves = append(validMoves, newState)
				if m.count == 2 {
					hasTwo = true
				} else {
					hasOne = true
				}
			}
		}

		// Optimization: when going UP, prefer 2-item moves; when going DOWN, prefer 1-item moves
		for _, m := range moves {
			// Skip if we're going up and have 2-item moves available but this is a 1-item move
			if dir == 1 && hasTwo && m.count == 1 {
				continue
			}
			// Skip if we're going down and have 1-item moves available but this is a 2-item move
			if dir == -1 && hasOne && m.count == 2 {
				continue
			}

			newState := s.copy()
			if newState.Move(dir, m.chips, m.generators) {
				ret = append(ret, newState)
			}
		}
	}

	return ret
}

type Combos map[string]Combo

// Hash returns a canonical string representation for state comparison.
// Element names don't matter - only the configuration of (generator_floor, chip_floor) pairs.
func (s State) Hash() string {
	var pairs []string
	for _, c := range s.combos {
		pairs = append(pairs, fmt.Sprintf("%d%d", c.generator, c.chip))
	}
	sort.Strings(pairs)
	return fmt.Sprintf("%d:%s", s.floor, strings.Join(pairs, ","))
}

func PrintTransition(s1, s2 State) (ret string) {
	l1 := strings.Split(s1.String(), "\n")
	l2 := strings.Split(s2.String(), "\n")

	for i := 0; i < 4; i++ {
		for j := 0; j < len(l1[i]); j++ {
			if l1[i][j] != l2[i][j] && l1[i][j] != '.' {
				ret += "\033[38;2;0;254;254m"
			}
			ret += string(l1[i][j])
			ret += "\033[0m"
		}
		ret += " -> "
		for j := 0; j < len(l2[i]); j++ {
			if l1[i][j] != l2[i][j] && l2[i][j] != '.' {
				ret += "\033[38;5;208m"
			}
			ret += string(l2[i][j])
			ret += "\033[0m"
		}
		ret += "\n"
	}

	return
}

// colWidth calculates column width based on longest element name
func colWidth(combos Combos) int {
	maxLen := 1
	for name := range combos {
		if len(name) > maxLen {
			maxLen = len(name)
		}
	}
	return maxLen + 1 // +1 for G/M suffix
}

// Make State hashable
func (c State) String() (ret string) {
	var names []string
	for name := range c.combos {
		names = append(names, name)
	}
	sort.Slice(names, func(i, j int) bool { return names[i] < names[j] })

	w := colWidth(c.combos)
	dotPad := strings.Repeat(" ", w-1) + "."

	for f := 4; f >= 1; f-- {
		ret += fmt.Sprintf("F%d ", f)
		if f == c.floor {
			ret += "E  "
		} else {
			ret += ".  "
		}
		for _, name := range names {
			if c.combos[name].generator == f {
				ret += fmt.Sprintf("%*s", w, name+"G")
			} else {
				ret += dotPad
			}

			ret += " "

			if c.combos[name].chip == f {
				ret += fmt.Sprintf("%*s", w, name+"M")
			} else {
				ret += dotPad
			}
			ret += " "
		}
		ret += "\n"
	}
	return ret
}

// ANSI color codes
const (
	colorReset  = "\033[0m"
	colorCyan   = "\033[38;2;0;254;254m"
	colorOrange = "\033[38;5;208m"
	colorGreen  = "\033[38;2;0;255;0m"
	colorYellow = "\033[38;5;226m"
)

// Animation line offset (set before running dfs)
var animationStartLine = 1

// moveCursor moves cursor to line (1-based)
func moveCursor(line int) {
	fmt.Printf("\033[%d;1H", line)
}

// clearFromCursor clears from cursor to end of screen
func clearFromCursor() {
	fmt.Print("\033[J")
}

// stateWidth calculates display width of a state
func stateWidth(state State) int {
	w := colWidth(state.combos)
	// "F4 " + "E  " + (width + " " + width + " ") * numCombos
	return 3 + 3 + len(state.combos)*(w+1+w+1)
}

// displayState shows current state at animation position with step info
func displayState(state State, step, queueSize, visited int) {
	moveCursor(animationStartLine)
	clearFromCursor()

	width := stateWidth(state)
	if width < 50 {
		width = 50
	}

	fmt.Printf("%sStep: %d%s | %sQueue: %d%s | %sVisited: %d%s\n",
		colorGreen, step, colorReset,
		colorYellow, queueSize, colorReset,
		colorCyan, visited, colorReset)
	fmt.Println(strings.Repeat("─", width))
	fmt.Print(colorizeState(state))
	fmt.Println(strings.Repeat("─", width))
}

// colorizeState adds colors to state output
func colorizeState(state State) string {
	var names []string
	for name := range state.combos {
		names = append(names, name)
	}
	sort.Slice(names, func(i, j int) bool { return names[i] < names[j] })

	w := colWidth(state.combos)
	dotPad := strings.Repeat(" ", w-1) + "."

	var ret string
	for f := 4; f >= 1; f-- {
		ret += fmt.Sprintf("F%d ", f)
		if f == state.floor {
			ret += colorOrange + "E" + colorReset + "  "
		} else {
			ret += ".  "
		}
		for _, name := range names {
			if state.combos[name].generator == f {
				ret += colorCyan + fmt.Sprintf("%*s", w, name+"G") + colorReset
			} else {
				ret += dotPad
			}

			ret += " "

			if state.combos[name].chip == f {
				ret += colorGreen + fmt.Sprintf("%*s", w, name+"M") + colorReset
			} else {
				ret += dotPad
			}
			ret += " "
		}
		ret += "\n"
	}
	return ret
}

func dfs(state State) (ret int) {
	queue := []State{state}
	visited := make(map[string]bool)
	visited[state.Hash()] = true

	for {
		ret++
		newqueue := []State{}
		for _, state := range queue {
			// Display current state being explored (only first state per step to reduce overhead)
			if len(newqueue) == 0 {
				displayState(state, ret, len(queue), len(visited))
				time.Sleep(30 * time.Millisecond) // Small delay for animation effect
			}

			for _, s := range state.nextStates() {
				if s.Assembled() {
					displayState(s, ret, len(queue), len(visited))
					return
				}

				if visited[s.Hash()] {
					continue
				}

				visited[s.Hash()] = true
				newqueue = append(newqueue, s)
			}
		}
		queue = newqueue
		if len(queue) == 0 {
			panic("no solution")
		}
	}
}

func partOne(state State) (ret int) {
	return dfs(state)
}

func partTwo(state State) (ret int) {
	return dfs(state)
}

func main() {
	// Clear screen at start
	fmt.Print("\033[2J\033[H")

	// The first floor contains a promethium generator and a promethium-compatible microchip.
	// The second floor contains a cobalt generator, a curium generator, a ruthenium generator, and a plutonium generator.
	// The third floor contains a cobalt-compatible microchip, a curium-compatible microchip, a ruthenium-compatible microchip, and a plutonium-compatible microchip.
	// The fourth floor contains nothing relevant.

	// Part 1: animate starting at line 1
	animationStartLine = 1
	result1 := partOne(State{
		combos: Combos{
			"Pm": {1, 1},
			"Co": {2, 3},
			"Cm": {2, 3},
			"Ru": {2, 3},
			"Pu": {2, 3},
		},
		floor: 1,
	})

	// Move below part 1 display (1 header + 1 line + 4 floor lines + 1 line = 7 lines)
	// Print result
	fmt.Printf("\n%sPart one: %d%s\n\n", colorGreen, result1, colorReset)

	// Part 2: animate starting below part 1 (line 10 or so)
	animationStartLine = 10
	result2 := partTwo(State{
		combos: Combos{
			"Pm": {1, 1},
			"Co": {2, 3},
			"Cm": {2, 3},
			"Ru": {2, 3},
			"Pu": {2, 3},
			"El": {1, 1}, // elerium generator and chip on floor 1
			"Di": {1, 1}, // dilithium generator and chip on floor 1
		},
		floor: 1,
	})
	fmt.Printf("\n%sPart two: %d%s\n", colorGreen, result2, colorReset)
}
