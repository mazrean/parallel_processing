package main

import (
	"log"
	"sync"
)

const (
	routineNum = 4
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
}
