package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

// For example, suppose you were given the following strategy guide:

// A Y
// B X
// C Z
// This strategy guide predicts and recommends the following:

// In the first round, your opponent will choose
// Rock (A), and you should choose
// Paper (Y).
// This ends in a win for you with a score of
// 8 (2 because you chose Paper + 6 because you won).
// In the second round, your opponent will choose Paper (B), and you should choose Rock (X).
// This ends in a loss for you with a score of
// 1 (1 + 0).
// The third round is a draw with both players choosing Scissors, giving you a score of
// 3 + 3 = 6.
// In this example, if you were to follow the strategy guide, you would get a total score of 15 (8 + 1 + 6).

// PART 2
// The Elf finishes helping with the tent and sneaks back over to you.
// "Anyway, the second column says how the round needs to end: X means you need to lose,
// Y means you need to end the round in a draw, and Z means you need to win. Good luck!"
// The total score is still calculated in the same way,
// but now you need to figure out what shape to choose so the round ends as indicated.
// The example above now goes like this:
// Rock (A), draw (Y), 1 + 3 = 4.
// Paper (B), lose (X), 1 + 0 = 1.
// Scissors (C), win (Z), 1 + 6 = 7.
// 12.
func round(them string, need string) int {
	var score int

	switch need {
	case "X":
		if them == "A" {
			score = 3 + 0
		} else if them == "B" {
			score = 1 + 0
		} else if them == "C" {
			score = 2 + 0
		}
	case "Y":
		if them == "A" {
			score = 1 + 3
		} else if them == "B" {
			score = 2 + 3
		} else if them == "C" {
			score = 3 + 3
		}
	case "Z":
		if them == "A" {
			score = 2 + 6
		} else if them == "B" {
			score = 3 + 6
		} else if them == "C" {
			score = 1 + 6
		}
	}

	fmt.Printf("SCORE: %d - ", score)
	return score
}

func getInput() []string {
	fpath, err := os.Getwd()
	if err != nil {
		fmt.Printf("Getwd Error:: %v", err)
	}

	p := path.Join(fpath, "day2/d.txt")
	f, e := os.ReadFile(p)
	if e != nil {
		fmt.Printf("Read Error:: %v", e)
	}

	return strings.Split(string(f), "\n")
}

func game() int {
	input := getInput()
	var total int
	for i, r := range input {
		plays := strings.Split(r, " ")

		fmt.Printf("PLAYS: %v\n", plays)
		if len(plays) > 1 {
			total += round(plays[0], plays[1])
			fmt.Printf("ROUND - %d: %s - %s = %d \n", i, plays[0], plays[1], round(plays[0], plays[1]))
		}
	}
	return total
}

func main() {
	fmt.Println(game())
}
