package main

import (
	"context"
	"fmt"
	"time"
)

func despair(ctx context.Context) {
	innerCtx, cancel := context.WithCancel(ctx)
	//この時点でエラーが発生。
	fmt.Println("error!")
	//cancel
	cancel()
	for {
		fmt.Println("Help me!")

		select {
		case _ = <-innerCtx.Done():
			fmt.Println("...")
			return
		default:
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go despair(ctx)

	time.Sleep(time.Second * 1)

	select {
	case _ = <-ctx.Done():
		fmt.Println(ctx.Err())
	default:
		//親は何も分からない
		fmt.Println("???")
		time.Sleep(time.Millisecond * 100)
	}
}
