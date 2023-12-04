package main

import (
  "os"
	"fmt"
	"math"
	"path"
  "strings"
)

func findUniq(str string) string {
	l := int(math.Round(float64(len(str)/2)))
  lc := make(map[rune]bool)
	one, two := str[:l], str[l:]
  for _, k := range one {
    lc[k] = true
  }
  var arr string
  for _, k := range two {
    if lc[k] {
      fmt.Printf("RUNE: %c, ", k)
      arr = string(k)
    }
  }
  return arr
}


func getInput() []string {
	fpath, err := os.Getwd()
	if err != nil {
		fmt.Printf("Getwd Error:: %v", err)
	}
	p := path.Join(fpath, "day3/d.txt")
	f, e := os.ReadFile(p)
	if e != nil {
		fmt.Printf("Read Error:: %v", e)
	}
	return strings.Split(string(f), "\n")
}

func tally(data []string) int {
	k := map[string]int{
		"a": 1, "b": 2, "c": 3,
		"d": 4, "e": 5, "f": 6,
		"g": 7, "h": 8, "i": 9,
		"j": 10, "k": 11, "l": 12,
		"m": 13, "n": 14, "o": 15,
		"p": 16, "q": 17, "r": 18,
		"s": 19, "t": 20, "u": 21,
		"v": 22, "w": 23, "x": 24,
		"y": 25, "z": 26, "A": 27,
		"B": 28, "C": 29, "D": 30,
		"E": 31, "F": 32, "G": 33,
		"H": 34, "I": 35, "J": 36,
		"K": 37, "L": 38, "M": 39,
		"N": 40, "O": 41, "P": 42,
		"Q": 43, "R": 44, "S": 45,
		"T": 46, "U": 47, "V": 48,
		"W": 49, "X": 50, "Y": 51, "Z": 52,
	}
  var arr []int
  for _, v := range data {
    arr = append(arr, k[findUniq(v)])
  }

  var res int
  for _, i := range arr {
    res += i
  }

  return res
}

func main() {
  in := getInput()
  out := tally(in)
  fmt.Printf("FINAL TOTAL: %d", out)
}
