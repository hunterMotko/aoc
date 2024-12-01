package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

type Leaf struct {
	Key   string
	Left  *Leaf
	Right *Leaf
}
func Haunted(m map[string][]string, d string) int {
	var res int
	here := "AAA"
	for here != "ZZZ" {
		for _, cur := range d {
			res++
			if cur == 'R' {
				here = m[here][1]
			} else {
				here = m[here][0]
			}
		}
	}
	return res
}


func Traverse(bra *Leaf, dirs string, isFinish func(k string) bool) int {
	cur := bra
	curI := 0
	cnt := 0
	for cur != nil && !check(cur.Key) {
		cnt++
		if dirs[curI] == 'L' {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
		curI = (curI + 1) % len(dirs)
	}
	return cnt
}
func check(k string) bool { 
  return k[len(k)-1] == 'Z' 
}
func H2() int {
	in := strings.Split(input, "\n\n")
	re := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)
	list := strings.Split(in[1], "\n")
	branch := make(map[string]*Leaf, 0)
	branches := make([]*Leaf, 0)
	for _, v := range list {
		if v == "" {
			continue
		}
		mt := re.FindStringSubmatch(v)
		k, l, r := mt[1], mt[2], mt[3]
		left := upsert(branch, l, nil, nil)
		right := upsert(branch, r, nil, nil)
		br := upsert(branch, k, left, right)
		if k[2] == 'A' {
			branches = append(branches, br)
		}
	}

	res := 1
	for _, bra := range branches {
    fmt.Println(bra)
		res = Lcm(res, Traverse(bra, in[0], LastCheck))
	}

	return res
}

func upsert(nodes map[string]*Leaf, key string, left, right *Leaf) *Leaf {
	if nodes[key] == nil {
		nodes[key] = &Leaf{
			Key:   key,
			Left:  left,
			Right: right,
		}
	} else {
		if left != nil {
			nodes[key].Left = left
		}
		if right != nil {
			nodes[key].Right = right
		}
	}
	return nodes[key]
}

func Lcm(a, b int) int {
	return a * b / Gcd(a, b)
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	// 	n := make(map[string][]string)
	// 	d := parse(n)
	// 	f := Haunted(n, d)
	// 	fmt.Println(f)
	// in := strings.TrimRight(input, "\n")
	// r := part2(in)
  res := H2()
  fmt.Println(res)
}
