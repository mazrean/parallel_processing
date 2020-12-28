package main

import (
	"sync"
	"testing"
)

func BenchmarkBad(b *testing.B) {
	l := new(sync.Mutex)
	s := make([]int32, 0, 2*b.N)
	wg := sync.WaitGroup{}

	b.ResetTimer()
	wg.Add(1)
	go func(l sync.Locker, s *[]int32) {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			l.Lock()
			*s = append(*s, int32(i))
			l.Unlock()
		}
	}(l, &s)

	wg.Add(1)
	go func(l sync.Locker, s *[]int32) {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			l.Lock()
			*s = append(*s, int32(i))
			l.Unlock()
		}
	}(l, &s)

	wg.Wait()
	for _, v := range s {
		f(v)
	}
}
