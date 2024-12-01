package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

const (
	_ int = iota
	HighCard
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var VALUES = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func Value(b rune) int {
	return VALUES[b]
}

type Cards []byte
func (c Cards) Len() int           { return len(c) }
func (c Cards) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c Cards) Less(i, j int) bool { return Value(rune(c[i])) < Value(rune(c[j])) }

type Hands []Hand
func (h Hands) Len() int      { return len(h) }
func (h Hands) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h Hands) Less(i, j int) bool {
	if h[i].Level == h[j].Level {
		return h[i].Points < h[j].Points
	}
	return h[i].Level < h[j].Level
}
func (h Hands) String() string {
	s := ""
	for i := range h {
		s += fmt.Sprintf("%5d : %s [L:%d] (%7d pts)\n", i+1, h[i].Raw, h[i].Level, h[i].Points)
	}
	return s
}

type Hand struct {
	Raw    string
	Bid    int
	Points int
	Level  int
}

//go:embed input.txt
var in string

func ParseInput() []Hand {
	rounds := strings.Split(in, "\n")
	hands := make([]Hand, 0)
	for _, line := range rounds {
		if line != "" {
			data := strings.Split(line, " ")
			raw := data[0]
			bid, _ := strconv.Atoi(data[1])
			h := Hand{
				Raw: raw,
				Bid: bid,
			}
			hands = append(hands, h)
		}
	}
	return hands
}

func HandLevel(hand string) int {
	sorted := Cards([]byte(hand))
	sort.Sort(sorted)
  fmt.Println(sorted)
	c1, c2, c3, c4, c5 := sorted[0], sorted[1], sorted[2], sorted[3], sorted[4]
	switch {
	case c1 == c5:
		return FiveOfAKind
	case c1 == c4 || c2 == c5:
		return FourOfAKind
	case (c1 == c3 && c4 == c5) || (c1 == c2 && c3 == c5):
		return FullHouse
	case c1 == c3 || c2 == c4 || c3 == c5:
		return ThreeOfAKind
	case (c1 == c2 && c3 == c4) || (c2 == c3 && c4 == c5) || (c1 == c2 && c4 == c5):
		return TwoPair
	case c1 == c2 || c2 == c3 || c3 == c4 || c4 == c5:
		return OnePair
	default:
		return HighCard
	}
}

func PointHand(hand string) int {
	points := 0
	for i := range hand {
		points += Value(rune(hand[i])) * int(math.Pow10(10-i*2))
	}
	return points
}

func ReplaceJokers(hand string) string {
	if hand == "JJJJJ" {
		return hand
	}
	counts := make(map[rune]int)
	var maxRune rune
	var maxCount int
	for _, c := range hand {
		if c == 'J' {
			continue
		}
		counts[c]++
		if counts[c] > maxCount {
			maxRune = c
			maxCount = counts[c]
		}
		if counts[c] == maxCount && Value(c) > Value(maxRune) {
			maxRune = c
		}
	}
	mod := strings.ReplaceAll(hand, "J", string(maxRune))

	return mod
}

func PartOne(hands Hands) int {
	for i := range hands {
		hands[i].Level = HandLevel(hands[i].Raw)
		hands[i].Points = PointHand(hands[i].Raw)
	}
	sort.Sort(hands)
	fmt.Println(hands)

	winnings := 0
	for i := range hands {
		winnings += hands[i].Bid * (i + 1)
	}
	return winnings
}

func PartTwo(hands Hands) int {
	VALUES['J'] = 1
	for i := range hands {
		hand := hands[i].Raw
		if strings.Contains(hand, "J") {
			hand = ReplaceJokers(hand)
		}
		hands[i].Level = HandLevel(hand)
		hands[i].Points = PointHand(hands[i].Raw)
	}
	sort.Sort(hands)
	fmt.Println(hands)
	winnings := 0
	for i := range hands {
		winnings += hands[i].Bid * (i + 1)
	}
	return winnings
}

func runP() {
	fmt.Println(PartTwo(ParseInput()))
}

// func main() {
// 	fmt.Println(PartOne(ParseInput("input.txt")))
// 	fmt.Println(PartTwo(ParseInput("input.txt")))
// }
