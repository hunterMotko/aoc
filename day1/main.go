package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseTxtFile(filename string) string {
  b, e := os.ReadFile(filename)
  if e != nil {
    fmt.Printf("Read Error:: %v", e)
  }

  return string(b)
}

func sumInt(arr []int) int {
  s := 0
  for i := 0; i < len(arr); i++ {
    s += arr[i]
  }
  return s
}
func sumStr(arr []string) int {
  s := 0
    for _, y := range arr {
      if y == "" {
        continue
      }
      a, err := strconv.Atoi(y)
      if err != nil {
        fmt.Println(err)
      }
      s += a
    }
    return s
}

func findElf(str string) int {
  some := strings.Split(str, "\n\n")
  var calories []int
  for _, v := range some {
    w := strings.Split(v, "\n")
    temp := sumStr(w)
    calories = append(calories, temp)
  }
  sort.Ints(calories)
  n := len(calories)
  asd := calories[n-3:]
  fmt.Println(asd)
  return sumInt(asd)
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("%v", err)
	} else {
		fmt.Println(dir)
	}

	filePath := dir + "/day1/calories.txt"
  s := parseTxtFile(filePath)
  i := findElf(s)
  fmt.Printf("elf?--%d", i)
}
