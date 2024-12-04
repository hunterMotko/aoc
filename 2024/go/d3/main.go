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

func part1(doc string) int {
  re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
  matches := re.FindAllStringSubmatch(doc, len(doc)+1)
  sum := 0
  for _, match := range matches {
    x, _ := strconv.Atoi(match[1])
    y, _ := strconv.Atoi(match[2])
    sum += (x * y)
  }
  return sum
}

func removeDont(doc string) string {
  var filter []string
  dos := strings.Split(doc, "do()")
  for _, v := range dos {
    p, _, _ := strings.Cut(v, "don't()")
    filter = append(filter, p)
  }
  return strings.Join(filter, "")
}
func part2(doc string) int {
  return part1(removeDont(doc))
}

func main() {
	// HIGH
	// 190 433 914
	// CORRECT 
  // 184 511 516
  // LOW
	// 183 583 220
	// 183 583 196
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}
