package main

import (
	"log"
	"sync"
	"sync/atomic"
)

func main() {
	l := &sync.Mutex{}
	m := map[int]int32{}

	go func() {
		var i int32 = 0
		for {
			l.Lock()
			m[1] = i
			l.Unlock()
			atomic.AddInt32(&i, 1)
		}
	}()

	for {
		l.Lock()
		log.Printf("%+v\n", m)
		l.Unlock()
	}
}
