package main

import (
	"fmt"
	"strings"

	"github.com/huntermotko/aoc/utils"
)

func main() {
  input := utils.GetInput("day1")

  calc := func(r *strings.Replacer) (result int) {
		for _, s := range strings.Fields(input) {
			s = r.Replace(r.Replace(s))
			result += 10 * int(s[strings.IndexAny(s, "123456789")]-'0')
			result += int(s[strings.LastIndexAny(s, "123456789")] - '0')
		}
    return
	}

	fmt.Println(calc(strings.NewReplacer()))
	fmt.Println(calc(strings.NewReplacer(
    "one", "o1e", "two", "t2o", "three", 
    "t3e", "four", "f4r", "five", "f5e", 
    "six", "s6x", "seven", "s7n", "eight", 
    "e8t", "nine", "n9e",
  )))
}
