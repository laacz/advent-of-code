package main

import (
	"fmt"
	"os"
	"strings"
)

type Brick struct {
	x1, y1, z1 int
	x2, y2, z2 int
	name       string
}

func (b *Brick) Contains(x, y int) bool {
	return x >= b.x1 && x <= b.x2 && y >= b.y1 && y <= b.y2
}

func (b *Brick) Drop(bricks Bricks) bool {
	newbrick := *b

	fallen := false
	for {
		if newbrick.OnGround() || len(newbrick.BricksBelow(bricks)) > 0 {
			return fallen
		}

		newbrick.z1--
		newbrick.z2--

		b.z1--
		b.z2--

		fallen = true
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// draws projection of bricks on xz axis
func (b Bricks) String() (ret string) {
	maxz := 0
	maxx := 0
	maxy := 0
	for _, brick := range b {
		maxz = max(maxz, brick.z2)
		maxx = max(maxx, brick.x2)
		maxy = max(maxy, brick.y2)
	}

	ret += fmt.Sprint("┌", strings.Repeat("─", maxx+4), "┬", strings.Repeat("─", maxx+4), "┐\n")
	for z := maxz; z > 0; z-- {
		hasbricks := false
		line := "│ "
		for x := 0; x <= maxx; x++ {
			name := ""
			cnt := 0
			for _, brick := range b {
				if brick.x1 <= x && brick.x2 >= x && brick.z1 <= z && brick.z2 >= z {
					name = brick.name
					cnt += 1
				}
			}

			hasbricks = hasbricks || cnt > 0

			switch cnt {
			case 0:
				line += "\033[90m·\033[0m"
			case 1:
				line += name
			default:
				line += "\033[93m?\033[0m"
			}
		}

		line += fmt.Sprintf("\033[90m %3d \033[0m", z)

		for y := 0; y <= maxy; y++ {
			cnt := 0
			name := ""
			for _, brick := range b {
				if brick.y1 <= y && brick.y2 >= y && brick.z1 <= z && brick.z2 >= z {
					name = brick.name
					cnt += 1
				}
			}

			hasbricks = hasbricks || cnt > 0

			switch cnt {
			case 0:
				line += "\033[90m·\033[0m"
			case 1:
				line += name
			default:
				line += "\033[93m?\033[0m"
			}
		}
		line += " │"
		// if hasbricks {
		ret += fmt.Sprintln(line)
		// }
	}
	ret += fmt.Sprint("└", strings.Repeat("─", maxx+4), "┴", strings.Repeat("─", maxx+4), "┘\n")

	return ret
}

func (b *Brick) OnGround() bool {
	return b.z1 <= 1
}

// CountTouching counts how many other bricks are touching this brick on given axis
func (b *Brick) Supports(bricks Bricks) (ret []string) {

	m := map[string]bool{}
	for _, a := range bricks {
		if a == b {
			continue
		}

		if b.z1 == a.z2+1 {
			overlap := !(a.x2 < b.x1 || b.x2 < a.x1 || a.y2 < b.y1 || b.y2 < a.y1)
			contained := (a.x1 <= b.x1 && a.x2 >= b.x2 && a.y1 <= b.y1 && a.y2 >= b.y2)
			contained = contained || (b.x1 <= a.x1 && b.x2 >= a.x2 && b.y1 <= a.y1 && b.y2 >= a.y2)

			if overlap || contained {
				m[a.name] = true
			}
		}
	}

	for k := range m {
		ret = append(ret, k)
	}

	return ret
}

func (b *Brick) BricksAbove(bricks Bricks) (ret []*Brick) {
	for _, a := range bricks {
		if a == b {
			continue
		}

		if b.z2+1 == a.z1 {
			overlap := !(a.x2 < b.x1 || b.x2 < a.x1 || a.y2 < b.y1 || b.y2 < a.y1)
			contained := (a.x1 <= b.x1 && a.x2 >= b.x2 && a.y1 <= b.y1 && a.y2 >= b.y2)
			contained = contained || (b.x1 <= a.x1 && b.x2 >= a.x2 && b.y1 <= a.y1 && b.y2 >= a.y2)

			if overlap || contained {
				ret = append(ret, a)
			}
		}
	}

	return ret
}

func (b *Brick) BricksBelow(bricks Bricks) (ret []*Brick) {
	for _, a := range bricks {
		if a == b {
			continue
		}

		if b.z1 == a.z2+1 {
			overlap := !(a.x2 < b.x1 || b.x2 < a.x1 || a.y2 < b.y1 || b.y2 < a.y1)
			contained := (a.x1 <= b.x1 && a.x2 >= b.x2 && a.y1 <= b.y1 && a.y2 >= b.y2)
			contained = contained || (b.x1 <= a.x1 && b.x2 >= a.x2 && b.y1 <= a.y1 && b.y2 >= a.y2)

			if overlap || contained {
				ret = append(ret, a)
			}
		}
	}

	return ret
}

type Bricks []*Brick

func (b *Bricks) Sort() {
	for i := 0; i < len(*b); i++ {
		for j := i + 1; j < len(*b); j++ {
			if (*b)[i].z2 > (*b)[j].z2 {
				(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
			}
		}
	}
}

// parse parses
func parse(input string) (ret Bricks) {
	for i, line := range strings.Split(input, "\n") {
		var b Brick
		fmt.Sscanf(strings.TrimSpace(line), "%d,%d,%d~%d,%d,%d", &b.x1, &b.y1, &b.z1, &b.x2, &b.y2, &b.z2)
		ret = append(ret, &b)
		b.name = string(rune('A' + i))
	}

	return ret
}

// partOne returns the answer to part one of this day's puzzle.
func partOne(input string) (ret int) {
	bricks := parse(input)
	bricks.Sort()

	ret = 0

	pq := Bricks{}
	pq = append(pq, bricks...)

	bricks = Bricks{}
	for len(pq) > 0 {
		brick := pq[0]
		pq = pq[1:]

		brick.Drop(bricks)
		bricks = append(bricks, brick)
	}

	for _, b := range bricks {
		ab := b.BricksAbove(bricks)
		for _, a := range ab {
			bb := a.BricksBelow(bricks)
			if len(bb) == 1 && bb[0].name == b.name {
				ret--
				break
			}
		}
		ret++
	}

	return ret
}

// partTwo returns the answer to part one of this day's puzzle.
func partTwo(input string) (ret int) {
	bricks := parse(input)
	bricks.Sort()

	pq := Bricks{}
	pq = append(pq, bricks...)

	bricks = Bricks{}
	for len(pq) > 0 {
		brick := pq[0]
		pq = pq[1:]

		brick.Drop(bricks)
		bricks = append(bricks, brick)
	}

	// fmt.Println(bricks)

	for _, b := range bricks {
		ab := Bricks{}
		for _, a := range bricks {
			if a != b || a.z1 > b.z2 {
				ab = append(ab, &Brick{
					x1: a.x1,
					y1: a.y1,
					z1: a.z1,
					x2: a.x2,
					y2: a.y2,
					z2: a.z2,
				})
			}
		}

		for _, a := range ab {
			if a.Drop(ab) {
				ret += 1
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
