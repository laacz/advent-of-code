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
				Type:             TypeConjuction,
				Outputs:          outputs,
				Name:             name[1:],
				ConjuctionMemory: make(map[string]Pulse),
			})
		}
	}

	for _, module := range ret {
		for _, output := range module.Outputs {
			m := ret.Find(output)
			if m != nil && m.Type == TypeConjuction {
				m.ConjuctionMemory[module.Name] = LOW
			}
		}
	}

	return ret
}

const (
	TypeBroadcaster = iota
	TypeFlipFlop
	TypeConjuction
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
	Type             int
	Outputs          []string
	Name             string
	FlipFlopState    bool
	ConjuctionMemory map[string]Pulse
}

func (m *Module) String() string {
	t := map[int]string{
		TypeBroadcaster: "br",
		TypeFlipFlop:    "fl",
		TypeConjuction:  "con",
	}
	return fmt.Sprintf("[%3s] %s", t[m.Type], m.Name)
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

var hz, pv, qh, xm int

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
		case TypeConjuction:
			p.Target.ConjuctionMemory[p.Source.Name] = p.Pulse

			if p.Target.Name == "kh" && p.Pulse == HIGH {
				fmt.Printf("Processing pulse: %+v\n", p)
				fmt.Println("  kh memory: ", p.Target.ConjuctionMemory)
				if hz == 0 && p.Target.ConjuctionMemory["hz"] == HIGH {
					hz = i
					fmt.Println("hz", i)
				}
				if pv == 0 && p.Target.ConjuctionMemory["pv"] == HIGH {
					pv = i
					fmt.Println("pv", i)
				}
				if qh == 0 && p.Target.ConjuctionMemory["qh"] == HIGH {
					qh = i
					fmt.Println("qh", i)
				}
				if xm == 0 && p.Target.ConjuctionMemory["xm"] == HIGH {
					xm = i
					fmt.Println("xm", i)
				}

			}

			out := LOW
			for _, v := range p.Target.ConjuctionMemory {
				if v == LOW {
					out = HIGH
					break
				}
			}

			// fmt.Println("    memory:", p.Target.ConjuctionMemory, "output:", out)

			for _, output := range p.Target.Outputs {
				qq = append(qq, QQ{output, out})
			}
		}

		for _, pulse := range qq {
			module := m.Find(pulse.Name)
			// fmt.Println(p.Target, "-", pulse.Pulse, "-", pulse.Name)
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
	fmt.Println(high, low)

	return high * low
}

// partTwo returns the answer to part one of this day's puzzle.
func partTwo(input string) (ret int) {
	modules := parse(input)

	i := 0
	for hz == 0 || pv == 0 || qh == 0 || xm == 0 {
		i += 1
		modules.PushButton(i)
		if i > 10000 {
			break
		}
	}

	return hz * pv * qh * xm
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
