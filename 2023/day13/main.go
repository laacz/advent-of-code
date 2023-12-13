package main

import (
	"fmt"
	"os"
	"strings"
)

type Field struct {
	cols    []uint
	rows    []uint
	revcols []uint
	revrows []uint
	w, h    int

	left, top int
}

func (f Field) String() (ret string) {
	for y, row := range f.rows {
		if f.top > 0 && f.top == y {
			ret += strings.Repeat("─", f.w) + "\n"
		}
		for x := f.w - 1; x >= 0; x-- {
			if f.left > 0 && f.left == f.w-x-1 {
				ret += "│"
			}

			if row&(1<<x) != 0 {
				ret += "#"
			} else {
				ret += "."
			}
		}
		ret += "\n"
	}

	return ret
}

// FindMirrorline finds the mirrorline on a field
func (f *Field) FindMirrorline(left, top int) int {
	f.top = 0
	f.left = 0

	for i := 1; i < f.h; i++ {
		if top != i && partialeq(reverse(f.rows[:i]), f.rows[i:]) {
			f.top = i
			return f.top * 100
		}
	}

	for i := 1; i < f.w; i++ {
		if left != i && partialeq(reverse(f.cols[:i]), f.cols[i:]) {
			f.left = i
			return f.left
		}
	}

	return 0
}

// min returns the smaller of a and b.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// partialeq returns true if the a starts with b or vice versa.
func partialeq(a, b []uint) bool {
	for i := 0; i < min(len(a), len(b)); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// reverse returns a copy of a with the order of elements reversed.
func reverse(a []uint) []uint {
	ret := make([]uint, len(a))
	copy(ret, a)

	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}

	return ret
}

func parse(input string) (ret []Field) {
	for _, field := range strings.Split(input, "\n\n") {
		f := Field{}
		rows := strings.Split(field, "\n")
		f.w = len(rows[0])

		for _, row := range strings.Split(field, "\n") {
			var r uint
			for _, c := range row {
				r <<= 1
				if c == '#' {
					r |= 1
				}
			}
			f.rows = append(f.rows, r)
		}

		f.h = len(f.rows)

		for i := f.w - 1; i >= 0; i-- {
			var c uint
			for j := 0; j < f.h; j++ {
				c <<= 1
				if f.rows[j]&(1<<i) != 0 {
					c |= 1
				}
			}
			f.cols = append(f.cols, c)
		}

		f.revrows = reverse(f.rows)
		f.revcols = reverse(f.cols)

		ret = append(ret, f)

	}

	return ret
}

// partOne returns the answer to part one of this day's puzzle.
func partOne(input string) (ret int) {
	for _, field := range parse(input) {
		ret += field.FindMirrorline(0, 0)
	}
	return ret
}

// partTwo returns the answer to part one of this day's puzzle.
func partTwo(input string) (ret int) {
	for _, field := range parse(input) {
		var f Field
		field.FindMirrorline(0, 0)

		mirrors := make(map[int]bool)
		for y := 0; y < field.h; y++ {
			for x := 0; x < field.w; x++ {
				f = Field{
					w:    field.w,
					h:    field.h,
					cols: make([]uint, field.w),
					rows: make([]uint, field.h),
				}

				copy(f.cols, field.cols)
				copy(f.rows, field.rows)

				f.rows[y] ^= 1 << (field.w - x - 1)
				f.cols[x] ^= 1 << y

				f.FindMirrorline(field.left, field.top)

				if f.top+f.left != 0 && (f.top != field.top || f.left != field.left) {
					// there can be more than one, can there?
					mirrors[f.top*100+f.left] = true
				}
			}
		}

		for k := range mirrors {
			ret += k
		}
	}
	return ret
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
