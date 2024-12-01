package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/huntermotko/aoc/utils"
)

func main() {
  input := utils.GetInput("day2")
	re := regexp.MustCompile(`(\d+) (\w+)`)

	part1, part2 := 0, 0
	for i, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		min := map[string]int{}

		for _, m := range re.FindAllStringSubmatch(s, -1) {
			n, _ := strconv.Atoi(m[1])
			min[m[2]] = slices.Max([]int{min[m[2]], n})
		}

		if min["red"] <= 12 && min["green"] <= 13 && min["blue"] <= 14 {
			part1 += i + 1
		}
		part2 += min["red"] * min["green"] * min["blue"]
	}
	fmt.Println(part1)
	fmt.Println(part2)
}
