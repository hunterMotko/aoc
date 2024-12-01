package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Time:      7  15   30
// Distance:  9  40  200

// race lasts 7 milliseconds. The record distance in this race is 9 millimeters.
// race lasts 15 milliseconds. The record distance in this race is 40 millimeters.
// race lasts 30 milliseconds. The record distance in this race is 200 millimeters.

// Your toy boat has a starting speed of zero millimeters per millisecond.
// each whole millisecond you spend at the beginning of the race holding down the button,
// the boat's speed increases by one millimeter per millisecond.
// hold down ms
// boatspeed mm per ms
// first race lasts 7 milliseconds,
// t hold the button at all (that is, hold it for 0 milliseconds) at the start of the race.
// The boat won't move; it will have traveled 0 millimeters by the end of the race.

// Hold the button for 1 millisecond at the start of the race.
// boat will travel at a speed of 1 millimeter per millisecond for 6 milliseconds,
// reaching a total distance traveled of 6 millimeters.

// Hold the button for 2 milliseconds,
// giving the boat a speed of 2 millimeters per millisecond.
// It will then get 5 milliseconds to move, reaching a total distance of 10 millimeters.

// Hold the button for 3 milliseconds.
// After its remaining 4 milliseconds of travel time,
// the boat will have gone 12 millimeters.

// Hold the button for 4 milliseconds.
// After its remaining 3 milliseconds of travel time,
// the boat will have gone 12 millimeters.

// Hold the button for 5 milliseconds,
// causing the boat to travel a total of 10 millimeters.
// Hold the button for 6 milliseconds,
// causing the boat to travel a total of 6 millimeters.
// Hold the button for 7 milliseconds.
// That's the entire duration of the race. You never let go of the button.
// The boat can't move until you let go of the button.
// Please make sure you let go of the button so the boat gets to move. 0 millimeters.
// Since the current record for this race is 9 millimeters,
// there are actually 4 different ways you could win:
// you could hold the button for 2, 3, 4, or 5 milliseconds at the start of the race.

// In the second race, you could hold the button for at least 4 milliseconds and at most
// 11 milliseconds and beat the record, a total of 8 different ways to win.
// In the third race, you could hold the button for at least 11 milliseconds and no more
// than 19 milliseconds and still beat the record, a total of 9 ways you could win.

// you get 288 (4 * 8 * 9).

func ReaderToStrings(input io.Reader) []string {
	scanner := bufio.NewScanner(input)
	var res []string
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}

func main() {
	f, err := os.Open("./wait/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	in := ReaderToStrings(f)
	fmt.Println("Wait for it", wait(in))
}
func wait(a []string) int {
  var x []int

	for _, v := range a {
		arr := strings.Fields(strings.Split(v, ":")[1])
    str := strings.Join(arr, "")
    i, e := strconv.Atoi(str)
    if e != nil {
      fmt.Println("ERROR: ", str)
    }
    x = append(x, i)
	}
  fmt.Println(x)

  return crazyTimeConv(x[0], x[1])
}

// func waitForIt(a []string) int {
//   res := 1
// 	var matrix [2][4]int
//
// 	for i, v := range a {
// 		ints := strings.Fields(strings.Split(v, ":")[1])
// 		if len(ints) > 0 {
// 			for j, x := range ints {
// 				z, err := strconv.Atoi(x)
// 				if err != nil {
// 					fmt.Println("ERROR: ", err)
// 				}
// 				matrix[i][j] = z
// 			}
// 		}
// 	}
//
// 	for i := 0; i < len(matrix[0]); i++ {
// 		time := matrix[0][i]
// 		dist := matrix[1][i]
// 		res *= crazyTimeConv(time, dist)
// 	}
//
// 	return res
// }

func crazyTimeConv(time, dist int) int {
	var res int
	remaining := time
	for i := 0; i < time; i++ {
		eq := i * remaining
		if eq > dist {
			res++
		}
		remaining--
	}
	return res
}
