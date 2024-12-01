package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/huntermotko/aoc/utils"
)

type game struct {
	id    int
  good bool
}

const (
	MaxRed   = 12
	MaxGreen = 13
	MaxBlue  = 14
)

func newGame(id int) game {
	return game{
		id:    id,
    good: true, 
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
      if g.good {
        total += g.id
      }
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
		case color == "red" && i > MaxRed:
      g.good = false
		case color == "blue" && i > MaxBlue:
      g.good = false
		case color == "green" && i > MaxGreen:
      g.good = false
		}
	}
}

func main() {
  str := utils.GetInput("day2")
  arr := strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n")
	total := Game(arr)
	fmt.Printf("RESULT: %d", total)
}
