package main

import (
	"fmt"
	"time"
)

func despair() {
	for {
		fmt.Println("Help me!")
	}
}

func main() {
	go despair()

	time.Sleep(time.Second * 1)
	/*この時点でエラーが発生。
	しかしdespairは知らされず苦しみ続ける…*/
	fmt.Println("error!")

	time.Sleep(time.Second * 1)
}
