package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
)

// So, if card 10 were to have 5 matching numbers, you would win one copy each of cards 11, 12, 13, 14, and 15.
// Copies of scratchcards are scored like normal scratchcards and have the same card number as the card they copied.
// So, if you win a copy of card 10 and it has 5 matching numbers,
// it would then win a copy of the same cards that the original card 10 won:
// cards 11, 12, 13, 14, and 15. This process repeats until none of the copies cause you to win any more cards.
// (Cards will never make you copy a card past the end of the table.)
// This time, the above example goes differently:
// Card 1 has four matching numbers, so you win one copy each of the next four cards: cards 2, 3, 4, and 5.
// Your original card 2 has two matching numbers, so you win one copy each of cards 3 and 4.
// Your copy of card 2 also wins one copy each of cards 3 and 4.
// Your four instances of card 3 (one original and three copies) have two matching numbers, so you win four copies each of cards 4 and 5.
// Your eight instances of card 4 (one original and seven copies) have one matching number, so you win eight copies of card 5.
// Your fourteen instances of card 5 (one original and thirteen copies) have no matching numbers and win no more cards.
// Your one instance of card 6
// you end up with 1 instance of card 1,
// 2 instances of card 2,
// 4 instances of card 3,
// 8 instances of card 4,
// 14 instances of card 5,
// and 1 instance of card 6.
// In total, this example pile of scratchcards causes you to ultimately have 30 scratchcards!

func GetInput() []string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("WORKDIR_Error: %v\n", err)
	}

	p := path.Join(wd, "cards", "input.txt")
	b, err := os.ReadFile(p)
	if err != nil {
		fmt.Printf("WORKDIR_Error: %v\n", err)
	}

	return strings.Split(string(b), "\n")
}

func ScratchCards(a []string) int {
	ln := len(a) - 1
	cards := make(map[int]int, ln)
	for i := 0; i < ln; i++ {
		cards[i] = 1
	}

	res := ln
	for i, v := range a {
		blah := strings.Split(v, ": ")
		if len(blah) > 1 {
			c := strings.Split(blah[1], " | ")
			wins := 0
			for _, a := range strings.Fields(c[0]) {
				for _, b := range strings.Fields(c[1]) {
					if a == b {
						wins++
						res += cards[i]
						cards[i+wins] += cards[i]
						break
					}
				}
			}
		}
	}

	return res
}

func Cards(a []string) int {
	ln := len(a) - 1
	cards := make(map[int]int, ln)
	for i := 0; i < ln; i++ {
		cards[i] = 1
	}

	res := ln

	linesCh := make(chan string, len(a))
	var wg sync.WaitGroup

	for i := 0; i < len(a); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			linesCh <- a[i]
		}(i)
	}

	go func() {
		wg.Wait()
		close(linesCh)
	}()

	for line := range linesCh {
		v := line
		idx := strings.Index(v, ": ")
		if idx > 0 {
			fieldsA := strings.Fields(v[idx+2:])
			for j := range linesCh {
        x, _ := strconv.Atoi(j)
				v2 := a[x]
				idx2 := strings.Index(v2, ": ")
				if idx2 > 0 {
					fieldsB := strings.Fields(v2[idx2+2:])
					wins := 0
					for _, a := range fieldsA {
						for _, b := range fieldsB {
							if a == b {
								wins++
								res += cards[x]
								cards[x+wins] += cards[x]
								break
							}
						}
					}
				}
			}
		}
	}

	return res
}


func main() {
	x := Cards(GetInput())
	y := ScratchCards(GetInput())
	fmt.Println(x, y)
}

//
//func cti(s string) int {
//	i, err := strconv.Atoi(s)
//	if err != nil {
//		fmt.Println("ERR", err)
//		return 0
//	}
//	return i
//}
//
//func cardBigFacts(w, m []string) int {
//	count := 0
//	for _, a := range w {
//	  for _, b := range m {
//      if a == b && count > 0 {
//        count = count + count
//      } else if a == b {
//        count++
//      }
//    }
//	}
//	return count
//}
//
//func scratchCards(a []string) int {
//	res := 0
//	for _, v := range a {
//		blah := strings.Split(v, ": ")
//		if len(blah) > 1 {
//			card := strings.Split(blah[1], " | ")
//			res += cardBigFacts(strings.Fields(card[0]), strings.Fields(card[1]))
//		}
//	}
//	return res
//}
//
//
