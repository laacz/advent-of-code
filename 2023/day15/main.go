package main

import (
	"fmt"
	"os"
	"strings"
)

func parse(input string) []string {
	input = strings.ReplaceAll(input, "\n", "")
	return strings.Split(input, ",")
}

func hash(s string) (ret int) {
	for _, c := range s {
		ret += int(c)
		ret *= 17
		ret %= 256
	}

	return ret
}

// partOne returns the answer to part one of this day's puzzle.
func partOne(input string) (ret int) {
	parse(input)
	for _, s := range parse(input) {
		ret += hash(s)
	}
	return ret
}

type Lens struct {
	label string
	fl    int
}

type Box struct {
	lenses []Lens
}

func (b *Box) add(l Lens) {
	for idx, lens := range b.lenses {
		if lens.label == l.label {
			b.lenses[idx].fl = l.fl
			return
		}
	}

	b.lenses = append(b.lenses, l)
}

func (b *Box) remove(l Lens) {
	for idx, lens := range b.lenses {
		if lens.label == l.label {
			b.lenses = append(b.lenses[:idx], b.lenses[idx+1:]...)
			return
		}
	}
}

func (b *Box) power(boxn int) (ret int) {
	for i, l := range b.lenses {
		ret += (i + 1) * (boxn + 1) * l.fl
	}
	return ret
}

// partTwo returns the answer to part one of this day's puzzle.
func partTwo(input string) (ret int) {
	parse(input)
	boxes := make([]Box, 265)
	steps := parse(input)
	for _, s := range steps {
		var label string
		var lens int
		if strings.HasSuffix(s, "-") {
			label = s[:len(s)-1]
		} else if strings.Contains(s, "=") {
			label = s[:strings.Index(s, "=")]
			fmt.Sscanf(s[strings.Index(s, "=")+1:], "%d", &lens)
		} else {
			label = s
		}
		boxidx := hash(label)
		l := Lens{label, lens}
		if lens > 0 {
			boxes[boxidx].add(l)
		} else {
			boxes[boxidx].remove(l)
		}
	}

	for i, b := range boxes {
		ret += b.power(i)
	}

	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
