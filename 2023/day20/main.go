package main

import (
	"fmt"
	"os"
	"strings"
)

// parse parses the input into a Grid according to the rules of the first part of the puzzle
func parse(input string) (ret Modules) {
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		name := strings.Split(line, " -> ")[0]
		outputs := strings.Split(line, " -> ")[1]

		switch name[0] {
		case 'b':
			ret = append(ret, &Module{
				Type:    TypeBroadcaster,
				Outputs: strings.Split(outputs, ", "),
				Name:    "broadcaster",
			})
		case '%':
			ret = append(ret, &Module{
				Type:    TypeFlipFlop,
				Outputs: strings.Split(outputs, ", "),
				Name:    name[1:],
			})
		case '&':
			outputs := strings.Split(outputs, ", ")

			ret = append(ret, &Module{
				Type:              TypeConjunction,
				Outputs:           outputs,
				Name:              name[1:],
				ConjunctionMemory: make(map[string]Pulse),
			})
		}
	}

	for _, module := range ret {
		for _, output := range module.Outputs {
			m := ret.Find(output)
			if m != nil && m.Type == TypeConjunction {
				m.ConjunctionMemory[module.Name] = LOW
			}
		}
	}

	return ret
}

const (
	TypeBroadcaster = iota
	TypeFlipFlop
	TypeConjunction
)

const LOW Pulse = true
const HIGH Pulse = false

type Pulse bool

func (p Pulse) String() string {
	if p {
		return "LOW"
	}
	return "HIGH"
}

type Entry struct {
	Target *Module
	Source *Module
	Pulse  Pulse
}

type Module struct {
	Type              int
	Outputs           []string
	Name              string
	FlipFlopState     bool
	ConjunctionMemory map[string]Pulse
}

type Modules []*Module

func (m *Modules) Find(name string) *Module {
	for _, module := range *m {
		if module.Name == name {
			return module
		}
	}
	return nil
}

type QQ struct {
	Name  string
	Pulse Pulse
}

var watch = map[string]int{
	"hz": 0,
	"pv": 0,
	"qh": 0,
	"xm": 0,
}

func (m *Modules) PushButton(i int) (high, low int) {
	broadcaster := m.Find("broadcaster")
	queue := []Entry{
		{broadcaster, nil, LOW},
	}

	low = 1

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		qq := []QQ{}

		if p.Target == nil {
			continue
		}

		switch p.Target.Type {
		case TypeBroadcaster:
			for _, output := range p.Target.Outputs {
				qq = append(qq, QQ{output, p.Pulse})
			}
		case TypeFlipFlop:
			if p.Pulse == LOW {
				out := HIGH
				if p.Target.FlipFlopState {
					out = LOW
				}
				p.Target.FlipFlopState = !p.Target.FlipFlopState
				for _, output := range p.Target.Outputs {
					qq = append(qq, QQ{output, out})
				}
			}
		case TypeConjunction:
			p.Target.ConjunctionMemory[p.Source.Name] = p.Pulse

			if p.Target.Name == "kh" && p.Pulse == HIGH {
				for k, v := range p.Target.ConjunctionMemory {
					if v == HIGH && watch[k] == 0 {
						watch[k] = i
					}
				}
			}

			out := LOW
			for _, v := range p.Target.ConjunctionMemory {
				if v == LOW {
					out = HIGH
					break
				}
			}

			for _, output := range p.Target.Outputs {
				qq = append(qq, QQ{output, out})
			}
		}

		for _, pulse := range qq {
			module := m.Find(pulse.Name)

			if pulse.Pulse == HIGH {
				high++
			} else {
				low++
			}

			if module != nil {
				queue = append(queue, Entry{module, p.Target, pulse.Pulse})
			}
		}
	}

	return high, low
}

// partOne returns the answer to part one of this day's puzzle.
func partOne(input string) (ret int) {
	modules := parse(input)
	high, low := 0, 0

	for i := 0; i < 1000; i++ {
		h, l := modules.PushButton(i)
		high += h
		low += l
	}

	return high * low
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

// partTwo returns the answer to part one of this day's puzzle.
func partTwo(input string) (ret int) {
	modules := parse(input)

	i := 0
	for watch["hz"] == 0 || watch["pv"] == 0 || watch["qh"] == 0 || watch["xm"] == 0 {
		i += 1
		modules.PushButton(i)
	}

	return lcm(watch["hz"], watch["pv"], watch["qh"], watch["xm"])
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
