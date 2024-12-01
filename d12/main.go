package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// . operational
// # damaged
// ? unknown
func parse(file string) []string {
	f, err := os.Open("./d12/" + file)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	b := bufio.NewScanner(f)
	var rows []string
	for b.Scan() {
		rows = append(rows, b.Text())
	}
	return rows
}

func countDamaged(s string, keys []int) int {
	if len(keys) == 0 {
		return 1
	}
	var total int
	j := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			j = 0
			continue
		}
		if s[i] == '?' && (i == len(s)-1 || s[i+1] == '.' || s[i+1] == '?' || s[i+1] == '#') {
			j++
		}
		if s[i] == '#' {
			j++
		}
		if j == keys[0] {
			if i < len(s)-1 && s[i+1] == '#' {
				j--
				continue
			}
			k := 0
			if i < len(s)-1 && s[i+1] == '?' {
				k = 1
			}
			total += countDamaged(s[i+k+1:], keys[1:])
			i -= j - 1
			for s[i] == '#' {
				if i < len(s)-1 &&  s[i+1] == '#' || s[i+1] == '?' && s[i] == '#' {
					i++
				} else {
					break
				}
			}
			j = 0
		}
	}
	return total
}

func main() {
	rows := parse("ex_one")
	for _, cols := range rows {
		col := strings.Split(cols, " ")
		ints := asInts(strings.Split(col[1], ","))
		count := countDamaged(col[0], ints)
		fmt.Println(count)
	}
	fmt.Println("Hot Springs")
}

func asInts(s []string) []int {
	var ints []int
	for _, str := range s {
		i, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}
	return ints
}
