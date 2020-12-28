package main

import (
	"bytes"
	"io"
	"testing"
)

const (
	CopyFile1 = "cp1.tar.gz"
	//CP_FILE2 = "cp2.tar.gz"
	FromFile = "go1.15.6.linux-amd64.tar.gz"
)

func BenchmarkPipe(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pr, pw := io.Pipe()
		defer pr.Close()

		go func() {
			defer pw.Close()
			FileReader(FromFile, pw)
			return
		}()

		FileWriter(CopyFile1, pr)
	}

	return
}

func BenchmarkCopy(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bf := &bytes.Buffer{}

		FileReader(FromFile, bf)

		FileWriter(CopyFile1, bf)
	}

	return
}
