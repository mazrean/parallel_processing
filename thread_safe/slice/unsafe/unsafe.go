package main

import (
	"log"
	"sort"
	"sync"
)

func main() {
	m := make([]int32, 0, 100)

	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int32) {
			defer wg.Done()
			m = append(m, i)
		}(int32(i))
	}

	wg.Wait()
	log.Println(len(m))
	sort.Slice(m, func(i, j int) bool { return m[i] < m[j] })
	log.Println(m)
}
