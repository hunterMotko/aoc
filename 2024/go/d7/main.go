package main

import (
	"aoc/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	lines, err := utils.ReadFile("./d7/in.txt")
	if err != nil {
		log.Fatalf("READ ERR: %v\n", err)
	}
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(lines []string) int {
	var res int
	for _, v := range lines {
		line := strings.Split(v, ": ")
		cal, _ := strconv.Atoi(line[0])
		nums := strings.Split(line[1], " ")
		var vals []int
		for _, v := range nums {
			n, _ := strconv.Atoi(v)
			vals = append(vals, n)
		}
		if generate(cal, vals, vals[0], 1) {
			res += cal
		}
	}
	return res
}

func generate(cal int, vals []int, res int, i int) bool {
	if cal == res {
		return true
	}
	if i == len(vals) {
		return cal == res
	}
	if generate(cal, vals, res+vals[i], i+1) {
		return true
	}
	if res != 0 && generate(cal, vals, res*vals[i], i+1) {
		return true
	}
	return false
}

func part2(lines []string) int64 {
	var res int64
	for _, v := range lines {
		line := strings.Split(v, ": ")
		cal, _ := strconv.Atoi(line[0])
		nums := strings.Split(line[1], " ")
		var vals []int
		for _, v := range nums {
			n, _ := strconv.Atoi(v)
			vals = append(vals, n)
		}
		if dfs(cal, 0, 0, vals, "+") {
			res += int64(cal)
		}
	}
	return res
}

func dfs(cal int, res int, i int, vals []int, op string) bool {
	if i >= len(vals) {
		return cal == res
	}
	cur := vals[i]
	if op == "+" {
		res += cur
	} else if op == "*" {
		res *= cur
	} else {
		n, _ := strconv.Atoi(fmt.Sprintf("%d%d", res, cur))
		res = n
	}
	ops := []string{"+", "*", "||"}
	for _, o := range ops {
		if dfs(cal, res, i+1, vals, o) {
			return true
		}
	}
	return false
}
