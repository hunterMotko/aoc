package main

import (
	"aoc/utils"
	"fmt"
	"sort"
)

func main() {
	matrix, _ := utils.ReadFileToMatrix("./d12/test1.txt")
	// fmt.Println(part1(matrix))
	fmt.Println(part2(matrix))
}

// Coord represents a point in the m
type Coord struct {
	row, col int
}

var dirs = []Coord{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func part1(m [][]string) int {
	vs := make(map[Coord]bool)
	var res int
	for row := range m {
		for col := range m[row] {
			if !vs[Coord{row, col}] {
				var a, p int
				dfs(Coord{row, col}, m, vs, m[row][col], &a, &p)
				res += (a + 1) * p
			}
		}
	}
	return res
}

func dfs(cur Coord, m [][]string, vs map[Coord]bool, c string, a, p *int) {
	nx := 0
	vs[cur] = true
	for _, dir := range dirs {
		next := Coord{cur.row + dir.row, cur.col + dir.col}
		if isValidCoord(m, next) && m[next.col][next.col] == c {
			nx++
			if !vs[next] {
				(*a)++
				dfs(next, m, vs, c, a, p)
			}
		}
	}
	(*p) += 4 - nx
}

func isValidCoord(m [][]string, c Coord) bool {
	return c.row >= 0 && c.row < len(m) && c.col >= 0 && c.col < len(m[0])
}

const (
	North int = iota
	East
	South
	West
)

type Garden struct {
	ch     string
	size       int
	coords     map[Coord]bool
	boundaries map[int][]boundary
}

// boundary represents an edge of a garden
type boundary struct {
	row, col int
	counted  bool
}

func part2(m [][]string) int {
	gards := findGardens(m)
	return calculateScore(gards)
}

func findGardens(m [][]string) []Garden {
	s := make(map[Coord]bool)
	var gs []Garden
	for row := range m {
		for col := range m[row] {
			coord := Coord{row, col}
			if !s[coord] {
				if g := exploreGarden(m, coord, s); g.size > 0 {
					gs = append(gs, g)
				}
			}
		}
	}
	return gs
}

func exploreGarden(m [][]string, start Coord, seen map[Coord]bool) Garden {
	if seen[start] || !isValidCoord(m, start) {
		return Garden{}
	}
	ch := m[start.row][start.col]
	garden := Garden{
		ch:     ch,
		coords:     make(map[Coord]bool),
		boundaries: make(map[int][]boundary),
	}
	queue := []Coord{start}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if seen[curr] || !isValidCoord(m, curr) || m[curr.row][curr.col] != ch {
			continue
		}
		seen[curr] = true
		garden.coords[curr] = true
		garden.size++
		// Check boundaries in all directions
		checkBoundary(m, curr, ch, North, &garden)
		checkBoundary(m, curr, ch, South, &garden)
		checkBoundary(m, curr, ch, East, &garden)
		checkBoundary(m, curr, ch, West, &garden)
		// Add adjacent cells to queue
		for _, next := range getAdjacent(curr) {
			if isValidCoord(m, next) && m[next.row][next.col] == ch {
				queue = append(queue, next)
			}
		}
	}
	return garden
}

func checkBoundary(m [][]string, pos Coord, ch string, dir int, garden *Garden) {
	var isEdge bool
	switch dir {
	case North:
		isEdge = pos.row == 0 || m[pos.row-1][pos.col] != ch
	case South:
		isEdge = pos.row == len(m)-1 || m[pos.row+1][pos.col] != ch
	case East:
		isEdge = pos.col == len(m[0])-1 || m[pos.row][pos.col+1] != ch
	case West:
		isEdge = pos.col == 0 || m[pos.row][pos.col-1] != ch
	}
	if isEdge {
		garden.boundaries[dir] = append(garden.boundaries[dir], boundary{pos.row, pos.col, true})
	}
}

func calculateScore(gardens []Garden) int {
	total := 0
	for _, garden := range gardens {
		pruneRedundantBoundaries(&garden)
		boundaryCount := countValidBoundaries(garden)
		total += garden.size * boundaryCount
	}
	return total
}

func pruneRedundantBoundaries(garden *Garden) {
	for dir := range garden.boundaries {
		if dir == North || dir == South {
			pruneBoundariesAlongAxis(garden.boundaries[dir], true)
		} else {
			pruneBoundariesAlongAxis(garden.boundaries[dir], false)
		}
	}
}

func pruneBoundariesAlongAxis(bounds []boundary, sortByCol bool) {
	if sortByCol {
		sort.Slice(bounds, func(i, j int) bool { return bounds[i].col < bounds[j].col })
	} else {
		sort.Slice(bounds, func(i, j int) bool { return bounds[i].row < bounds[j].row })
	}
	for i := range bounds {
		var next int
		if sortByCol {
			next = bounds[i].col
		} else {
			next = bounds[i].row
		}
		for {
			next++
			found := false
			for j := range bounds {
				var matches bool
				if sortByCol {
					matches = bounds[j].col == next && bounds[j].row == bounds[i].row
				} else {
					matches = bounds[j].row == next && bounds[j].col == bounds[i].col
				}
				if matches {
					bounds[j].counted = false
					found = true
					break
				}
			}
			if !found {
				break
			}
		}
	}
}

func countValidBoundaries(garden Garden) int {
	count := 0
	for _, bounds := range garden.boundaries {
		for _, b := range bounds {
			if b.counted {
				count++
			}
		}
	}
	return count
}
func getAdjacent(c Coord) []Coord {
	return []Coord{
		{c.row - 1, c.col}, // North
		{c.row + 1, c.col}, // South
		{c.row, c.col + 1}, // East
		{c.row, c.col - 1}, // West
	}
}
