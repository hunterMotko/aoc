package main

import (
	"fmt"
  _ "embed"
	"strings"
)

//go:embed ex_three
var input string

func main() {
  fmt.Println(strings.Split(input, "\n"))
  RunDay10Part1()
  RunDay10Part2()
}
