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
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	go despair(ctx)

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
