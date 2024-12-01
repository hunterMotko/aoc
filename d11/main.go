package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
)

func main() {
	matrix := parse("input")
  fmt.Printf("\n------\nP1 Res: %d\n------\n", part_one(matrix))
  fmt.Printf("\n------\nP2 Res: %d\n------\n", part_two(matrix))
}

func part_one(matrix []string) int {
	h, v, hash := expand(matrix)
  res := find(h, v, hash, 2)
  return res
}
func part_two(matrix []string) int {
	h, v, hash := expand(matrix)
  res := find(h, v, hash, 1000000)
  return res
}

func parse(filename string) []string {
	f, err := os.Open("./d11/" + filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scn := bufio.NewScanner(f)
	var matrix []string
	for scn.Scan() {
		line := scn.Text()
		fmt.Println(line)
		matrix = append(matrix, line)
	}

	if err := scn.Err(); err != nil {
		log.Fatal(err)
	}
	return matrix
}

type Pos struct {
	x int
	y int
}

// takes a matrix []string representation of the galaxy layout
// returns spaced columns map, spaced rows map, and hashtags positions
func expand(matrix []string) (map[int]interface{}, map[int]interface{}, []Pos) {
	hashtag := regexp.MustCompile("#")
	vertical := map[int]interface{}{}
	horizontal := map[int]interface{}{}
	hashes := []Pos{}

	for i := 0; i < len(matrix[0]); i++ {
		matches := hashtag.FindAllStringIndex(matrix[i], -1)
		if len(matches) == 0 {
			horizontal[i] = struct{}{}
		}
		for _, m := range matches {
			hashes = append(hashes, Pos{
				x: i,
				y: m[0],
			})
		}
		countY := 0
		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] == '#' {
				countY++
			}
		}
		if countY == 0 {
			vertical[i] = struct{}{}
		}
	}
  // returns spaced columns map, spaced rows map, and hashtags positions
	return horizontal, vertical, hashes
}


func find(h, v map[int]interface{}, hash []Pos, times int) int {
  mod := []Pos{}
  for _, ch := range hash {
    a, b := ch.x, ch.y
    for x := range h {
      if x < ch.x {
        a += times - 1
      }
    }
    for y := range v {
      if y < ch.y {
        b += times - 1
      }
    }
    mod = append(mod, Pos{a, b})
  }
  c := 0
  for _, m := range mod {
    for _, h := range mod {
      if m == h {
        continue
      }
      x := int(math.Abs(float64(h.x - m.x)))
      y := int(math.Abs(float64(h.y - m.y)))
      c += x + y
    }
  }
  return int(c / 2)
}

