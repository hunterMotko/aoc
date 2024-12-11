package main

import (
	"aoc/utils"
	"fmt"
	"log"
)

func main() {
	matrix, err := utils.ReadFileToIntMatrix("./d10/in.txt")
	if err != nil {
		log.Fatalf("READ ERR: %v\n", err)
	}
	fmt.Println(part1(matrix))
}

type pos struct {
	y, x int
}

var th = make(map[pos][]pos)

func part2(matrix [][]int) int {
	res := 0
	for y, row := range matrix {
		for x, c := range row {
			if c == 0 {
				ps := pos{y, x}
				th[ps] = make([]pos, 0)
				traverse(matrix, x, y, 0, ps)
			}
		}
	}
	for _, v := range th {
		res += len(v)
	}
	return res
}
func part1(matrix [][]int) int {
	res := 0
	for y, row := range matrix {
		for x, c := range row {
			if c == 0 {
				ps := pos{y, x}
				th[ps] = make([]pos, 0)
				traverse(matrix, x, y, 0, ps)
			}
		}
	}
	for _, v := range th {
		res += len(v)
	}
	return res
}
func traverse(m [][]int, x, y, nx int, start pos) {
	if y < 0 || y > len(m)-1 || x < 0 || x > len(m[0])-1 {
		return
	}
	cur := m[y][x]
	if cur != nx {
		return
	}
	if cur == 9 && nx == 9 {
		th[start] = append(th[start], pos{y,x})
		// ps := th[start]
		// if !slices.Contains(ps, pos{y, x}) {
		// 	th[start] = append(th[start], pos{y, x})
		// }
		return
	}
	traverse(m, x, y-1, nx+1, start)
	traverse(m, x, y+1, nx+1, start)
	traverse(m, x-1, y, nx+1, start)
	traverse(m, x+1, y, nx+1, start)
}
