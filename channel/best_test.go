package main

import (
	"testing"
)

func BenchmarkBest(b *testing.B) {
	out := make(chan int32, 2*b.N)

	b.ResetTimer()
	go func(out chan int32) {
		for i := 0; i < b.N; i++ {
			out <- int32(i)
		}
	}(out)

	go func(out chan int32) {
		for i := 0; i < b.N; i++ {
			out <- int32(i)
		}
	}(out)

	for i := 0; i < 2*b.N; i++ {
		v := <-out
		f(v)
	}
	close(out)
}
