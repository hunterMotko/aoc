package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed test1.txt
var in string

func main() {
	grid, moves := parse()
	fmt.Println(part1(grid, moves))
}

type coord struct {
	x, y int
}

type warehouse struct {
	height, width int
	robot         coord
	grid          [][]string
}

func (wh warehouse) moveInBounds(x, y int) bool {
	return wh.robot.x+x >= 0 && wh.robot.x+x <= wh.width-1 && wh.robot.y+y >= 0 && wh.robot.y+y <= wh.height-1
}

func (wh warehouse) move(x, y int) {
	cur := coord{wh.robot.x + x, wh.robot.y + y}
}

func (wh warehouse) countBoxPositions() {
}

func parse() (warehouse, []string) {
	var grid [][]string
	arr := strings.Split(in, "\n\n")
	var robPos coord
	for y, line := range strings.Split(arr[0], "\n") {
		var cur []string
		for x, v := range strings.Split(line, "") {
			if v == "@" {
				robPos = coord{x, y}
			}
			cur = append(cur, v)
		}
		grid = append(grid, cur)
	}

	var moves []string
	for _, v := range strings.Split(arr[1], "") {
		if v != "\n" {
			moves = append(moves, v)
		}
	}

	return warehouse{
		height: len(grid),
		width:  len(grid[0]),
		robot:  robPos,
		grid:   grid,
	}, moves
}

func part1(wh warehouse, moves []string) int {
	for _, move := range moves {
		switch move {
		case "^":
			wh.move(0, -1)
		case "v":
			wh.move(0, 1)
		case ">":
			wh.move(1, 0)
		case "<":
			wh.move(-1, 0)
		}
	}
	return 0
}
