package main

import (
	"io"
	"os"
)

func FileReader(file string, w io.Writer) {
	fromFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fromFile.Close()

	_, err = io.Copy(w, fromFile)
	if err != nil {
		panic(err)
	}
	return
}

func FileWriter(file string, r io.Reader) {
	cpFile1, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer cpFile1.Close()

	_, err = io.Copy(cpFile1, r)
	if err != nil {
		panic(err)
	}

	return
}
