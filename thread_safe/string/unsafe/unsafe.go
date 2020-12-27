package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	m := ""

	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i string) {
			defer wg.Done()
			m += i
		}(fmt.Sprint(i % 10))
	}

	wg.Wait()
	log.Println(len(m))
	log.Println(m)
}
