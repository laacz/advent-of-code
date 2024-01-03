package main

import (
	"fmt"

	"github.com/laacz/aoc-2016/util"
)

type Address struct {
	nets, hypernets []string
}

func (a *Address) addAddress(s string, hyper bool) {
	if s == "" {
		return
	}

	if hyper {
		a.hypernets = append(a.hypernets, s)
	} else {
		a.nets = append(a.nets, s)
	}
}

func parse(input string) (ret []Address) {
	for _, line := range util.GetLines(input) {
		var a Address
		var s string
		var hyper bool

		for _, c := range line {
			if c == '[' || c == ']' {
				a.addAddress(s, hyper)
				hyper = !hyper
				s = ""
			} else {
				s += string(c)
			}
		}
		a.addAddress(s, hyper)
		ret = append(ret, a)
	}

	return ret
}

func hasABBA(a string) bool {
	for i := 0; i < len(a)-3; i++ {
		if a[i] == a[i+3] && a[i+1] == a[i+2] && a[i] != a[i+1] {
			return true
		}
	}

	return false
}

func hasABA(a string) (ret []string) {
	for i := 0; i < len(a)-2; i++ {
		if a[i] == a[i+2] && a[i] != a[i+1] {
			ret = append(ret, a[i:i+3])
		}
	}

	return ret
}

func hasBAB(a string, aba string) bool {
	for i := 0; i < len(a)-2; i++ {
		if a[i] == aba[1] && a[i+1] == aba[0] && a[i+2] == aba[1] {
			return true
		}
	}

	return false
}

func partOne(input string) (ret int) {
	ads := parse(input)
	for _, a := range ads {
		TLS := false

		for _, h := range a.nets {
			if hasABBA(h) {
				TLS = true
				for _, n := range a.hypernets {
					if hasABBA(n) {
						TLS = false
						break
					}
				}
			}
		}

		if TLS {
			ret++
		}
	}

	return ret
}

func partTwo(input string) (ret int) {
	ads := parse(input)
	for _, a := range ads {

		SSL := false
		for _, n := range a.nets {
			for _, aba := range hasABA(n) {
				for _, h := range a.hypernets {
					if hasBAB(h, aba) {
						SSL = true
						break
					}
				}
			}
		}

		if SSL {
			ret++
		}
	}

	return ret
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}
