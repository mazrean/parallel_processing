package main

import (
	"log"
	"sync"
	"sync/atomic"
)

const (
	routineNum = 4
)

type hoge struct {
	hoge int32
}

func main() {
	m := map[int]hoge{}

	wg := sync.WaitGroup{}

	wg.Add(1) //wg:1
	go func() {
		defer wg.Done()
		var i int32 = 0
		for {
			m[1] = hoge{
				hoge: i,
			}
			atomic.AddInt32(&i, 1)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			log.Printf("%+v\n", m[1])
		}
	}()

	wg.Wait()
}
