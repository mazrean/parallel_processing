package main

import (
	"log"
	"sync/atomic"
)

type hoge struct {
	hoge int32
	huga int32
}

func main() {
	m := hoge{0, 0}

	go func() {
		var i int32 = 0
		for {
			m = hoge{i, i}
			atomic.AddInt32(&i, 1)
			_ = m.hoge + m.huga
		}
	}()

	for {
		n := m
		if n.hoge != n.huga {
			log.Panicln(n.hoge, n.huga)
		}
	}
}
