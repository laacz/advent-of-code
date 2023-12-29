package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
)

type Hailstone struct {
	x, y, z    float64
	dx, dy, dz float64
}

// Intersection returns true if the two Hailstones intersect
func (a Hailstone) Intersection(b Hailstone) (x, y float64, future bool) {
	// Solving X and Y for general line equation: Y = m * X + b, where X and Y must be the same for
	// given line (am * X + bm = bm * X + bb)

	am := a.dy / a.dx
	ab := a.y - am*a.x
	bm := b.dy / b.dx
	bb := b.y - bm*b.x

	X := (bb - ab) / (am - bm)
	Y := am*X + ab

	// Check if the intersection is in the future
	// If the intersection is in the future, then the difference between the
	// X and Y coordinates of the two Hailstones will have the same sign as the
	// velocity of the Hailstone.

	future = (X-a.x)*(a.dx) > 0 && (X-b.x)*(b.dx) > 0 && (Y-a.y)*(a.dy) > 0 && (Y-b.y)*(b.dy) > 0

	return X, Y, future
}

type Hailstones []Hailstone

// Intersections find the number of intersections between the Hailstones in the given coords range
func (h Hailstones) Intersections(m, n int) (ret int) {
	var seen = make(map[[2]Hailstone]bool)
	for _, a := range h {
		for _, b := range h {
			if a == b || seen[[2]Hailstone{a, b}] || seen[[2]Hailstone{b, a}] {
				continue
			}
			x, y, future := a.Intersection(b)
			if !math.IsInf(x, 0) &&
				x >= float64(m) && x <= float64(n) &&
				y >= float64(m) && y <= float64(n) &&
				future {
				ret++
			}
			seen[[2]Hailstone{a, b}] = true
		}
	}

	return ret
}

// parse parses
func parse(input string) (ret Hailstones) {
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		var h Hailstone
		fmt.Sscanf(line, "%f, %f, %f @ %f, %f, %f", &h.x, &h.y, &h.z, &h.dx, &h.dy, &h.dz)
		ret = append(ret, h)
	}

	return ret
}

// partOne returns the answer to part one of this day's puzzle.
func partOne(input string, m, n int) (ret int) {
	h := parse(input)
	ret = h.Intersections(m, n)

	return ret
}

// From Rosetta Code
func GaussPartial(a0 [][]float64, b0 []float64) ([]float64, error) {
	m := len(b0)
	a := make([][]float64, m)
	for i, ai := range a0 {
		row := make([]float64, m+1)
		copy(row, ai)
		row[m] = b0[i]
		a[i] = row
	}
	for k := range a {
		iMax := 0
		max := -1.
		for i := k; i < m; i++ {
			row := a[i]
			// compute scale factor s = max abs in row
			s := -1.
			for j := k; j < m; j++ {
				x := math.Abs(row[j])
				if x > s {
					s = x
				}
			}
			// scale the abs used to pick the pivot.
			if abs := math.Abs(row[k]) / s; abs > max {
				iMax = i
				max = abs
			}
		}
		if a[iMax][k] == 0 {
			return nil, errors.New("singular")
		}
		a[k], a[iMax] = a[iMax], a[k]
		for i := k + 1; i < m; i++ {
			for j := k + 1; j <= m; j++ {
				a[i][j] -= a[k][j] * (a[i][k] / a[k][k])
			}
			a[i][k] = 0
		}
	}
	x := make([]float64, m)
	for i := m - 1; i >= 0; i-- {
		x[i] = a[i][m]
		for j := i + 1; j < m; j++ {
			x[i] -= a[i][j] * x[j]
		}
		x[i] /= a[i][i]
	}
	return x, nil
}

// GetMatrixes returns two matrices - coefficients and
func (h Hailstones) GetMatrixes(z bool) (coeffs [][]float64, rhs []float64) {
	for i := 0; i < 4; i++ {
		a, b := h[i], h[i+1]
		if z {
			a.y, a.dy = a.z, a.dz
			b.y, b.dy = b.z, b.dz
		}
		row := []float64{
			b.dy - a.dy,
			a.dx - b.dx,
			a.y - b.y,
			b.x - a.x,
		}

		coeffs = append(coeffs, row)
		rhs = append(rhs, b.x*b.dy-b.y*b.dx-a.x*a.dy+a.y*a.dx)
	}

	return coeffs, rhs
}

// partTwo returns the answer to part one of this day's puzzle.
func partTwo(input string) (ret int) {
	h := parse(input)

	coeffs, rhs := h.GetMatrixes(false)

	gp, err := GaussPartial(coeffs, rhs)
	if err != nil {
		panic(err)
	}

	x, y := gp[0], gp[1]

	coeffs, rhs = h.GetMatrixes(true)

	gp, err = GaussPartial(coeffs, rhs)
	if err != nil {
		panic(err)
	}

	z := gp[1]

	return int(math.Round(x + y + z))
}

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Printf("Part one: %d\n", partOne(string(data), 200000000000000, 400000000000000))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
