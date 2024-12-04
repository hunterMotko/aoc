package main

import (
	"aoc/utils"
	"fmt"
	"log"
	"slices"
	"strings"
)

func part1(lines []string) {
	var left []int64
	var right []int64
	for _, v := range lines {
		arr := strings.Split(v, " ")
		first, last := arr[0], arr[len(arr)-1]
		x, err := utils.ParseInt(first)
		if err != nil {
			log.Fatalf("PARSE ERROR: %v\n", err)
		}
		y, er := utils.ParseInt(last)
		if er != nil {
			log.Fatalf("PARSE ERROR: %v\n", er)
		}
		left = append(left, x)
		right = append(right, y)
	}

	slices.Sort(left)
	slices.Sort(right)

	var res int64
	for i := 0; i < len(right); i++ {
		if right[i] > left[i] {
			res += right[i] - left[i]
			continue
		}
		res += left[i] - right[i]
	}
	fmt.Println(res)
}

func part2(lines []string) {
  var left []int64
  right := make(map[int64]int)

	for _, v := range lines {
		arr := strings.Split(v, " ")
		first, last := arr[0], arr[len(arr)-1]
		x, err := utils.ParseInt(first)
		if err != nil {
			log.Fatalf("PARSE ERROR: %v\n", err)
		}
    left = append(left, x)
		y, er := utils.ParseInt(last)
		if er != nil {
			log.Fatalf("PARSE ERROR: %v\n", er)
		}
    if _, ok := right[y]; !ok {
      right[y] = 1
      continue
    }
    right[y]++
	}
  var res int64
  for _, v := range left {
    res += v * int64(right[v])
  }
  fmt.Println(res)
}

func main() {
	lines, err := utils.ReadFile("./d1/in")
	if err != nil {
		log.Fatalf("READ ERROR: %v\n", err)
	}
  // part1(lines)
  part2(lines)
}
