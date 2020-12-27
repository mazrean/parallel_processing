package main

import (
	"log"
	"sync"
)

func main() {
	var m int32 = 0

	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int32) {
			defer wg.Done()
			m += i
		}(int32(i % 10))
	}

	wg.Wait()
	log.Println(m)
}
