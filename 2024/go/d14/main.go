package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed in.txt
var in string

func main() {
	// to high - 234703728
	robots := parseIn()
	// fmt.Println(part1(robots))
	fmt.Println(part2(robots))
}

type coord struct {
	x, y int
}

type robot struct {
	pos coord
	vol coord
}

func parseIn() []robot {
	lines := strings.Split(in, "\n")
	re := regexp.MustCompile("(-*[0-9]+)")
	var robots []robot
	for _, line := range lines {
		fs := re.FindAllString(line, -1)
		if len(fs) != 4 {
			break
		}
		px, _ := strconv.Atoi(fs[0])
		py, _ := strconv.Atoi(fs[1])
		vx, _ := strconv.Atoi(fs[2])
		vy, _ := strconv.Atoi(fs[3])
		robots = append(robots, robot{coord{px, py}, coord{vx, vy}})
	}
	return robots
}

var height int = 103
var width int = 101
var time int = 100

func part2(robots []robot) int {
	for i := 0; i <= height*width; i++ {
		for j, r := range robots {
			move(robots, r, j)
		}
		if checkMove(robots) {
			visual(robots)
			return i + 1
		}
	}
	return 0
}
func move(robots []robot, r robot, j int) {
	robots[j].pos.x = (r.pos.x + r.vol.x) % width
	robots[j].pos.y = (r.pos.y + r.vol.y) % height
	if robots[j].pos.x < 0 {
		robots[j].pos.x += width
	}
	if robots[j].pos.y < 0 {
		robots[j].pos.y += height
	}
}

func checkMove(robots []robot) bool {
	var grid [][]int
	grid = make([][]int, height)
	for y := 0; y < height; y++ {
		grid[y] = make([]int, width)
	}
	for _, robot := range robots {
		grid[robot.pos.y][robot.pos.x]++
	}
	count := 0
	for y := 0; y < height; y++ {
		count = 0
		for x := 0; x < width; x++ {
			if grid[y][x] == 1 {
				count++
			}
			if count > 10 {
				return true
			}
			if grid[y][x] == 0 {
				count = 0
			}
		}
	}
	return false
}
func visual(robots []robot) {
	var grid [][]int
	grid = make([][]int, height)
	for y := 0; y < height; y++ {
		grid[y] = make([]int, width)
	}
	for _, robot := range robots {
		grid[robot.pos.y][robot.pos.x]++
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == 1 {
				fmt.Printf("^")
				continue
			}
			fmt.Printf("_")
		}
		fmt.Println()
	}
}
func part1(robots []robot) int {
	dx, dy := width/2, height/2
	for i, r := range robots {
		robots[i].pos.x = (r.pos.x + r.vol.x*time + width*time) % width
		robots[i].pos.y = (r.pos.y + r.vol.y*time + height*time) % height
	}
	var quad [4]int
	for _, r := range robots {
		if r.pos.y < dy && r.pos.x < dx {
			quad[0]++
		}
		if r.pos.y > dy && r.pos.x < dx {
			quad[1]++
		}
		if r.pos.y < dy && r.pos.x > dx {
			quad[2]++
		}
		if r.pos.y > dy && r.pos.x > dx {
			quad[3]++
		}
	}
	res := 1
	for _, v := range quad {
		res *= v
	}
	return res
}
