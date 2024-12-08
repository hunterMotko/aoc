package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed in.txt
var in string

func formatIn() ([][]string, [][]string) {
	arr := strings.Split(in, "\n\n")
	rules := strings.Split(arr[0], "\n")
	var r [][]string
	for _, v := range rules {
		r = append(r, strings.Split(v, "|"))
	}
	updates := strings.Split(arr[1], "\n")
	var u [][]string
	for _, v := range updates {
		u = append(u, strings.Split(v, ","))
	}
	return r, u
}

func part1(rules [][]string, updates [][]string) int {
	rMap := make(map[string][]string)
	for _, rule := range rules {
		key := rule[0]
		if _, ok := rMap[key]; !ok {
			rMap[key] = []string{rule[1]}
			continue
		}
		rMap[key] = append(rMap[key], rule[1])
	}

	res := 0
	for _, update := range updates[:len(updates)-1] {
		check := true
		for i, v := range update {
			if i < len(update)-1 {
				curList := rMap[v]
				next := update[i+1]
				if !slices.Contains(curList, next) {
					check = false
					break
				}
			}
		}

		if check {
			n, _ := strconv.Atoi(update[len(update)/2])
			res += n
		}
	}

	return res
}
func part2(rules [][]string, updates [][]string) int {
	rMap := make(map[string][]string)
	for _, rule := range rules {
		key := rule[0]
		if _, ok := rMap[key]; !ok {
			rMap[key] = []string{rule[1]}
			continue
		}
		rMap[key] = append(rMap[key], rule[1])
	}

	res := 0
	for _, update := range updates[:len(updates)-1] {
		check := true
		for i, v := range update {
			if i < len(update)-1 {
				curList := rMap[v]
				next := update[i+1]
				if !slices.Contains(curList, next) {
					check = false
					break
				}
			}
		}

		if !check {
      // Reorder
      for i := 0; i < len(update); i++ {
        for j := i+1; j < len(update); j++ {
          cur := update[j]
          list := rMap[update[i]]
          if slices.Contains(list, cur) {
            update[j] = update[i]
            update[i] = cur
          }
        }
      }
      fmt.Println(update)
			n, _ := strconv.Atoi(update[len(update)/2])
			res += n
		}
	}

	return res
}

func main() {
	r, u := formatIn()
	fmt.Println(part2(r, u))
}
