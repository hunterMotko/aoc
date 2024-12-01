package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Delimiter struct {
	Ind       []int
	s         string
	del       string
	trimSpace bool
}

type delimiterOptions struct {
	trimSpace bool
}

type DelimiterOption func(options *delimiterOptions)

func WithTrimSpace() DelimiterOption {
	return func(options *delimiterOptions) {
		options.trimSpace = true
	}
}

func NewDelimiter(s, del string, opts ...DelimiterOption) Delimiter {
	var options delimiterOptions
	for _, opt := range opts {
		opt(&options)
	}
	return Delimiter{
		Ind:       IndexAll(s, del),
		s:         s,
		del:       del,
		trimSpace: options.trimSpace,
	}
}

func (d Delimiter) GetStrings() []string {
	if len(d.Ind) == 0 {
		if d.s == "" {
			return nil
		}
		if d.trimSpace {
			return []string{strings.TrimSpace(d.s)}
		}
		return []string{d.s}
	}
	var res []string
	for i := 0; i <= len(d.Ind); i++ {
		s := d.GetString(i)
		if s == "" {
			continue
		}
		res = append(res, s)
	}
	return res
}

func (d Delimiter) GetInts() []int {
	return StringsToInts(d.GetStrings())
}

func (d Delimiter) GetString(i int) string {
	s := ""
	if i == 0 {
		s = d.s[:d.Ind[0]]
	} else if i == len(d.Ind) {
		s = d.s[d.Ind[len(d.Ind)-1]+len(d.del):]
	} else {
		s = d.s[d.Ind[i-1]+len(d.del) : d.Ind[i]]
	}
	if d.trimSpace {
		return strings.TrimSpace(s)
	}
	return s
}

func (d Delimiter) GetInt(i int) int {
	return StringToInt(d.GetString(i))
}

func (d Delimiter) TryGetInt(i int) (int, bool) {
	return TryStringToInt(d.GetString(i))
}

func (d Delimiter) IsInt(i int) bool {
	_, err := strconv.Atoi(d.GetString(i))
	return err == nil
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func TryStringToInt(s string) (int, bool) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}
	return i, true
}

func StringsToInts(s []string) []int {
	res := make([]int, len(s))
	for i, v := range s {
		res[i] = StringToInt(v)
	}
	return res
}

func Substring(s string, del string) string {
	idx := strings.Index(s, del)
	if idx == -1 {
		panic(fmt.Sprintf("substring: %s is not present in %s", del, s))
	}
	return s[idx+len(del):]
}

func IndexAll(s string, search string) []int {
	i := 0
	var res []int
	for i < len(s) {
		index := strings.Index(s[i:], search)
		if index == -1 {
			return res
		}
		res = append(res, index+i)
		i += index + len(search)
	}
	return res
}

func StringGroups(lines []string) [][]string {
	i := 0
	var res [][]string
	var row []string
	for {
		row = append(row, lines[i])
		i++
		if i >= len(lines) {
			res = append(res, row)
			break
		}
		for ; i < len(lines); i++ {
			if lines[i] == "" {
				break
			} else {
				row = append(row, lines[i])
			}
		}
		res = append(res, row)
		row = nil
		i++
		if i >= len(lines) {
			break
		}
	}
	return res
}

func ReaderToStrings(input io.Reader) []string {
	scanner := bufio.NewScanner(input)
	var res []string
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}

func fs1(input io.Reader) int {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	groups := StringGroups(ReaderToStrings(f))
	seeds := parseSeeds(groups[0][0])
	var maps []Map
	for i := 1; i < len(groups); i++ {
		maps = append(maps, parseMap(groups[i]))
	}
	lowest := math.MaxInt
	for _, v := range seeds {
		for _, m := range maps {
			if dst, contains := m.get(v); contains {
				v = dst
			}
		}
		lowest = min(lowest, v)
	}
	return lowest
}

func parseSeeds(line string) []int {
	line = Substring(line, ": ")
	del := NewDelimiter(line, " ")
	return del.GetInts()
}

type Range struct {
	from    int
	to      int
	transfo int
}

type Map struct {
	ranges []Range
}

func (m Map) get(v int) (int, bool) {
	l := 0
	r := len(m.ranges) - 1
	for l <= r {
		mid := l + (r-l)/2
		rng := m.ranges[mid]
		if v > rng.to {
			l = mid + 1
		} else if v < rng.from {
			r = mid - 1
		} else {
			return v + rng.transfo, true
		}
	}

	return 0, false
}

func parseMap(lines []string) Map {
	var ranges []Range
	for i := 0; i < len(lines); i++ {
		if i == 0 {
			// Discard header
			continue
		}
		del := NewDelimiter(lines[i], " ")
		ints := del.GetInts()
		dstRange := ints[0]
		srcRange := ints[1]
		rangeLength := ints[2]

		ranges = append(ranges, Range{
			from:    srcRange,
			to:      srcRange + rangeLength - 1,
			transfo: dstRange - srcRange,
		})
	}
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].from < ranges[j].from
	})
	return Map{ranges: ranges}
}

func fs2NonFinal(input io.Reader) int {
	groups := StringGroups(ReaderToStrings(input))

	seeds := parseSeedsRange(groups[0][0])
	var maps []Map
	for i := 1; i < len(groups); i++ {
		maps = append(maps, parseMap(groups[i]))
	}

	mu := sync.Mutex{}
	lowest := math.MaxInt
	wg, _ := errgroup.WithContext(context.Background())
	for _, seed := range seeds {
		seed := seed
		wg.Go(func() error {
			local := math.MaxInt
			for i := 0; i < seed[1]; i++ {
				local = min(local, transform(seed[0]+i, maps))
			}
			mu.Lock()
			lowest = min(lowest, local)
			mu.Unlock()
			return nil
		})
	}

	_ = wg.Wait()
	return lowest
}

func transform(v int, maps []Map) int {
	for _, m := range maps {
		if dst, contains := m.get(v); contains {
			v = dst
		}
	}
	return v
}

func parseSeedsRange(line string) [][2]int {
	line = Substring(line, ": ")
	del := NewDelimiter(line, " ")
	ints := del.GetInts()
	var res [][2]int
	for i := 0; i < len(ints); i += 2 {
		res = append(res, [2]int{ints[i], ints[i+1]})
	}
	return res
}

func fs2(input io.Reader) int {
	groups := StringGroups(ReaderToStrings(input))
	seeds := parseSeedsRange(groups[0][0])
	sort.Slice(seeds, func(i, j int) bool {
		return seeds[i][0] < seeds[j][0]
	})
	var maps []Map
	for i := 1; i < len(groups); i++ {
		maps = append(maps, parseRevMap(groups[i]))
	}
	for loc := 0; loc < math.MaxInt; loc++ {
		v := transformRev(loc, maps)
		if isWithinSeedRange(v, seeds) {
			return loc
		}
	}

	return -1
}

func transformRev(v int, maps []Map) int {
	for i := len(maps) - 1; i >= 0; i-- {
		m := maps[i]
		if dst, contains := m.get(v); contains {
			v = dst
		}
	}
	return v
}

func parseRevMap(lines []string) Map {
	var ranges []Range
	for i := 0; i < len(lines); i++ {
		if i == 0 {
			// Discard header
			continue
		}
		del := NewDelimiter(lines[i], " ")
		ints := del.GetInts()
		dstRange := ints[0]
		srcRange := ints[1]
		rangeLength := ints[2]

		ranges = append(ranges, Range{
			from:    dstRange,
			to:      dstRange + rangeLength - 1,
			transfo: srcRange - dstRange,
		})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].from < ranges[j].from
	})
	return Map{ranges: ranges}
}

func isWithinSeedRange(n int, seeds [][2]int) bool {
	l := 0
	r := len(seeds) - 1
	for l <= r {
		mid := l + (r-l)/2
		rng := seeds[mid]
		if n < rng[0] {
			r = mid - 1
		} else if n > rng[0]+rng[1]-1 {
			l = mid + 1
		} else {
			return true
		}
	}

	return false
}

func main() {
	f, err := os.Open("./seed/input.txt")
	if err != nil {
		fmt.Println(err)
	}
  fmt.Println(fs2(f))
}
