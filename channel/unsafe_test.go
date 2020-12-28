package main

import (
	"sync"
	"testing"
)

func BenchmarkUnsafe(b *testing.B) {
	s := make([]int32, 0, 2*b.N)
	wg := sync.WaitGroup{}

	b.ResetTimer()
	wg.Add(1)
	go func(s *[]int32) {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			*s = append(*s, int32(i))
		}
	}(&s)

	wg.Add(1)
	go func(s *[]int32) {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			*s = append(*s, int32(i))
		}
	}(&s)
	wg.Wait()
	for _, v := range s {
		f(v)
	}
}
