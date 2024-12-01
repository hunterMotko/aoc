package main

import (
	"testing"
)

var Result int
func BenchmarkScratchCards(b *testing.B) {
	var x int
	for i := 0; i < b.N; i++ {
		x = ScratchCards(GetInput())
	}
	Result = x
}

func BenchmarkCards(b *testing.B) {
	var x int
	for i := 0; i < b.N; i++ {
		x = Cards(GetInput())
	}
	Result = x
}
