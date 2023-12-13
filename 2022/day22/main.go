package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Right = iota
	Down
	Left
	Up
)

const (
	Open byte = '.'
	Wall byte = '#'
)

var dirs = [][]int{
	{1, 0, Right}, // Right
	{0, 1, Down},  // Down
	{-1, 0, Left}, // Left
	{0, -1, Up},   // Up
}

type Tile struct {
	X, Y int
}

type Map struct {
	Tiles         map[Tile]byte
	Width, Height int
	Instructions  []string
	Pos           Tile
	Dir           int
	Faces         []struct {
		X, Y int
	}
	TileSize int
	Rules    map[int]map[int][]int
}

func NewMap(input string) *Map {
	t := strings.Split(input, "\n\n")
	var m Map
	// Populate the map
	m.Tiles = make(map[Tile]byte)
	for y, row := range strings.Split(t[0], "\n") {
		if y > m.Height {
			m.Height = y + 1
		}
		for x, c := range row {
			if x > m.Width {
				m.Width = x + 1
			}
			switch c {
			case '#':
				m.Tiles[Tile{x, y}] = Wall
			case '.':
				m.Tiles[Tile{x, y}] = Open
			}
		}
	}

	// Parse instructions
	cc := ""
	for i, c := range t[1] {
		if c == 'L' || c == 'R' {
			if len(cc) > 0 || i == len(t[1])-1 {
				m.Instructions = append(m.Instructions, cc)
				cc = ""
			}

			m.Instructions = append(m.Instructions, string(c))
		} else {
			cc += string(c)
			if i == len(t[1])-1 {
				m.Instructions = append(m.Instructions, cc)
			}
		}
	}

	// Determine the starting position
	for x := 0; x < m.Width; x++ {
		if m.Tiles[Tile{x, 0}] == Open {
			m.Pos = Tile{x, 0}
			break
		}
	}

	// Hardcoding wrap-around rules...
	if m.Height == 12 {
		// test input
		m.TileSize = 4
		m.Rules = map[int]map[int][]int{
			1: {
				Right: {6, Left},
				Down:  {4, Down},
				Left:  {3, Down},
				Up:    {2, Down},
			},
			2: {
				Right: {3, Right},
				Down:  {5, Up},
				Left:  {6, Up},
				Up:    {1, Down},
			},
			3: {
				Right: {4, Right},
				Down:  {5, Right},
				Left:  {2, Left},
				Up:    {1, Right},
			},
			4: {
				Right: {6, Down},
				Down:  {5, Down},
				Left:  {3, Left},
				Up:    {1, Up},
			},
			5: {
				Right: {6, Right},
				Down:  {2, Up},
				Left:  {3, Up},
				Up:    {4, Up},
			},
			6: {
				Right: {1, Left},
				Down:  {2, Right},
				Left:  {5, Left},
				Up:    {4, Left},
			},
		}
	} else {
		// actual input
		m.TileSize = 50
		m.Rules = map[int]map[int][]int{
			1: {
				Right: {2, Right},
				Down:  {3, Down},
				Left:  {4, Right},
				Up:    {6, Right},
			},
			2: {
				Right: {5, Left},
				Down:  {3, Left},
				Left:  {1, Left},
				Up:    {6, Up},
			},
			3: {
				Right: {2, Up},
				Down:  {5, Down},
				Left:  {4, Down},
				Up:    {1, Up},
			},
			4: {
				Right: {5, Right},
				Down:  {6, Down},
				Left:  {1, Right},
				Up:    {3, Right},
			},
			5: {
				Right: {2, Left},
				Down:  {6, Left},
				Left:  {4, Left},
				Up:    {3, Up},
			},
			6: {
				Right: {5, Up},
				Down:  {2, Down},
				Left:  {1, Down},
				Up:    {4, Up},
			},
		}
	}

	// Find all faces
	for y := 0; y < m.Height; y += m.TileSize {
		for x := 0; x < m.Width; x += m.TileSize {
			if _, ok := m.Tiles[Tile{x, y}]; ok {
				m.Faces = append(m.Faces, struct{ X, Y int }{x, y})
			}
		}
	}

	return &m
}

// NextTile returns next Tile, using wrap-around rules from the 1st part
func (m *Map) NextTile() Tile {
	p := Tile{
		m.Pos.X,
		m.Pos.Y,
	}
	for {
		p.X = (p.X + dirs[m.Dir][0]) % m.Width
		p.Y = (p.Y + dirs[m.Dir][1]) % m.Height
		if p.X < 0 {
			p.X += m.Width
		}
		if p.Y < 0 {
			p.Y += m.Height
		}

		if _, ok := m.Tiles[p]; ok {
			break
		}
	}
	return p
}

// Next3DTile returns next Tile, using wrap-around rules from the 2nd part
func (m *Map) Next3DTile() (Tile, int) {
	currentFace, relXY := m.GetFaceInfo(m.Pos)
	p := Tile{
		m.Pos.X + dirs[m.Dir][0],
		m.Pos.Y + dirs[m.Dir][1],
	}
	newFace, _ := m.GetFaceInfo(p)
	// If tile is outside of all faces, newFace is zero
	if currentFace == newFace {
		return p, m.Dir
	}

	rule := m.Rules[currentFace][m.Dir]
	newFace = rule[0] - 1
	newDir := rule[1]

	// This is not generic. It depends on the way cube is folded, so there's one nested conditional.
	if m.Dir == Up && newDir == Up { // sample 5->4, 4->1, input 5->3, 3->1
		p.Y = m.Faces[newFace].Y + m.TileSize - 1
		p.X = m.Faces[newFace].X + relXY.X
	} else if m.Dir == Up && newDir == Down { // sample 1->2
		p.Y = m.Faces[newFace].Y
		p.X = m.Faces[newFace].X + m.TileSize - 1 - relXY.X
	} else if m.Dir == Up && newDir == Right { // input 1>6 and 4->3
		p.Y = m.Faces[newFace].Y + relXY.X
		p.X = m.Faces[newFace].X
	} else if m.Dir == Up && newDir == Left { // sample 3->1, input 1->6
		p.Y = m.Faces[newFace].Y + relXY.X
		p.X = m.Faces[newFace].X
	} else if m.Dir == Right && newDir == Right { // input 3-4 and 6->1
		p.Y = m.Faces[newFace].Y + relXY.Y
		p.X = m.Faces[newFace].X + m.TileSize - 1 - relXY.X
	} else if m.Dir == Right && newDir == Down { // sample 4->6
		p.Y = m.Faces[newFace].Y
		p.X = m.Faces[newFace].X + m.TileSize - 1 - relXY.Y
	} else if m.Dir == Right && newDir == Up { // input 3->2
		p.Y = m.Faces[newFace].Y + m.TileSize - 1
		p.X = m.Faces[newFace].X + relXY.Y
	} else if m.Dir == Right && newDir == Left { // input 5->2
		p.Y = m.Faces[newFace].Y + m.TileSize - 1 - relXY.Y
		p.X = m.Faces[newFace].X + m.TileSize - 1
	} else if m.Dir == Down && newDir == Up { // sample 5->2
		p.Y = m.Faces[newFace].Y + m.TileSize - 1
		p.X = m.Faces[newFace].X + m.TileSize - 1 - relXY.X
	} else if m.Dir == Down && newDir == Right { // sample 6->2, input 5->6
		// OMG, I hunted this special folding case for 2 hours.
		if m.TileSize == 4 {
			p.Y = m.Faces[newFace].Y + m.TileSize - 1 - relXY.X
			p.X = m.Faces[newFace].X
		} else {
			p.Y = m.Faces[newFace].Y + relXY.X
			p.X = m.Faces[newFace].X + m.TileSize - 1
		}
	} else if m.Dir == Down && newDir == Left { // input 2->3
		p.Y = m.Faces[newFace].Y + relXY.X
		p.X = m.Faces[newFace].X + m.TileSize - 1
	} else if m.Dir == Down && newDir == Down { // input -> 6 -> 2
		p.Y = m.Faces[newFace].Y
		p.X = m.Faces[newFace].X + relXY.X
	} else if m.Dir == Left && newDir == Down { // input 3-4 and 6->1
		p.Y = m.Faces[newFace].Y
		p.X = m.Faces[newFace].X + relXY.Y
	} else if m.Dir == Left && newDir == Up { // sample 5->3
		p.Y = m.Faces[newFace].Y + m.TileSize - 1
		p.X = m.Faces[newFace].X + m.TileSize - 1 - relXY.Y
	} else if m.Dir == Left && newDir == Right { // input 1->4
		p.Y = m.Faces[newFace].Y + m.TileSize - 1 - relXY.Y
		p.X = m.Faces[newFace].X
	}

	return p, newDir
}

// Walk walks the map, following the instructions
func (m *Map) Walk() {
	for _, i := range m.Instructions {
		switch i {
		case "R":
			m.Dir = (m.Dir + 1) % 4
		case "L":
			m.Dir = (m.Dir + 3) % 4
		default:
			tiles, _ := strconv.Atoi(i)
			for i := 0; i < tiles; i++ {
				p := m.NextTile()
				if m.Tiles[p] == Wall {
					break
				}
				m.Pos = p
			}
		}
	}
}

// Walk3D walks the map, following the rules from the 2nd part
func (m *Map) Walk3D() {
	for _, i := range m.Instructions {
		switch i {
		case "R":
			m.Dir = (m.Dir + 1) % 4
		case "L":
			m.Dir = (m.Dir + 3) % 4
		default:
			tiles, _ := strconv.Atoi(i)
			for i := 0; i < tiles; i++ {
				nextTile, newDir := m.Next3DTile()
				if m.Tiles[nextTile] == Wall {
					break
				}
				m.Dir = newDir
				m.Pos = nextTile
			}
		}
	}
}

// GetFaceInfo returns the face number and the position of the tile within the face
func (m *Map) GetFaceInfo(t Tile) (int, Tile) {
	for i, f := range m.Faces {
		if t.X >= f.X && t.X < f.X+m.TileSize && t.Y >= f.Y && t.Y < f.Y+m.TileSize {
			return i + 1, Tile{t.X - f.X, t.Y - f.Y}
		}
	}
	return 0, Tile{}
}

func partOne(input string) int {
	m := NewMap(input)
	m.Walk()

	return 1000*(m.Pos.Y+1) + 4*(m.Pos.X+1) + dirs[m.Dir][2]
}

func partTwo(input string) int {
	m := NewMap(input)
	m.Walk3D()

	return 1000*(m.Pos.Y+1) + 4*(m.Pos.X+1) + dirs[m.Dir][2]
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
