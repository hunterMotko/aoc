package main

import (
	"container/heap"
	"fmt"
	// "sort"
)

type kv struct {
	Key   string
	Value int
}

func getHeap(m map[string]int) *KVHeap {
	h := &KVHeap{}
	heap.Init(h)
	for k, v := range m {
		heap.Push(h, kv{k, v})
	}
	return h
}

func getMap() map[string]int {
	m := make(map[string]int)
	m["A"] = 13
	m["K"] = 12
	m["Q"] = 11
	m["T"] = 10
	m["9"] = 9
	m["8"] = 8
	m["7"] = 7
	m["6"] = 6
	m["5"] = 5
	m["4"] = 4
	m["3"] = 3
	m["2"] = 2
	m["J"] = 1
	return m
}

type KVHeap []kv

func (a KVHeap) Less(i, j int) bool { return a[i].Value < a[j].Value }
func (a KVHeap) Len() int           { return len(a) }
func (a KVHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (h *KVHeap) Push(x interface{}) {
	*h = append(*h, x.(kv))
}

func (h *KVHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func run() {
  m := getMap()
  heap := getHeap(m)
//  sort.Sort(heap)
  n := 3
  fmt.Println(heap)
  for i := 0; i < n; i++ {
    fmt.Printf("%d) %v#v\n", i+1, heap.Pop())
  }
}
