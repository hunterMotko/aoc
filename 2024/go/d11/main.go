package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed in.txt
var in string

func main() {
	arr := strings.Split(strings.ReplaceAll(in, "\n", ""), " ")
	// fmt.Println(part1(arr))
	fmt.Println(part2(arr))
}

// If the stone is engraved with the number 0, it is replaced by a stone
// engraved with the number 1.
// If the stone is engraved with a number that has an even number of digits, it
// is replaced by two stones. The left half of the digits are engraved on the
// new left stone, and the right half of the digits are engraved on the new
// right stone. (The new numbers don't keep extra leading zeroes: 1000 would
// become stones 10 and 0.)
// If none of the other rules apply, the stone is replaced by a new stone; the
// old stone's number multiplied by 2024 is engraved on the new stone.

func part1(arr []string) int {
	for i := 0; i < 75; i++ {
		var temp []string
		for _, v := range arr {
			if v == "0" {
				temp = append(temp, "1")
			} else if len(v)%2 == 0 {
				l := len(v) / 2
				a, b := v[:l], v[l:]
				n, _ := strconv.Atoi(b)
				b = fmt.Sprintf("%d", n)
				temp = append(temp, a)
				temp = append(temp, b)
			} else {
				n, _ := strconv.Atoi(v)
				temp = append(temp, fmt.Sprintf("%d", n*2024))
			}
		}
		arr = temp
	}
	return len(arr)
}

type item struct {
	v, n int
}

var cache = make(map[item]int)

func part2(arr []string) int {
	res := 0
	n := 75
	for _, v := range arr {
		v, _ := strconv.Atoi(v)
		res += walk(v, n)
	}
	return res
}

func walk(v, n int) int {
	if n == 0 {
		return 1
	}
	cur := item{v, n}
	if r, ok := cache[cur]; ok {
		return r
	}
	if v == 0 {
		res := walk(1, n-1)
		cache[cur] = res
		return res
	}
	s := fmt.Sprintf("%d", v)
	if len(s)%2 == 0 {
		a, _ := strconv.Atoi(s[:len(s)/2])
		b, _ := strconv.Atoi(s[len(s)/2:])
		res := walk(a, n-1) + walk(b, n-1)
		cache[cur] = res
		return res
	}
	res := walk(v*2024, n-1)
	cache[cur] = res
	return res
}
