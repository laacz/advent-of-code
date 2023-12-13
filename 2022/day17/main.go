package main

import (
	"fmt"
	"os"
	"strconv"
)

type Block struct {
	Shape []byte
	Y     int
}

type Game struct {
	Block           *Block
	BlockIdx        int
	Board           *[]byte
	AvailableBlocks []*Block
	Movements       string
	MoveIdx         int
	StoppedBlocks   uint64
	Height          uint64
	PatternFound    bool
	Memory          map[string][]uint64
}

// SpawnBlock generates a new block, if previous one has been stopped
func (g *Game) SpawnBlock() bool {
	if g.Block != nil {
		return false
	}

	// Update the next block's index
	g.BlockIdx = (g.BlockIdx + 1) % (len(g.AvailableBlocks))

	// Find first empty line
	fel := 0
	for i := 0; i < len(*g.Board); i++ {
		if (*g.Board)[i] == 0 {
			fel = i
			break
		}
	}

	// Ensure top 7 lines ar empty
	for len(*g.Board) < fel+8 {
		*g.Board = append(*g.Board, 0)
	}

	// Add the block to the board
	shp := g.AvailableBlocks[g.BlockIdx].Shape
	g.Block = &Block{}
	g.Block.Y = fel + 3
	g.Block.Shape = make([]byte, len(shp))
	// g.Block.Shape = g.AvailableBlocks[g.NextBlockIdx%(len(g.AvailableBlocks))].Shape

	copy(g.Block.Shape, shp)
	// g.Block = g.AvailableBlocks[g.NextBlockIdx%(len(g.AvailableBlocks))]

	return true
}

func shft(b byte, n byte) byte {
	if n == '<' {
		return b << 1
	} else {
		return b >> 1
	}
}

// Slide moves block sideways, if possible
func (g *Game) Slide() {
	g.MoveIdx = (g.MoveIdx + 1) % len(g.Movements)
	idx := g.MoveIdx
	mv := g.Movements[idx]
	shp := g.Block.Shape
	movable := true

	for i := 0; i < len(shp); i++ {
		// Check if we're at the edge
		edge := (shp[i]&0b0000001 == 1 && mv == '>') || (shp[i]&0b1000000 == 0b1000000 && mv == '<')
		if edge {
			movable = false
			break
		}

		// Check if we anything else blocks us
		nb := shft(shp[i], mv)
		l := (*g.Board)[g.Block.Y+len(shp)-i-1]
		if nb^l != nb|l {
			movable = false
			break
		}
	}

	if movable {
		for i := 0; i < len(shp); i++ {
			switch mv {
			case '<':
				g.Block.Shape[i] = shp[i] << 1
			case '>':
				g.Block.Shape[i] = shp[i] >> 1
			}
		}
	}

}

// CanFall returns true if the block can fall one line
func (g *Game) CanFall() bool {
	y := g.Block.Y - 1
	ret := true

	if y == -1 {
		ret = false
	} else {
		for i := 0; i < len(g.Block.Shape); i++ {
			bl := g.Block.Shape[len(g.Block.Shape)-i-1]
			br := (*g.Board)[y+i]
			if bl^br != bl|br {
				ret = false
				break
			}
		}
	}

	return ret
}

// BoardHeight returns the actual board height, ignoring padded lines
func (g *Game) BoardHeight() uint64 {
	for i := len(*g.Board) - 1; i >= 0; i-- {
		if (*g.Board)[i] != 0 {
			return uint64(i)
		}
	}
	return 0
}

// StopBlock drops the block and renders it on the board
func (g *Game) StopBlock() {
	y := g.Block.Y
	for i := 0; i < len(g.Block.Shape); i++ {
		(*g.Board)[y+i] |= g.Block.Shape[len(g.Block.Shape)-i-1]
	}

	g.Block = nil
	g.StoppedBlocks += 1
}

// Tick advances the game by one step (spawn block, if needed, slide, then fall and/or stop)
func (g *Game) Tick() uint64 {
	blockNum := g.StoppedBlocks
	if g.SpawnBlock() {
		// I worked on this for far too long, but eventually here we are
		// The essence is just to check if current interval of cycle is
		// already memorized.
		idx := fmt.Sprintf("%d,%d", g.BlockIdx, g.MoveIdx)
		mem, ok := g.Memory[idx]
		if ok {
			prevBlockNum := mem[0]
			prevHeight := mem[1]
			interval := blockNum - prevBlockNum
			if blockNum%interval == 1_000_000_000_000%interval {
				intervalHeight := g.BoardHeight() - prevHeight
				blocksLeft := 1_000_000_000_000 - blockNum
				intervalsLeft := blocksLeft/interval + 1
				return prevHeight + intervalHeight*intervalsLeft + 1
			}
		} else {
			g.Memory[idx] = []uint64{blockNum, g.BoardHeight()}
		}
	}

	g.Slide()
	if !g.CanFall() {
		g.StopBlock()
	} else {
		g.Block.Y -= 1
	}

	return 0
}

func NewGame(input string) *Game {
	g := &Game{
		Board:     &[]byte{},
		BlockIdx:  -1,
		MoveIdx:   -1,
		Movements: input,
		AvailableBlocks: []*Block{
			{
				Shape: []byte{
					0b0011110,
				},
			},
			{
				Shape: []byte{
					0b0001000,
					0b0011100,
					0b0001000,
				},
			},
			{
				Shape: []byte{
					0b0000100,
					0b0000100,
					0b0011100,
				},
			},
			{
				Shape: []byte{
					0b0010000,
					0b0010000,
					0b0010000,
					0b0010000,
				},
			},
			{
				Shape: []byte{
					0b0011000,
					0b0011000,
				},
			},
		},
		Memory: make(map[string][]uint64),
	}

	return g
}

func partOne(input string) int {
	ret := 0
	game := NewGame(input)
	i := 0

	z, err := strconv.Atoi(os.Args[len(os.Args)-1])
	var stopped uint64 = 2022
	if err == nil {
		stopped = uint64(z)
	}

	for {
		i++
		game.Tick()
		if game.StoppedBlocks == stopped {
			ret = int(game.BoardHeight())
			break
		}
	}

	return ret + 1
}

func partTwo(input string) uint64 {
	ret := uint64(0)
	game := NewGame(input)
	i := 0

	for {
		i++
		estimatedHeight := game.Tick()
		if estimatedHeight > 0 {
			ret = estimatedHeight
			break
		}
	}

	return ret
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(fmt.Errorf("while reading %s, got %v", "input.txt", err))
	}

	fmt.Printf("Part one: %d\n", partOne(string(data)))
	fmt.Printf("Part two: %d\n", partTwo(string(data)))
}
