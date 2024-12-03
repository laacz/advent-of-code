package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Mul struct {
	x int
	y int
}

func (m Mul) Apply() int {
	return m.x * m.y
}

type State struct {
	Args        []int
	ArgStr      strings.Builder
	ReadingArgs bool
}

func parseInput(input string, startStopMuls bool) []Mul {
	var ops []Mul
	var doMuls = true

	state := State{}
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '(':
			state = State{}
			if startStopMuls && i > 1 && input[i-2:i+1] == "do(" {
				doMuls = true
			} else if startStopMuls && i > 4 && input[i-5:i+1] == "don't(" {
				doMuls = false
			} else if i > 2 && input[i-3:i+1] == "mul(" && doMuls {
				state.ReadingArgs = true
			}
		case ')':
			if state.ReadingArgs && state.ArgStr.Len() > 0 {
				arg, _ := strconv.Atoi(state.ArgStr.String())
				state.Args = append(state.Args, arg)
				if len(state.Args) == 2 {
					ops = append(ops, Mul{state.Args[0], state.Args[1]})
				}
			}
			state = State{}
		case ',':
			if state.ReadingArgs && state.ArgStr.Len() > 0 {
				arg, _ := strconv.Atoi(state.ArgStr.String())
				state.Args = append(state.Args, arg)
				state.ArgStr.Reset()
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if state.ReadingArgs {
				state.ArgStr.WriteByte(input[i])
			}
		}
	}

	return ops
}
func part1(input string) int {
	var ret int

	ops := parseInput(input, false)
	for _, op := range ops {
		ret += op.Apply()
	}

	return ret
}

func part2(input string) int {
	var ret int

	ops := parseInput(input, true)
	for _, op := range ops {
		ret += op.Apply()
	}

	return ret
}

func main() {

	input, _ := os.ReadFile("input.txt")

	fmt.Println("part1", part1(string(input)))
	fmt.Println("part2", part2(string(input)))
}
