package main

import "testing"

func BenchmarkOverhead(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() {}()
	}
}

func BenchmarkNone(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
	}
}
