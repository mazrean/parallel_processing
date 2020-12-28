package main

import (
	"log"
	"sort"
	"sync"
)

func main() {
	l := &sync.Mutex{}
	m := make([]int32, 0, 100)

	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int32) {
			defer wg.Done()
			l.Lock()
			m = append(m, i)
			l.Unlock()
		}(int32(i))
	}

	wg.Wait()
	log.Println(len(m))
	sort.Slice(m, func(i, j int) bool { return m[i] < m[j] })
	log.Println(m)
}
