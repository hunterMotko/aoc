package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Round struct {
	Hand string
	Bid  int
}

type Rounds []Round

// have to change this function to acount for jokers
//p1
// func greatestChar(s string) int {
// 	freq := make(map[rune]int)
// 	for _, v := range s { freq[v]++ }
// 	list := make([]int, 0, len(freq))
// 	for _, v := range freq { list = append(list, v) }
// 	var res int
// 	if len(list) == 1 {
// 		res = 7
// 	} else if len(list) == 2 && list[0] == 4 || list[1] == 4 {
// 		res = 6
// 	} else if len(list) == 2 && (list[0] == 3 && list[1] == 2) || (list[0] == 2 && list[1] == 3) {
// 		res = 5
// 	} else if len(list) == 3 && list[0] == 3 || list[1] == 3 || list[2] == 3 {
// 		res = 4
// 	} else if len(list) == 3 && (list[0] == 2 && list[1] == 2) ||
// 		(list[0] == 2 && list[2] == 2) || (list[1] == 2 && list[2] == 2) {
// 		res = 3
// 	} else if len(list) == 4 {
// 		res = 2
// 	} else if len(list) == 5 {
// 		res = 1
// 	}
// 	return res
// }

func Greatest(s string) int {
	freq := make(map[rune]int)
	great := 0
	jokers := 0
	for _, v := range s {
		if v == 74 {
			jokers++
		} else {
			freq[v]++
			if freq[v] > great {
				great = freq[v]
			}
		}
	}
	return great + jokers
}

var ranks = func() map[string]int {
	return map[string]int{
		"A": 13, "K": 12, "Q": 11, "T": 10, "9": 9, "8": 8,
		"7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2, "J": 1,
	}
}

func getRank(k string) int {
	return ranks()[k]
}

func (r Rounds) Len() int { return len(r) }
func (r Rounds) Less(i, j int) bool {
	cur := Greatest(r[i].Hand)
	next := Greatest(r[j].Hand)
	if cur == next {
		for k := 0; k < r.Len(); k++ {
			c := r[i].Hand[k]
			n := r[j].Hand[k]
      fmt.Println(c, cur, n, next)
			if c == n {
				continue
			}
			return getRank(string(c)) < getRank(string(n))
		}
	}
	return cur < next
}

func (r Rounds) Swap(i, j int) { r[i], r[j] = r[j], r[i] }

func CamelCards(rs Rounds) int {
	var res int
	for i, v := range rs {
		res += (i + 1) * v.Bid
    fmt.Println(res, i, v)
	}
	return res
}

func main() {
	rounds := strings.Split(input, "\n")
	var rs []Round

	for _, v := range rounds {
		if v != "" {
			a := strings.Split(v, " ")
			bid, err := strconv.Atoi(a[1])
			if err != nil {
				fmt.Println("ERR", err)
			}
			rs = append(rs, Round{
				Hand: a[0],
				Bid:  bid,
			})
		}
	}

	sort.Sort(Rounds(rs))
	// --- right answer ---
  //     253499763
}
