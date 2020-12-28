package main

import (
	"context"
	"fmt"
	"time"
)

func despair(ctx context.Context) {
	for {
		fmt.Println("Help me!")

		select {
		case _ = <-ctx.Done():
			fmt.Println("...")
			return
		default:
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go despair(ctx)

	time.Sleep(time.Second * 1)
	//この時点でエラーが発生。
	fmt.Println("error!")
	//despairに知らされる
	cancel()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		time.Sleep(time.Second * 1)
	}
}
