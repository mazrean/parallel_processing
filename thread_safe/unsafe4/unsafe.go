package main

import (
	"sync/atomic"
)

const (
	routineNum = 4
)

func main() {
	var m int32 = 0

	go func() {
		var i int32 = 0
		for i < 0x7fff {
			m += 1
			_ = m
			atomic.AddInt32(&i, 1)
		}
	}()

	var i int32 = 0x7fff
	for i >= 0 {
		m = i
		if m < 0 {
		}
		atomic.AddInt32(&i, -1)
	}
}
