package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/laacz/aoc-2016/util"
)

type Instruction interface {
	Apply(string) string
	Unapply(string) string
}

func dbg(i Instruction, s string) {
	switch i.(type) {
	case SwapPosition:
		fmt.Printf("SwapPosition: %v", i)
	case SwapLetter:
		fmt.Printf("SwapLetter: %v", i)
	case RotateLeft:
		fmt.Printf("RotateLeft: %v", i)
	case RotateRight:
		fmt.Printf("RotateRight: %v", i)
	case RotateBased:
		fmt.Printf("RotateBased: %+v", i)
	case Reverse:
		fmt.Printf("Reverse: %v", i)
	case Move:
		fmt.Printf("Move: %v", i)
	default:
		panic("Unknown instruction")
	}

	fmt.Printf(", str: %s\n", s)
}

type Instructions []Instruction
type SwapPosition struct {
	X, Y int
}

func (i SwapPosition) Apply(s string) string {
	b := []byte(s)
	b[i.X], b[i.Y] = b[i.Y], b[i.X]

	return string(b)
}

func (i SwapPosition) Unapply(s string) string {
	return i.Apply(s)
}

type SwapLetter struct {
	X, Y byte
}

func (i SwapLetter) Apply(s string) string {
	b := []byte(s)
	for j, c := range b {
		if c == i.X {
			b[j] = i.Y
		} else if c == i.Y {
			b[j] = i.X
		}
	}

	return string(b)
}

func (i SwapLetter) Unapply(s string) string {
	return i.Apply(s)
}

type RotateLeft struct {
	Steps int
}

func (i RotateLeft) Apply(s string) string {
	b := []byte{}

	b = append(b, s[i.Steps%len(s):]...)
	b = append(b, s[:i.Steps%len(s)]...)

	return string(b)
}

func (i RotateLeft) Unapply(s string) string {
	return RotateRight{i.Steps}.Apply(s)
}

type RotateRight struct {
	Steps int
}

func (i RotateRight) Apply(s string) string {
	b := []byte{}

	b = append(b, s[(len(s)*2-i.Steps)%len(s):]...)
	b = append(b, s[:(len(s)*2-i.Steps)%len(s)]...)

	return string(b)
}

func (i RotateRight) Unapply(s string) string {
	return RotateLeft{i.Steps}.Apply(s)
}

type RotateBased struct {
	Letter string
}

type Reverse struct {
	X, Y int
}

func (i Reverse) Apply(s string) string {
	b := []byte(s)

	for i, j := i.X, i.Y; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

func (i Reverse) Unapply(s string) string {
	return i.Apply(s)
}

func (i RotateBased) Apply(s string) string {
	idx := strings.Index(s, i.Letter)

	m := map[int]int{
		0: 1,
		1: 2,
		2: 3,
		3: 4,
		4: 6,
		5: 7,
		6: 8,
		7: 9,
	}

	return RotateRight{m[idx]}.Apply(s)
}

func (i RotateBased) Unapply(s string) string {
	idx := strings.Index(s, i.Letter)

	m := map[int]int{
		1: 1,
		3: 2,
		5: 3,
		7: 4,
		2: 6,
		4: 7,
		6: 8,
		0: 9,
	}

	return RotateLeft{m[idx]}.Apply(s)
}

type Move struct {
	X, Y int
}

func (i Move) Apply(s string) string {
	b := []byte(s)

	c := b[i.X]
	b = append(b[:i.X], b[i.X+1:]...)
	b = append(b[:i.Y], append([]byte{c}, b[i.Y:]...)...)

	return string(b)
}

func (i Move) Unapply(s string) string {
	return Move{i.Y, i.X}.Apply(s)
}

func parse(input string) (ret Instructions, str, unstr string) {
	lines := util.GetLines(input)
	str = lines[0]
	unstr = lines[1]
	for _, line := range lines[2:] {
		var i Instruction
		switch {
		case strings.HasPrefix(line, "swap position"):
			var x, y int
			fmt.Sscanf(line, "swap position %d with position %d", &x, &y)
			i = SwapPosition{x, y}
		case strings.HasPrefix(line, "swap letter"):
			var x, y byte
			fmt.Sscanf(line, "swap letter %c with letter %c", &x, &y)
			i = SwapLetter{x, y}
		case strings.HasPrefix(line, "rotate left"):
			var x int
			fmt.Sscanf(line, "rotate left %d steps", &x)
			i = RotateLeft{x}
		case strings.HasPrefix(line, "rotate right"):
			var x int
			fmt.Sscanf(line, "rotate right %d steps", &x)
			i = RotateRight{x}
		case strings.HasPrefix(line, "rotate based"):
			var x byte
			fmt.Sscanf(line, "rotate based on position of letter %c", &x)
			i = RotateBased{string(rune(x))}
		case strings.HasPrefix(line, "reverse positions"):
			var x, y int
			fmt.Sscanf(line, "reverse positions %d through %d", &x, &y)
			i = Reverse{x, y}
		case strings.HasPrefix(line, "move position"):
			var x, y int
			fmt.Sscanf(line, "move position %d to position %d", &x, &y)
			i = Move{x, y}
		default:
			panic("Unknown instruction: " + line)
		}

		ret = append(ret, i)
	}

	return
}

func partOne(input string) (ret string) {
	instructions, ret, _ := parse(input)

	for _, i := range instructions {
		// dbg(i, ret)
		ret = i.Apply(ret)
		// fmt.Println("    result: ", ret)
	}

	return
}

func partTwo(input string) (ret string) {
	instructions, _, ret := parse(input)

	slices.Reverse(instructions)

	for _, i := range instructions {
		// dbg(i, ret)
		ret = i.Unapply(ret)
		// fmt.Println("    result: ", ret)
	}

	return
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}
