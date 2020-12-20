package main

import (
	"log"
	"sync/atomic"
)

const (
	routineNum = 4
)

func main() {
	m := map[int]int32{}

	go func() {
		var i int32 = 0
		for {
			m[1] = i
			atomic.AddInt32(&i, 1)
		}
	}()

	for {
		log.Printf("%+v\n", m[1])
	}
}
