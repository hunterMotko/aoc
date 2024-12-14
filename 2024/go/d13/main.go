package main

import (
	"aoc/utils"
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed in.txt
var in string

type pos struct {
	x, y int64
}
type claw struct {
	btnA  pos
	btnB  pos
	prize pos
}

func main() {
	claws := parseBlocks()
	fmt.Println(part1(claws))
}

var add int64 = 10000000000000

func parseBlocks() []claw {
	blocks := strings.Split(in, "\n\n")
	var claws []claw
	re := regexp.MustCompile("([0-9]+)")
	for _, block := range blocks {
		rows := strings.Split(block, "\n")
		cl := claw{}
		for i, row := range rows {
			m := re.FindAllString(row, -1)
			switch i {
			case 0:
				x, _ := utils.ParseInt(m[0])
				y, _ := utils.ParseInt(m[1])
				cl.btnA = pos{x, y}
			case 1:
				x, _ := utils.ParseInt(m[0])
				y, _ := utils.ParseInt(m[1])
				cl.btnB = pos{x, y}
			case 2:
				x, _ := utils.ParseInt(m[0])
				y, _ := utils.ParseInt(m[1])
				cl.prize = pos{add + x, add + y}
			}
		}
		claws = append(claws, cl)
	}
	return claws
}

func part1(claws []claw) int {
	var res int
	for _, claw := range claws {
		res += int(calc2(claw))
	}
	return res
}

func calc2(cl claw) int64 {
  // x = bx * -cy - -cx * by
	x := cl.btnB.x*(-cl.prize.y) - (-cl.prize.x)*cl.btnB.y
  // y = -cx * ay - ax * -cy
	y := (-cl.prize.x)*cl.btnA.y - cl.btnA.x*(-cl.prize.y)
  // z = ax * by - bx * ay
	z := cl.btnA.x*cl.btnB.y - cl.btnB.x*cl.btnA.y
	if x%z != 0 || y%z != 0 {
		return 0
	}
	// x = x / z
	x /= z
	// y = y / z
	y /= z
	// Only add to the total if x and y are non-negative
	if x >= 0 && y >= 0 {
		return x*3 + y
	}
	return 0
}

func calculateMoves(cl claw) int {
	a, b := 0, 0
	for n := 0; n <= int(cl.prize.x/cl.btnA.x); n++ {
		for m := 0; m <= int(cl.prize.y/cl.btnB.y); m++ {
			if (int64(n)*cl.btnA.x+int64(m)*cl.btnB.x == cl.prize.x) &&
				(int64(n)*cl.btnA.y+int64(m)*cl.btnB.y == cl.prize.y) {
				a = n
				b = m
			}
		}
	}
	return a*3 + b
}
