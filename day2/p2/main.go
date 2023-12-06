package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/huntermotko/aoc/utils"
)

type game struct {
	id    int
	red   int
	blue  int
	green int
}

func newGame(id int) game {
	return game{
		id:    id,
		red:   0,
		blue:  0,
		green: 0,
	}
}

func CheckAtoi(e error) {
	if e != nil {
		fmt.Printf("ATOI ERR: %v", e)
	}
}

func Game(d []string) int {
	var total int
	for i, v := range d {
		gameRounds := strings.Split(v, ":")
		if len(gameRounds) > 1 {
			r := gameRounds[1]
			g := newGame(i + 1)
			rounds := strings.Split(r, ";")
			for _, v := range rounds {
				match := strings.Split(v, ",")
				g.Match(match)
			}
      fmt.Println(g.red * g.blue * g.green)
			total += g.red * g.blue * g.green
		}
	}
	return total
}

func (g *game) Match(m []string) {
	for _, v := range m {
		x := strings.Split(v, " ")
		cubes, color := x[1], x[2]
		i, e := strconv.Atoi(cubes)
		CheckAtoi(e)
		switch {
		case color == "red" && i > g.red:
			g.red = i
		case color == "blue" && i > g.blue:
			g.blue = i
		case color == "green" && i > g.green:
			g.green = i
		}
	}
}

func main() {
	str := utils.GetInput("day2")
	arr := strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n")
  // 	arr := []string{
  // 		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
  // 		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
  // 		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
  // 		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
  // 		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
  // 	}
	total := Game(arr)
	fmt.Printf("RESULT: %d", total)
}
