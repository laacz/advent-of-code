package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Tile struct {
	id    int
	grid  []string
	edges [4]string
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Part 1:", part1(string(data)))
	fmt.Println("Part 2:", part2(string(data)))
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func parse(input string) []Tile {
	blocks := strings.Split(strings.TrimSpace(input), "\n\n")
	var tiles []Tile

	for _, block := range blocks {
		lines := strings.Split(block, "\n")

		idStr := strings.TrimPrefix(lines[0], "Tile ")
		idStr = strings.TrimSuffix(idStr, ":")
		id, _ := strconv.Atoi(idStr)

		grid := lines[1:]

		top := grid[0]
		bottom := grid[len(grid)-1]

		var right, left strings.Builder
		for _, row := range grid {
			right.WriteByte(row[len(row)-1])
			left.WriteByte(row[0])
		}

		tiles = append(tiles, Tile{
			id:    id,
			grid:  grid,
			edges: [4]string{top, right.String(), bottom, left.String()},
		})
	}

	return tiles
}

func part1(input string) int {
	tiles := parse(input)

	edgeCount := make(map[string]int)

	for _, tile := range tiles {
		for _, edge := range tile.edges {
			rev := reverse(edge)
			canonical := edge
			if rev < edge {
				canonical = rev
			}
			edgeCount[canonical]++
		}
	}

	result := 1
	for _, tile := range tiles {
		uniqueEdges := 0
		for _, edge := range tile.edges {
			rev := reverse(edge)
			canonical := edge
			if rev < edge {
				canonical = rev
			}
			if edgeCount[canonical] == 1 {
				uniqueEdges++
			}
		}
		if uniqueEdges == 2 {
			result *= tile.id
		}
	}

	return result
}

func rotateGrid(grid []string) []string {
	n := len(grid)
	result := make([]string, n)
	for i := 0; i < n; i++ {
		var sb strings.Builder
		for j := n - 1; j >= 0; j-- {
			sb.WriteByte(grid[j][i])
		}
		result[i] = sb.String()
	}
	return result
}

func flipGrid(grid []string) []string {
	result := make([]string, len(grid))
	for i, row := range grid {
		result[i] = reverse(row)
	}
	return result
}

func getEdges(grid []string) [4]string {
	top := grid[0]
	bottom := grid[len(grid)-1]
	var right, left strings.Builder
	for _, row := range grid {
		right.WriteByte(row[len(row)-1])
		left.WriteByte(row[0])
	}
	return [4]string{top, right.String(), bottom, left.String()}
}

func getAllOrientations(grid []string) [][]string {
	var orientations [][]string
	current := grid
	for i := 0; i < 4; i++ {
		orientations = append(orientations, current)
		orientations = append(orientations, flipGrid(current))
		current = rotateGrid(current)
	}
	return orientations
}

func canonical(edge string) string {
	rev := reverse(edge)
	if rev < edge {
		return rev
	}
	return edge
}

func findCorners(tiles []Tile, edgeCount map[string]int) []Tile {
	var corners []Tile
	for _, tile := range tiles {
		uniqueEdges := 0
		for _, edge := range tile.edges {
			if edgeCount[canonical(edge)] == 1 {
				uniqueEdges++
			}
		}
		if uniqueEdges == 2 {
			corners = append(corners, tile)
		}
	}
	return corners
}

func assembleTiles(input string) [][]Tile {
	tiles := parse(input)
	size := int(math.Sqrt(float64(len(tiles))))

	edgeCount := make(map[string]int)
	for _, tile := range tiles {
		for _, edge := range tile.edges {
			edgeCount[canonical(edge)]++
		}
	}

	edgeToTiles := make(map[string][]int)
	for i, tile := range tiles {
		for _, edge := range tile.edges {
			c := canonical(edge)
			edgeToTiles[c] = append(edgeToTiles[c], i)
		}
	}

	corners := findCorners(tiles, edgeCount)

	startTile := corners[0]
	startGrid := startTile.grid

	for _, oriented := range getAllOrientations(startGrid) {
		edges := getEdges(oriented)
		topUnique := edgeCount[canonical(edges[0])] == 1
		leftUnique := edgeCount[canonical(edges[3])] == 1
		if topUnique && leftUnique {
			startGrid = oriented
			break
		}
	}

	grid := make([][]Tile, size)
	for i := range grid {
		grid[i] = make([]Tile, size)
	}

	used := make(map[int]bool)
	grid[0][0] = Tile{id: startTile.id, grid: startGrid, edges: getEdges(startGrid)}
	used[startTile.id] = true

	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if row == 0 && col == 0 {
				continue
			}

			var needLeft, needTop string
			if col > 0 {
				needLeft = grid[row][col-1].edges[1]
			}
			if row > 0 {
				needTop = grid[row-1][col].edges[2]
			}

			for _, tile := range tiles {
				if used[tile.id] {
					continue
				}

				for _, oriented := range getAllOrientations(tile.grid) {
					edges := getEdges(oriented)

					leftMatch := col == 0 || edges[3] == needLeft
					topMatch := row == 0 || edges[0] == needTop

					if leftMatch && topMatch {
						grid[row][col] = Tile{id: tile.id, grid: oriented, edges: edges}
						used[tile.id] = true
						goto found
					}
				}
			}
		found:
		}
	}

	return grid
}

func buildImage(tileGrid [][]Tile) []string {
	tileSize := len(tileGrid[0][0].grid)

	var image []string
	for row := range tileGrid {
		for y := 1; y < tileSize-1; y++ {
			var line string
			for col := range tileGrid[row] {
				line += tileGrid[row][col].grid[y][1 : tileSize-1]
			}
			image = append(image, line)
		}
	}
	return image
}

var seaMonster = []string{
	"                  # ",
	"#    ##    ##    ###",
	" #  #  #  #  #  #   ",
}

func countSeaMonsters(image []string) int {
	count := 0

	for y := 0; y <= len(image)-3; y++ {
		for x := 0; x <= len(image[0])-20; x++ {
			if matchesMonster(image, x, y) {
				count++
			}
		}
	}
	return count
}

func matchesMonster(image []string, startX, startY int) bool {
	for dy, row := range seaMonster {
		for dx, ch := range row {
			if ch == '#' && image[startY+dy][startX+dx] != '#' {
				return false
			}
		}
	}
	return true
}

func rotateImage(img []string) []string {
	n := len(img)
	result := make([]string, n)
	for i := 0; i < n; i++ {
		var sb strings.Builder
		for j := n - 1; j >= 0; j-- {
			sb.WriteByte(img[j][i])
		}
		result[i] = sb.String()
	}
	return result
}

func flipImage(img []string) []string {
	result := make([]string, len(img))
	for i, row := range img {
		result[i] = reverse(row)
	}
	return result
}

func getAllImageOrientations(img []string) [][]string {
	var orientations [][]string
	current := img
	for i := 0; i < 4; i++ {
		orientations = append(orientations, current)
		orientations = append(orientations, flipImage(current))
		current = rotateImage(current)
	}
	return orientations
}

func countHashes(image []string) int {
	count := 0
	for _, row := range image {
		for _, ch := range row {
			if ch == '#' {
				count++
			}
		}
	}
	return count
}

func monsterHashes() int {
	count := 0
	for _, row := range seaMonster {
		for _, ch := range row {
			if ch == '#' {
				count++
			}
		}
	}
	return count
}

func markSeaMonsters(image []string) []string {

	grid := make([][]byte, len(image))
	for i, row := range image {
		grid[i] = []byte(row)
	}

	for y := 0; y <= len(image)-3; y++ {
		for x := 0; x <= len(image[0])-20; x++ {
			if matchesMonster(image, x, y) {

				for dy, row := range seaMonster {
					for dx, ch := range row {
						if ch == '#' {
							grid[y+dy][x+dx] = 'O'
						}
					}
				}
			}
		}
	}

	result := make([]string, len(grid))
	for i, row := range grid {
		result[i] = string(row)
	}
	return result
}

func printImageWithMonsters(image []string) {
	grey := "\033[90m"
	silver := "\033[37m"
	silverBg := "\033[47m"
	black := "\033[30m"
	reset := "\033[0m"

	fmt.Println()
	for _, row := range image {
		for _, ch := range row {
			switch ch {
			case '#':
				fmt.Print(silver + "#" + reset)
			case 'O':
				fmt.Print(silverBg + black + "O" + reset)
			default:
				fmt.Print(grey + "·" + reset)
			}
		}
		fmt.Println()
	}
}

func part2(input string) int {
	tileGrid := assembleTiles(input)
	image := buildImage(tileGrid)

	for _, oriented := range getAllImageOrientations(image) {
		monsters := countSeaMonsters(oriented)
		if monsters > 0 {
			marked := markSeaMonsters(oriented)
			printImageWithMonsters(marked)

			totalHashes := countHashes(oriented)
			monsterHashCount := monsterHashes() * monsters
			return totalHashes - monsterHashCount
		}
	}

	return 0
}
