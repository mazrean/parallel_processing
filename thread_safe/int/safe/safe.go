package main

import (
	"log"
	"sync"
	"sync/atomic"
)

func main() {
	var m int32 = 0

	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int32) {
			defer wg.Done()
			atomic.AddInt32(&m, i)
		}(int32(i % 10))
	}

	wg.Wait()
	log.Println(m)
}
