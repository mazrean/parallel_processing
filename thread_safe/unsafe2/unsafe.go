package main

import (
	"log"
	"sync/atomic"
)

const (
	routineNum = 4
)

func main() {
	m := []int32{0}

	go func() {
		var i int32 = 0
		for {
			m = append(m, i)
			atomic.AddInt32(&i, 1)
			log.Printf("a: %d", i)
		}
	}()

	var i int32 = 0
	for {
		m = append(m, i)
		atomic.AddInt32(&i, 1)
	}
}
