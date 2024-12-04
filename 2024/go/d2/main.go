package main

import (
	"aoc/utils"
	"fmt"
	"log"
)

func isSafe(arr []int) bool {
	asc := true
	for i := 0; i < len(arr)-1; i++ {
		l, r := arr[i], arr[i+1]
		if i == 0 && l > r {
			asc = false
		}
		if asc {
			dif := r - l
			if l > r || dif < 1 || dif > 3 {
				return false
			}
		} else {
			dif := l - r
			if l < r || dif < 1 || dif > 3 {
				return false
			}
		}
	}
  return true
}

func part1(levels []string) int {
	var res int
	for _, level := range levels {
		arr := utils.StrToint(level, " ")
		if isSafe(arr) {
			res++
		}
	}
	return res
}

func part2(levels []string) int {
	safe, rmSafe := 0, 0
	for _, level := range levels {
		arr := utils.StrToint(level, " ")
		if isSafe(arr) {
			safe++
		} else if checkWithDeletion(arr) {
			rmSafe++
		}
	}
	fmt.Println(safe)
	return safe + rmSafe
}

func checkWithDeletion(report []int) bool {
	for i := 0; i < len(report); i++ {
		if isSafeWithDeletion(report, i) {
			return true
		}
	}
	return false
}

func isSafeWithDeletion(report []int, del int) bool {
	cr := make([]int, len(report))
	copy(cr, report)
	if del == len(cr)-1 {
		cr = cr[:del]
	} else {
		cr = append(cr[:del], cr[del+1:]...)
	}
	return isSafe(cr)
}

func main() {
	levels, err := utils.ReadFile("./d2/in")
	if err != nil {
		log.Fatalf("READ ERROR: %v\n", err)
	}
	res := part1(levels)
	fmt.Println("RES", res)
	// res = 442
	res = part2(levels)
	fmt.Println("RES", res)
}
