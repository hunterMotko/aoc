package main

import (
	"aoc/utils"
	"fmt"
	"log"
)

func main() {
	lines, err := utils.ReadFile("./d8/in.txt")
	if err != nil {
		log.Fatalf("READ ERR: %v\n", err)
	}
	fmt.Println(part1(lines))
}

func part1(lines []string) int {
	locs := findLoc(lines)
	return findAnt(locs, lines)
}

type loc struct {
	x int
	y int
}

func (l loc) isValid(aw, ah int) bool {
	return l.x >= 0 && l.x < aw && l.y >= 0 && l.y < ah
}

func findLoc(lines []string) map[rune][]loc {
	locs := make(map[rune][]loc)
	for y, row := range lines {
		for x, ant := range row {
			if ant != '.' {
				locs[ant] = append(locs[ant], loc{x: x, y: y})
			}
		}
	}
	return locs
}

func findAnt(locs map[rune][]loc, ants []string) int {
	uniq := make(map[loc]bool)
	for _, lc := range locs {
		l := len(lc)
		for i := 0; i < l-1; i++ {
			for j := i + 1; j < l; j++ {
				dx := lc[j].x - lc[i].x
				dy := lc[j].y - lc[i].y
        // part 2 need to check for out of bounds in case the anntennas line up
				ob := 0
				for k := 0; ob < 2; k++ {
					ob = 0
					a1 := loc{lc[i].x - k*dx, lc[i].y - k*dy}
					if a1.isValid(len(ants[0]), len(ants)) {
						uniq[a1] = true
					} else {
						ob++
					}
					a2 := loc{lc[j].x + k*dx, lc[j].y + k*dy}
					// fmt.Println(lc[j], lc[i], dx, dy, a1, a2)
					if a2.isValid(len(ants[0]), len(ants)) {
						uniq[a2] = true
					} else {
						ob++
					}
				}
			}
		}
	}
	return len(uniq)
}
