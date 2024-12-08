package main

import (
	"aoc/utils"
	"errors"
	"fmt"
	"log"
)

func main() {
	matrix, err := utils.ReadFileToMatrix("./d6/in.txt")
	if err != nil {
		log.Fatalf("PARSE FILE ERROR:  %v\n", err)
	}
	// fmt.Println(part1(matrix))
	var g Grid
	g = matrix
	visited, _ := g.traverse()
	fmt.Println(part2(&g, visited))
}

type Pos struct {
	x   int
	y   int
	dir string
}

type Position struct {
	x int
	y int
}
type Dir struct {
	dx int
	dy int
}

var (
	DOWN  Dir    = Dir{0, -1}
	UP    Dir    = Dir{0, 1}
	LEFT  Dir    = Dir{-1, 0}
	RIGHT Dir    = Dir{1, 0}
	OBJ   string = "#"
	START string = "^"
	EMPTY string = "."
)

type Visit struct {
	pos Position
	dir Dir
}

func (v *Visit) step() Visit {
	return Visit{
		Position{v.pos.x + v.dir.dx, v.pos.y + v.dir.dy},
		v.dir,
	}
}
func (v *Visit) rotate() Visit {
	dir := v.dir
	switch dir {
	case DOWN:
		dir = RIGHT
	case UP:
		dir = LEFT
	case LEFT:
		dir = DOWN
	case RIGHT:
		dir = UP
	}
	return Visit{v.pos, dir}
}

type Grid [][]string
type Row []string
type PosSet map[Position]bool

func (g *Grid) start() (*Visit, error) {
	for y, row := range *g {
		for x, cell := range row {
			if cell == START {
				return &Visit{Position{x, y}, DOWN}, nil
			}
		}
	}
	return nil, errors.New("START ERROR")
}
func (g *Grid) outside(pos *Position) bool {
	if pos.x < 0 || pos.x >= len((*g)[0]) {
		return true
	} else if pos.y < 0 || pos.y >= len(*g) {
		return true
	}
	return false
}
func (g *Grid) step(visit *Visit) (*Visit, bool) {
	nextV := visit.step()
	nextP := nextV.pos
	if g.outside(&nextP) {
		return nil, true
	}
	for (*g)[nextP.y][nextP.x] == OBJ {
		nextV = visit.rotate()
		nextP = nextV.pos
	}
	return &nextV, false
}
func (g *Grid) traverse() (*PosSet, bool) {
	visited := PosSet{}
	visits := map[Visit]bool{}
	visit, _ := g.start()
	var exit bool
	for visit != nil {
		if visits[*visit] {
			break
		}
		visited[visit.pos] = true
		visits[*visit] = true
		visit, exit = g.step(visit)
	}
	return &visited, exit
}
func (g *Grid) createObj(pos *Position) {
	(*g)[pos.y][pos.x] = OBJ
}
func (g *Grid) removeObj(pos *Position) {
	(*g)[pos.y][pos.x] = EMPTY
}

func part2(g *Grid, visited *PosSet) int {
	loops := 0
	start, _ := g.start()
	delete(*visited, start.pos)
	for pos := range *visited {
		g.createObj(&pos)
		_, exit := g.traverse()
		g.removeObj(&pos)
		if !exit {
			loops++
		}
	}
	return loops
}

func part1(matrix [][]string) int {
	pos := findStart(matrix)
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	for (pos.x > top && pos.x < bottom) && (pos.y > left && pos.y < right) {
		if pos.dir == "^" {
			for up := pos.x; up >= top; up-- {
				cur := matrix[up][pos.y]
				if cur == "." {
					matrix[up][pos.y] = "X"
					pos.x = up
				} else if cur == "#" {
					pos.dir = ">"
					pos.x = up + 1
					break
				}
			}
		} else if pos.dir == ">" {
			for rg := pos.y; rg <= right; rg++ {
				cur := matrix[pos.x][rg]
				if cur == "." {
					matrix[pos.x][rg] = "X"
					pos.y = rg
				} else if cur == "#" {
					pos.dir = "v"
					pos.y = rg - 1
					break
				}
			}
		} else if pos.dir == "<" {
			for lf := pos.y; lf >= left; lf-- {
				cur := matrix[pos.x][lf]
				if cur == "." {
					matrix[pos.x][lf] = "X"
					pos.y = lf
				} else if cur == "#" {
					pos.dir = "^"
					pos.y = lf + 1
					break
				}
			}
		} else if pos.dir == "v" {
			for down := pos.x; down <= bottom; down++ {
				cur := matrix[down][pos.y]
				if cur == "." {
					matrix[down][pos.y] = "X"
					pos.x = down
				} else if cur == "#" {
					pos.dir = "<"
					pos.x = down - 1
					break
				}
			}
		}
	}

	res := 0
	for _, row := range matrix {
		fmt.Println(row)
		for _, col := range row {
			if col == "X" || col == "^" {
				res++
			}
		}
	}

	return res
}

func findStart(matrix [][]string) Pos {
	var pos Pos
	for i, line := range matrix {
		for j, cur := range line {
			if cur == "v" || cur == ">" || cur == "<" || cur == "^" {
				pos = Pos{x: i, y: j, dir: cur}
				matrix[i][j] = "+"
			}
		}
	}
	return pos
}
