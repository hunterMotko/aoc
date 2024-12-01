package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

func parse() [][]int {
	lines := strings.Split(input, "\n")
	var res [][]int
	for _, line := range lines {
		if line != "" {
			var temp []int
			arr := strings.Fields(line)
			for j := 0; j < len(arr); j++ {
				cur := toInt(arr[j])
				temp = append(temp, cur)
			}
			res = append(res, temp)
		}
	}

	return res
}

func toInt(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		fmt.Println(e)
	}
	return i
}

func getDiff(a, b int) int {
	return b - a
}

func pryamid(a []int) int {
	stack := [][]int{a}
	fmt.Println("stack", [][]int{a})
	check := true

	for check {
		var bucket []int
		work := stack[len(stack)-1]
		check = false
		for i := 1; i < len(work); i++ {
			diff := work[i] - work[i-1]
			bucket = append(bucket, diff)
			if diff != 0 {
				check = true
			}
		}
		if len(bucket) > 0 {
			stack = append(stack, bucket)
		}
	}

	var res int = 0
	for i := len(stack) - 1; i >= 0; i-- {
		res += stack[i][len(stack[i])-1]
	}
	return res
}

func Mirage(m [][]int) int {
	var res int
	for _, curArr := range m {
		res += pryamid(curArr)
	}
	return res
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i + 1, j - 1 {
		a[i], a[j] = a[j], a[i]
	}
}

func M2(m [][]int) int {
	var res int
	for _, curArr := range m {
    // For part two
    reverse(curArr)
		res += pryamid(curArr)
	}
	return res
}

func main() {
	matrix := parse()
	fmt.Println("RES", M2(matrix))
	fmt.Println("Mirage Maintenace")
}
